package usecase_tempcart

import (
	"net/http"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/config"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/handler"
	"github.com/edwinjordan/e-canteen-backend/pkg/exceptions"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type UseCaseImpl struct {
	TempCartRepository repository.TempCartRepository
	VarianRepository   repository.VarianRepository
	Validate           *validator.Validate
}

func NewUseCase(tempcartRepo repository.TempCartRepository, productRepo repository.VarianRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:           validate,
		TempCartRepository: tempcartRepo,
		VarianRepository:   productRepo,
	}
}

func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.TempCart{}

	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	/* cek apakah masih ada sisa di product itu */
	varian, _ := controller.VarianRepository.FindById(r.Context(), dataRequest.TempCartProductVarianId)
	if varian.ProductVarianQtyBooth == 0 {
		panic(exceptions.NewBadRequestError("Tidak dapat menambah jumlah lagi karena sudah habis dipesan"))
	}
	/* cek apakah data sudah ada sebelumnya */
	tempCart := controller.TempCartRepository.FindSpesificData(r.Context(), entity.TempCart{
		TempCartProductVarianId: dataRequest.TempCartProductVarianId,
		TempCartOrderId:         dataRequest.TempCartOrderId,
	})

	if tempCart != nil {
		controller.TempCartRepository.Delete(r.Context(), tempCart[0].TempCartId)
		dataRequest.TempCartQty = dataRequest.TempCartQty + tempCart[0].TempCartQty
	}
	dataRequest.TempCartOrderId = dataRequest.TempCartOrderId
	controller.TempCartRepository.Create(r.Context(), dataRequest)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    varian,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productVarianId := vars["productVarianId"]
	userId := vars["userId"]
	dataRequest := entity.TempCart{}
	helpers.ReadFromRequestBody(r, &dataRequest)
	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	/* ambil data temp cart */
	cart := controller.TempCartRepository.FindSpesificData(r.Context(), entity.TempCart{
		TempCartProductVarianId: productVarianId,
		TempCartUserId:          userId,
	})

	if cart == nil {
		panic(exceptions.NewNotFoundError("Data tidak ditemukan"))
	}

	productVarian, _ := controller.VarianRepository.FindById(r.Context(), productVarianId)
	/* check apakah masih ada quantity tersisa */
	if (productVarian.ProductVarianQtyBooth + cart[0].TempCartQty) < dataRequest.TempCartQty {
		panic(exceptions.NewBadRequestError("Tidak dapat mengubah data karena jumlah tersisa kurang"))
	}

	/* update data temporary */
	dataRequest.TempCartId = cart[0].TempCartId
	dataRequest.TempCartProductVarianId = cart[0].TempCartProductVarianId
	dataRequest.TempCartUserId = cart[0].TempCartUserId
	dataRequest.TempCartOrderId = cart[0].TempCartOrderId
	controller.TempCartRepository.Update(r.Context(), dataRequest, cart[0].TempCartId)
	//controller.TempCartRepository.Delete(r.Context(), cart[0].TempCartId)

	/* ambil lagi data temporary terbaru */
	data, _ := controller.VarianRepository.FindById(r.Context(), productVarianId)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    data,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	producVarianId := vars["productVarianId"]
	userId := vars["userId"]

	data := controller.TempCartRepository.FindSpesificData(r.Context(), entity.TempCart{
		TempCartProductVarianId: producVarianId,
		TempCartUserId:          userId,
	})

	if data != nil {
		controller.TempCartRepository.Delete(r.Context(), data[0].TempCartId)
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		panic(exceptions.NewBadRequestError("Parameter userId diperlukan"))
	}

	dataResponse := controller.TempCartRepository.FindByUserId(r.Context(), userId)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) ClearCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		panic(exceptions.NewBadRequestError("Parameter userId diperlukan"))
	}

	controller.TempCartRepository.DeleteByUserId(r.Context(), userId)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
