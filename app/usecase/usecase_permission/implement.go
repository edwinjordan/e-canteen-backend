package usecase_permission

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
	PermissionRepository     repository.PermissionRepository
	PermissionRoleRepository repository.PermissionRoleRepository
	Validate                 *validator.Validate
}

func NewUseCase(
	permissionRepo repository.PermissionRepository,
	permissionRoleRepo repository.PermissionRoleRepository,
	validate *validator.Validate,
) UseCase {
	return &UseCaseImpl{
		PermissionRepository:     permissionRepo,
		PermissionRoleRepository: permissionRoleRepo,
		Validate:                 validate,
	}
}

func (uc *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	request := entity.CreatePermissionRequest{}
	helpers.ReadFromRequestBody(r, &request)

	err := uc.Validate.Struct(request)
	helpers.PanicIfError(err)

	permission := entity.Permission{
		PermissionName:        request.PermissionName,
		PermissionResource:    request.PermissionResource,
		PermissionAction:      request.PermissionAction,
		PermissionDescription: request.PermissionDescription,
	}

	result := uc.PermissionRepository.Create(r.Context(), permission)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    result,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (uc *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["permissionId"])

	request := entity.UpdatePermissionRequest{}
	helpers.ReadFromRequestBody(r, &request)

	// Check if permission exists
	existing, err := uc.PermissionRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError("Permission tidak ditemukan"))
	}

	// Update only provided fields
	if request.PermissionName != "" {
		existing.PermissionName = request.PermissionName
	}
	if request.PermissionResource != "" {
		existing.PermissionResource = request.PermissionResource
	}
	if request.PermissionAction != "" {
		existing.PermissionAction = request.PermissionAction
	}
	if request.PermissionDescription != "" {
		existing.PermissionDescription = request.PermissionDescription
	}

	result := uc.PermissionRepository.Update(r.Context(), *existing)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    result,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (uc *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["permissionId"])

	// Check if exists
	_, err := uc.PermissionRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError("Permission tidak ditemukan"))
	}

	err = uc.PermissionRepository.Delete(r.Context(), id)
	helpers.PanicIfError(err)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (uc *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["permissionId"])

	result, err := uc.PermissionRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError("Permission tidak ditemukan"))
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    result,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (uc *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	result := uc.PermissionRepository.FindAll(r.Context())

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    result,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (uc *UseCaseImpl) FindByRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleId, _ := strconv.Atoi(vars["roleId"])

	result := uc.PermissionRepository.FindByRoleAsTree(r.Context(), roleId)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    result,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (uc *UseCaseImpl) AssignToRole(w http.ResponseWriter, r *http.Request) {
	request := entity.AssignPermissionRequest{}
	helpers.ReadFromRequestBody(r, &request)

	err := uc.Validate.Struct(request)
	helpers.PanicIfError(err)

	err = uc.PermissionRoleRepository.Assign(r.Context(), request.RoleId, request.PermissionIds)
	helpers.PanicIfError(err)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Permission berhasil di-assign ke role",
		Data:    nil,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (uc *UseCaseImpl) RevokeFromRole(w http.ResponseWriter, r *http.Request) {
	request := entity.RevokePermissionRequest{}
	helpers.ReadFromRequestBody(r, &request)

	err := uc.Validate.Struct(request)
	helpers.PanicIfError(err)

	err = uc.PermissionRoleRepository.Revoke(r.Context(), request.RoleId, request.PermissionId)
	helpers.PanicIfError(err)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Permission berhasil di-revoke dari role",
		Data:    nil,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
