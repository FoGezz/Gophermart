package handler

import (
	"Gophermart/cmd/gophermart/config"
	"Gophermart/internal/app/domain/service"
	"Gophermart/internal/app/middleware"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
)

func register(app *config.App, usrService *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := new(service.AuthRequest)
		err := json.NewDecoder(r.Body).Decode(request)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			app.Logger.Warn("Incorrect request from post_register", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userId, err := usrService.RegisterNewUser(r.Context(), request)
		if err != nil {
			if errors.Is(err, service.ErrUserAlreadyExists) {
				w.WriteHeader(http.StatusConflict)
				return
			}
			app.Logger.Error("Unexpected error while registering user", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, middleware.GophermartClaims{
			RegisteredClaims: jwt.RegisteredClaims{},
			UserID:           userId,
		})
		strJwt, _ := newToken.SignedString([]byte(middleware.JwtSecretKey))

		http.SetCookie(w, &http.Cookie{
			Name:  "Token",
			Value: strJwt,
		})
		w.WriteHeader(http.StatusOK)
	}
}
