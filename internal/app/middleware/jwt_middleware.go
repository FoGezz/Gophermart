package middleware

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type GophermartClaims struct {
	jwt.RegisteredClaims
	UserID int
}

type CookieKey string

const JwtSecretKey string = "supersecretkeyforgophermart"
const JwtUserIDKey CookieKey = "UserId"

func JwtAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/user/register" || r.RequestURI == "/api/user/login" {
			next.ServeHTTP(w, r)
			return
		}

		token, err := r.Cookie("Token")
		if errors.Is(http.ErrNoCookie, err) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := &GophermartClaims{}

		_, err = jwt.ParseWithClaims(token.Value, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(JwtSecretKey), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), JwtUserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
