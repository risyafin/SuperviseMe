package main

import (
	"context"
	"net/http"
	"superviseMe/core/entity"

	"github.com/golang-jwt/jwt"
)

func jwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tokenString := request.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, &entity.MyClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("Bolong"), nil
		})
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte(err.Error()))
			return
		}
		claims, ok := token.Claims.(*entity.MyClaims) //forward ke usecase create transaction
		if !ok || !token.Valid {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("token invalid"))
			return
		}
		contextClaims := request.Context()
		// contextClaimsUser := context.Background()

		contextClaims = context.WithValue(contextClaims, "gmail", claims.Email)
		contextClaims = context.WithValue(contextClaims, "userID", claims.Id)
		request = request.WithContext(contextClaims)
		next(writer, request)

	}
}
