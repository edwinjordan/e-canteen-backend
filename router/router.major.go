package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_major"
	"github.com/edwinjordan/e-canteen-backend/repository/major_repository"
	"gorm.io/gorm"
)

func MajorRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	majorRepository := major_repository.New(db)
	majorController := usecase_major.NewUseCase(majorRepository, validate)
	router.HandleFunc("/api/major/{majorId}", majorController.FindById).Methods("GET")
	router.HandleFunc("/api/major", majorController.FindAll).Methods("GET")
	router.HandleFunc("/api/major", majorController.Create).Methods("POST")
	router.HandleFunc("/api/major/{majorId}", majorController.Update).Methods("PUT")
	router.HandleFunc("/api/major/{majorId}", majorController.Delete).Methods("DELETE")
}
