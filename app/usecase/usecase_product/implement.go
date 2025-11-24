package usecase_product

import (
	"encoding/json"
	"html"
	"net/http"
	"strconv"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/app/service"
	"github.com/edwinjordan/e-canteen-backend/config"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/handler"
	"github.com/edwinjordan/e-canteen-backend/pkg/exceptions"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type UseCaseImpl struct {
	ProductRepository repository.ProductRepository
	MinioService      service.MinioService
	Validate          *validator.Validate
}

func NewUseCase(ProductRepo repository.ProductRepository, minioService service.MinioService, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:          validate,
		ProductRepository: ProductRepo,
		MinioService:      minioService,
	}
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["productId"]
	dataResponse, err := controller.ProductRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	Qlimit := query.Get("limit")
	Qoffset := query.Get("offset")
	Qpage := query.Get("page")
	search := query.Get("search")

	if Qlimit == "" {
		Qlimit = "10"
	}

	page := 1
	if Qpage != "" {
		page, _ = strconv.Atoi(Qpage)
		if page < 1 {
			page = 1
		}
	}

	limit, _ := strconv.Atoi(Qlimit)
	offset := (page - 1) * limit

	if Qoffset != "" {
		offset, _ = strconv.Atoi(Qoffset)
	}

	conf := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
		"search": search,
	}

	where := entity.Product{
		ProductCategoryId: query.Get("category_id"),
	}
	dataResponse, totalRows := controller.ProductRepository.FindAll(r.Context(), where, conf)

	// Calculate pagination info
	totalPage := (totalRows + limit - 1) / limit
	if totalPage < 1 {
		totalPage = 1
	}

	webResponse := handler.WebResponseWithPagination{
		Code:    200,
		Success: true,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
		Page: &handler.PageInfo{
			CurrentPage: page,
			TotalPage:   totalPage,
			TotalRows:   totalRows,
			PerPage:     limit,
		},
	}
	helpers.WriteToResponseBody(w, webResponse)
}

type ProductInsertRequest struct {
	ProductCode       string          `json:"product_code" validate:"required"`
	ProductName       string          `json:"product_name" validate:"required"`
	ProductCategoryId string          `json:"product_category_id" validate:"required"`
	ProductDesc       string          `json:"product_desc"`
	ProductPhoto      string          `json:"product_photo"`
	Varians           []VarianRequest `json:"varians" validate:"required,min=1,dive"`
}

type VarianRequest struct {
	ProductVarianId           string `json:"product_varian_id" validate:"omitempty"`
	VarianId                  string `json:"varian_id" validate:"omitempty"`
	VarianName                string `json:"varian_name" validate:"required"`
	ProductVarianPrice        int    `json:"product_varian_price" validate:"required,min=0"`
	ProductVarianQtyBooth     int    `json:"product_varian_qty_booth" validate:"required,min=0"`
	ProductVarianQtyWarehouse int    `json:"product_varian_qty_warehouse" `
}

func (controller *UseCaseImpl) Insert(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		panic(exceptions.NewBadRequestError("Gagal memproses data: " + err.Error()))
	}

	dataRequest := ProductInsertRequest{}
	dataRequest.ProductCode = r.FormValue("product_code")
	dataRequest.ProductName = r.FormValue("product_name")
	dataRequest.ProductCategoryId = r.FormValue("product_category_id")
	dataRequest.ProductDesc = r.FormValue("product_desc")

	// Parse varians from JSON string
	variansStr := r.FormValue("varians")
	if variansStr != "" {
		err = json.Unmarshal([]byte(variansStr), &dataRequest.Varians)
		if err != nil {
			panic(exceptions.NewBadRequestError("Format varians tidak valid: " + err.Error()))
		}
	}

	// Handle Image Upload
	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		// Upload to MinIO
		uploadedPath, err := controller.MinioService.UploadFile(r.Context(), header, "products")
		if err != nil {
			panic(exceptions.NewInternalServerError("Gagal mengunggah gambar: " + err.Error()))
		}

		// Get URL
		url, _ := controller.MinioService.GetFileUrl(r.Context(), uploadedPath)
		dataRequest.ProductPhoto = url
	}

	err = controller.Validate.Struct(dataRequest)
	if err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	// Sanitize product data
	productId := helpers.GenUUID()
	product := entity.Product{
		ProductId:         productId,
		ProductCode:       html.EscapeString(dataRequest.ProductCode),
		ProductName:       html.EscapeString(dataRequest.ProductName),
		ProductCategoryId: dataRequest.ProductCategoryId,
		ProductDesc:       html.EscapeString(dataRequest.ProductDesc),
		ProductPhoto:      dataRequest.ProductPhoto,
	}

	// Prepare varians
	var varians []entity.Varian
	for _, v := range dataRequest.Varians {
		varianId := helpers.GenUUID()
		varians = append(varians, entity.Varian{
			ProductVarianId:           varianId,
			ProductId:                 productId,
			VarianId:                  v.VarianId,
			VarianName:                v.VarianName,
			ProductVarianPrice:        v.ProductVarianPrice,
			ProductVarianQtyBooth:     v.ProductVarianQtyBooth,
			ProductVarianQtyWarehouse: v.ProductVarianQtyWarehouse,
		})
	}

	// Insert product and varians
	err = controller.ProductRepository.Insert(r.Context(), product, varians)
	if err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	dataResponse := map[string]interface{}{
		"product_id": productId,
		"message":    "Product and variants created successfully",
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

type ProductUpdateRequest struct {
	ProductCode       string          `json:"product_code" validate:"required"`
	ProductName       string          `json:"product_name" validate:"required"`
	ProductCategoryId string          `json:"product_category_id" validate:"required"`
	ProductDesc       string          `json:"product_desc"`
	ProductPhoto      string          `json:"product_photo"`
	Varians           []VarianRequest `json:"varians" validate:"omitempty,dive"`
}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		panic(exceptions.NewBadRequestError("Gagal memproses data: " + err.Error()))
	}

	dataRequest := ProductUpdateRequest{}
	dataRequest.ProductCode = r.FormValue("product_code")
	dataRequest.ProductName = r.FormValue("product_name")
	dataRequest.ProductCategoryId = r.FormValue("product_category_id")
	dataRequest.ProductDesc = r.FormValue("product_desc")

	// Parse varians from JSON string
	variansStr := r.FormValue("varians")
	if variansStr != "" {
		err = json.Unmarshal([]byte(variansStr), &dataRequest.Varians)
		if err != nil {
			panic(exceptions.NewBadRequestError("Format varians tidak valid: " + err.Error()))
		}
	}

	// Handle Image Upload
	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		// Upload to MinIO
		uploadedPath, err := controller.MinioService.UploadFile(r.Context(), header, "products")
		if err != nil {
			panic(exceptions.NewInternalServerError("Gagal mengunggah gambar: " + err.Error()))
		}

		// Get URL
		url, _ := controller.MinioService.GetFileUrl(r.Context(), uploadedPath)
		dataRequest.ProductPhoto = url
	} else {
		// Keep existing photo if not updated (optional, depends on frontend logic)
		// If frontend sends empty image, we might want to keep old one.
		// But here we just set what we have. If empty, it might overwrite with empty if we don't check.
		// Let's check existing product first to be safe, or assume frontend sends old URL if no change?
		// Usually for update, if image is not provided, we keep old one.
		// But here we are constructing a new object.
		// Let's fetch existing product to get old photo if new one is not provided.
		existingProduct, err := controller.ProductRepository.FindById(r.Context(), productId)
		if err == nil {
			dataRequest.ProductPhoto = existingProduct.ProductPhoto
		}
	}

	err = controller.Validate.Struct(dataRequest)
	if err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	// Sanitize product data
	product := entity.Product{
		ProductId:         productId,
		ProductCode:       html.EscapeString(dataRequest.ProductCode),
		ProductName:       html.EscapeString(dataRequest.ProductName),
		ProductCategoryId: dataRequest.ProductCategoryId,
		ProductDesc:       html.EscapeString(dataRequest.ProductDesc),
		ProductPhoto:      dataRequest.ProductPhoto,
	}

	// Prepare varians if provided
	var varians []entity.Varian

	if len(dataRequest.Varians) > 0 {
		for _, v := range dataRequest.Varians {

			// Generate UUID for new varians (if product_varian_id is empty)
			varianId := v.VarianId
			// if varianId == "" {
			// 	varianId = helpers.GenUUID()
			// }

			productVarianId := ""

			varians = append(varians, entity.Varian{
				ProductVarianId:           varianId,
				ProductId:                 productId,
				VarianId:                  productVarianId,
				VarianName:                v.VarianName,
				ProductVarianPrice:        v.ProductVarianPrice,
				ProductVarianQtyBooth:     v.ProductVarianQtyBooth,
				ProductVarianQtyWarehouse: v.ProductVarianQtyWarehouse,
			})
		}
	}

	// Update product and varians
	err = controller.ProductRepository.Update(r.Context(), product, varians)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	dataResponse := map[string]interface{}{
		"product_id": productId,
		"message":    "Product and variants updated successfully",
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	// Delete product (soft delete)
	err := controller.ProductRepository.Delete(r.Context(), productId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	dataResponse := map[string]interface{}{
		"product_id": productId,
		"message":    "Product deleted successfully",
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
