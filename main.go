package main

import (
	"net/http"

	"github.com/edwinjordan/e-canteen-backend/config"
	_ "github.com/edwinjordan/e-canteen-backend/docs"
	"github.com/edwinjordan/e-canteen-backend/middleware"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"github.com/edwinjordan/e-canteen-backend/pkg/mysql"
	"github.com/edwinjordan/e-canteen-backend/router"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title E-Canteen Cashier API
// @version 1.2.0
// @description API documentation for E-Canteen Cashier System - A comprehensive canteen management solution

// @contact.name API Support
// @contact.email support@ecanteen.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host 127.0.0.1:3000
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	route := mux.NewRouter()

	/* setting cors */
	corsOpt := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPut,
		},
		AllowedHeaders: []string{
			"*",
		},
	})
	/* load middleware */
	// route.Use(middleware.Recovery)
	route.Use(middleware.Authentication)

	/* load router */
	router.KasirRouter(db, validate, route)
	router.VarianRouter(db, validate, route)
	router.TempCartRouter(db, validate, route)
	router.ProductRouter(db, validate, route)
	router.CategoryRouter(db, validate, route)
	router.CustomerRouter(db, validate, route)
	router.CustomerAddressRouter(db, validate, route)
	router.MajorRouter(db, validate, route)
	router.TerritoryRouter(db, validate, route)
	router.OrderRouter(db, validate, route)
	router.DashboardRouter(db, route)
	router.PermissionRouter(db, validate, route)
	router.DashboardCustomerRouter(db, route)

	/* Swagger documentation */
	route.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://127.0.0.1:3000/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	server := http.Server{
		Addr:    config.GetEnv("HOST_ADDR"),
		Handler: corsOpt.Handler(route),
	}
	err := server.ListenAndServe()
	helpers.PanicIfError(err)

}
