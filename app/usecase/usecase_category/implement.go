package usecase_category

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
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewUseCase(categoryRepo repository.CategoryRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:           validate,
		CategoryRepository: categoryRepo,
	}
}

func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.Category{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	dataRequest.CategoryName = html.EscapeString(dataRequest.CategoryName)
	category := controller.CategoryRepository.Create(r.Context(), dataRequest)

	dataResponse := map[string]interface{}{
		"category": category,
	}
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["categoryId"]
	dataResponse, err := controller.CategoryRepository.FindById(r.Context(), id)
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
		"limit":    limit,
		"offset":   offset,
		"search":   search,
		"category": query.Get("category"),
	}

	dataResponse, totalRows := controller.CategoryRepository.FindAll(r.Context(), conf)

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

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["categoryId"]
	dataRequest := entity.Category{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	dataCategory := entity.Category{
		CategoryName: dataRequest.CategoryName,
	}
	dataResponse := controller.CategoryRepository.Update(r.Context(), []string{"category_name"}, dataCategory, id)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["categoryId"]
	_, err := controller.CategoryRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	controller.CategoryRepository.Delete(r.Context(), id)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
