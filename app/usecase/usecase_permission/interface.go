package usecase_permission

import "net/http"

type UseCase interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindByRole(w http.ResponseWriter, r *http.Request)
	AssignToRole(w http.ResponseWriter, r *http.Request)
	RevokeFromRole(w http.ResponseWriter, r *http.Request)
}
