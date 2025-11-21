package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_tempcart"
	"github.com/edwinjordan/e-canteen-backend/repository/tempcart_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/varian_repository"
	"gorm.io/gorm"
)

// TempCartRouter sets up temporary cart routes
func TempCartRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	tempCartRepo := tempcart_repository.New(db)
	varianRepository := varian_repository.New(db)
	tempCartController := usecase_tempcart.NewUseCase(tempCartRepo, varianRepository, validate)

	// @Summary Get cart items by user ID
	// @Description Get all temporary cart items for a specific user
	// @Tags Cart
	// @Produce json
	// @Security BearerAuth
	// @Param userId path string true "User ID"
	// @Success 200 {object} handler.WebResponse{data=[]entity.TempCart} "Cart items"
	// @Failure 400 {object} handler.WebResponse "Invalid request"
	// @Router /tempcart/{userId} [get]
	router.HandleFunc("/api/tempcart/{userId}", tempCartController.FindByUserId).Methods("GET")

	// @Summary Add item to cart
	// @Description Add product variant to temporary cart
	// @Tags Cart
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Param request body entity.TempCart true "Cart item"
	// @Success 201 {object} handler.WebResponse{data=entity.TempCart} "Item added to cart"
	// @Failure 400 {object} handler.WebResponse "Invalid request"
	// @Router /tempcart [post]
	router.HandleFunc("/api/tempcart", tempCartController.Create).Methods("POST")

	// @Summary Update cart item quantity
	// @Description Update quantity of item in temporary cart
	// @Tags Cart
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Param productVarianId path string true "Product Variant ID"
	// @Param userId path string true "User ID"
	// @Success 200 {object} handler.WebResponse "Cart item updated"
	// @Failure 404 {object} handler.WebResponse "Cart item not found"
	// @Router /tempcart/{productVarianId}/{userId} [put]
	router.HandleFunc("/api/tempcart/{productVarianId}/{userId}", tempCartController.Update).Methods("PUT")

	// @Summary Remove item from cart
	// @Description Remove product variant from temporary cart
	// @Tags Cart
	// @Produce json
	// @Security BearerAuth
	// @Param productVarianId path string true "Product Variant ID"
	// @Param userId path string true "User ID"
	// @Success 200 {object} handler.WebResponse "Item removed from cart"
	// @Failure 404 {object} handler.WebResponse "Cart item not found"
	// @Router /tempcart/{productVarianId}/{userId} [delete]
	router.HandleFunc("/api/tempcart/{productVarianId}/{userId}", tempCartController.Delete).Methods("DELETE")
}
