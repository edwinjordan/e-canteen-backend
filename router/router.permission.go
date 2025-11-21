package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_permission"
	"github.com/edwinjordan/e-canteen-backend/repository/permission_repository"
	"gorm.io/gorm"
)

func PermissionRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	permissionRepo := permission_repository.New(db)
	permissionRoleRepo := permission_repository.NewPermissionRole(db)
	permissionController := usecase_permission.NewUseCase(permissionRepo, permissionRoleRepo, validate)

	// Permission CRUD
	router.HandleFunc("/api/permission", permissionController.Create).Methods("POST")
	router.HandleFunc("/api/permission", permissionController.FindAll).Methods("GET")
	router.HandleFunc("/api/permission/{permissionId}", permissionController.FindById).Methods("GET")
	router.HandleFunc("/api/permission/{permissionId}", permissionController.Update).Methods("PUT")
	router.HandleFunc("/api/permission/{permissionId}", permissionController.Delete).Methods("DELETE")

	// Permission Role Assignment
	router.HandleFunc("/api/permission/role/{roleId}", permissionController.FindByRole).Methods("GET")
	router.HandleFunc("/api/permission/assign", permissionController.AssignToRole).Methods("POST")
	router.HandleFunc("/api/permission/revoke", permissionController.RevokeFromRole).Methods("DELETE")
}
