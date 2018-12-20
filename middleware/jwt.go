package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"todo/models"
	"todo/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

func badReq(w http.ResponseWriter) {
	res := utils.Message(false, "Invalid token")
	w.WriteHeader(http.StatusForbidden)
	utils.Respond(w, res)
}

// JwtMiddleware - parse token and add user id to context
func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			badReq(w)
			return
		}

		headerKV := strings.Split(header, " ")
		if len(headerKV) != 2 {
			badReq(w)
			return
		}

		tokenVal := headerKV[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenVal, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			badReq(w)
			return
		}

		if !token.Valid {
			badReq(w)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
