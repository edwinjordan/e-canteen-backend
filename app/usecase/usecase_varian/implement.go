package usecase_varian

import (
	"net/http"

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
	VarianRepository repository.VarianRepository
	Validate         *validator.Validate
}

func NewUseCase(varianRepo repository.VarianRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:         validate,
		VarianRepository: varianRepo,
	}
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["varianId"]
	dataResponse, err := controller.VarianRepository.FindById(r.Context(), id)
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
	vars := r.URL.Query()
	where := entity.Varian{}

	if vars.Get("product_id") != "" {
		where.ProductId = vars.Get("product_id")
	}

	if vars.Get("varian_id") != "" {
		where.VarianId = vars.Get("varian_id")
	}
	dataResponse := controller.VarianRepository.FindSpesificData(r.Context(), where)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
