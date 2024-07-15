package handler

import (
	"Gophermart/cmd/gophermart/config"
	"Gophermart/internal/app/domain/service"
	"Gophermart/internal/app/middleware"
	"Gophermart/internal/app/repository"
	"github.com/go-chi/chi/v5"
)

func InitRoutes(mux *chi.Mux, app *config.App) {
	//middlewares
	mux.Use(middleware.JwtAuthorization)

	//repos
	userRepo := repository.NewUserRepository(app.DB)

	//services
	userService := service.NewUserService(userRepo, app.Logger)

	//routes
	mux.Post("/api/user/register", register(app, userService))
	mux.Post("/api/user/login", login(app, userService))
}
