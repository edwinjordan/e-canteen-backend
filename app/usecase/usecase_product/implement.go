package usecase_product

import (
	"html"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/config"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/handler"
	"github.com/edwinjordan/e-canteen-backend/pkg/exceptions"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
)

type UseCaseImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewUseCase(ProductRepo repository.ProductRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:          validate,
		ProductRepository: ProductRepo,
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
	dataRequest := ProductInsertRequest{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
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

	dataRequest := ProductUpdateRequest{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
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
			if varianId == "" {
				varianId = helpers.GenUUID()
			}

			productVarianId := v.ProductVarianId
			if productVarianId == "" {
				productVarianId = helpers.GenUUID()
			}

			varians = append(varians, entity.Varian{
				ProductVarianId:           productVarianId,
				ProductId:                 productId,
				VarianId:                  varianId,
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
