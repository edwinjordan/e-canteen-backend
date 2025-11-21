package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_territory"
	"github.com/edwinjordan/e-canteen-backend/repository/territory_repository"
	"gorm.io/gorm"
)

func TerritoryRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	repositoryRepository := territory_repository.New(db)
	repositoryController := usecase_territory.NewUseCase(repositoryRepository, validate)
	router.HandleFunc("/api/province", repositoryController.GetProvince).Methods("GET")
	router.HandleFunc("/api/city", repositoryController.GetCity).Methods("GET")
	router.HandleFunc("/api/subdistrict", repositoryController.GetSubdistrict).Methods("GET")
	router.HandleFunc("/api/village", repositoryController.GetVillage).Methods("GET")
}
