package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_category"
	"github.com/edwinjordan/e-canteen-backend/repository/category_repository"
	"gorm.io/gorm"
)

func CategoryRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	categoryRepository := category_repository.New(db)
	categoryController := usecase_category.NewUseCase(categoryRepository, validate)
	router.HandleFunc("/api/category", categoryController.FindAll).Methods("GET")
	router.HandleFunc("/api/category", categoryController.Create).Methods("POST")
	router.HandleFunc("/api/category/{categoryId}", categoryController.Update).Methods("PUT")
	router.HandleFunc("/api/category/{categoryId}", categoryController.FindById).Methods("GET")
	router.HandleFunc("/api/category/{categoryId}", categoryController.Delete).Methods("DELETE")
}
