package usecase_major

import (
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
	MajorRepository repository.MajorRepository
	Validate        *validator.Validate
}

func NewUseCase(majorRepo repository.MajorRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:        validate,
		MajorRepository: majorRepo,
	}
}

func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.Major{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	dataResponse := controller.MajorRepository.Create(r.Context(), dataRequest)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["majorId"]
	dataRequest := entity.Major{}
	helpers.ReadFromRequestBody(r, &dataRequest)
	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	_, err = controller.MajorRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	dataRequest.MajorId = id
	dataResponse := controller.MajorRepository.Update(r.Context(), dataRequest, id)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["majorId"]
	_, err := controller.MajorRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	controller.MajorRepository.Delete(r.Context(), id)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["majorId"]
	dataResponse, err := controller.MajorRepository.FindById(r.Context(), id)
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
		"major":  query.Get("major"),
	}

	dataResponse, totalRows := controller.MajorRepository.FindAll(r.Context(), conf)

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
