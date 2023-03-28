package infrastructure

import (
	"net/http"

	v1 "backend_crudgo/domain/users/domain/handler/v1"
	"backend_crudgo/infrastructure/database"
	"backend_crudgo/infrastructure/kit/enum"

	"github.com/go-chi/chi"
)

// RoutesProducts aa
func RoutesUsers(conn *database.DataDB) http.Handler {
	router := chi.NewRouter()
	users := v1.NewUserHandler(conn) //domain
	router.Mount("/", routesUser(users))
	return router
}

// Router user
func routesUser(handler *v1.UserRouter) http.Handler {
	router := chi.NewRouter()
	router.Post(enum.LoginUserPath, handler.GetUserHandler)
	router.Post(enum.RegisterPath, handler.CreateUserHandler)

	return router
}
