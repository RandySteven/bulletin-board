package middlewares

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"task_mission/enums"
	"task_mission/pkg/securities"
	"task_mission/utils"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.ContentType(w, "application/json")
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 || auth == "" {
			utils.ResponseHandler(w, http.StatusUnauthorized, `Invalid get token from auth`, nil, nil, nil)
			return
		}
		tokenStr := auth[len("Bearer "):]
		if tokenStr == "" {
			utils.ResponseHandler(w, http.StatusUnauthorized, `Invalid token failed to split from bearer`, nil, nil, nil)
			return
		}
		claims := &securities.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(j *jwt.Token) (interface{}, error) {
			return securities.JWT_KEY, nil
		})
		if err != nil || !token.Valid {
			utils.ResponseHandler(w, http.StatusUnauthorized, `Invalid token`, nil, nil, err)
			return
		}
		ctx := context.WithValue(r.Context(), enums.UserID, claims.UserID)
		ctx2 := context.WithValue(ctx, enums.RoleID, claims.RoleID)
		r = r.WithContext(ctx2)
		log.Println(r.Context())
		next.ServeHTTP(w, r)
	})
}
