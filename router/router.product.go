package router

import (
	"github.com/edwinjordan/e-canteen-backend/app/service"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_product"
	"github.com/edwinjordan/e-canteen-backend/config"
	"github.com/edwinjordan/e-canteen-backend/repository/product_repository"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// ProductRouter sets up product routes
func ProductRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	productRepository := product_repository.New(db)

	// Initialize MinIO
	minioClient := config.NewMinioClient()
	minioService := service.NewMinioService(minioClient)

	productController := usecase_product.NewUseCase(productRepository, minioService, validate)

	// @Summary Get all products
	// @Description Retrieve all products with their variants
	// @Tags Products
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {object} handler.WebResponse{data=[]entity.Product} "Products list"
	// @Router /products [get]
	router.HandleFunc("/api/products", productController.FindAll).Methods("GET")

	// @Summary Create a new product with variants
	// @Description Create a new product along with its variants
	// @Tags Products
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Param product body usecase_product.ProductInsertRequest true "Product data with variants"
	// @Success 200 {object} handler.WebResponse{data=map[string]interface{}} "Product created successfully"
	// @Router /products [post]
	router.HandleFunc("/api/products", productController.Insert).Methods("POST")

	// @Summary Get product by ID
	// @Description Retrieve a specific product by its ID with variants
	// @Tags Products
	// @Produce json
	// @Security BearerAuth
	// @Param productId path string true "Product ID"
	// @Success 200 {object} handler.WebResponse{data=entity.Product} "Product detail"
	// @Router /products/{productId} [get]
	router.HandleFunc("/api/products/{productId}", productController.FindById).Methods("GET")

	// @Summary Update a product
	// @Description Update a product's information
	// @Tags Products
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Param productId path string true "Product ID"
	// @Param product body usecase_product.ProductUpdateRequest true "Product update data"
	// @Success 200 {object} handler.WebResponse{data=map[string]interface{}} "Product updated successfully"
	// @Router /products/{productId} [put]
	router.HandleFunc("/api/products/{productId}", productController.Update).Methods("PUT")

	// @Summary Delete a product
	// @Description Soft delete a product by its ID
	// @Tags Products
	// @Produce json
	// @Security BearerAuth
	// @Param productId path string true "Product ID"
	// @Success 200 {object} handler.WebResponse{data=map[string]interface{}} "Product deleted successfully"
	// @Router /products/{productId} [delete]
	router.HandleFunc("/api/products/{productId}", productController.Delete).Methods("DELETE")
}
