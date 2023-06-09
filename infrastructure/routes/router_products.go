package infrastructure

import (
	"net/http"

	v1 "backend_crudgo/domain/products/domain/handler/v1"
	"backend_crudgo/infrastructure/database"
	"backend_crudgo/infrastructure/middlewares"

	"github.com/go-chi/chi"
)

// RoutesProducts aa
func RoutesProducts(conn *database.DataDB) http.Handler {
	router := chi.NewRouter()
	products := v1.NewProductHandler(conn) //domain
	router.Mount("/products", routesProduct(products))
	return router
}

// Router user
func routesProduct(handler *v1.ProductRouter) http.Handler {
	router := chi.NewRouter()
	router.Use(middlewares.AuthMiddleware)
	router.Post("/", handler.CreateProductHandler)
	router.Get("/", handler.GetProductsHandler)
	router.Get("/{id}", handler.GetProductHandler)
	router.Put("/{id}", handler.UpdateProductHandler)
	router.Delete("/{id}", handler.DeleteProductHandler)
	return router
}
