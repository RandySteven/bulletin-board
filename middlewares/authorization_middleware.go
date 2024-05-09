package middlewares

import (
	"net/http"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//roleMap := make(map[string][]string)
		//roleMap[enums.UserRouter]
		//utils.ContentType(w, "application/json")
		//auth := r.Header.Get("Authorization")
		//if len(auth) == 0 || auth == "" {
		//	utils.ResponseHandler(w, http.StatusUnauthorized, `Invalid get token from auth`, nil, nil, nil)
		//	return
		//}
		//tokenStr := auth[len("Bearer "):]
		//if tokenStr == "" {
		//	utils.ResponseHandler(w, http.StatusUnauthorized, `Invalid token failed to split from bearer`, nil, nil, nil)
		//	return
		//}
		//claims := &securities.JWTClaim{}
		//token, err := jwt.ParseWithClaims(tokenStr, claims, func(j *jwt.Token) (interface{}, error) {
		//	return securities.JWT_KEY, nil
		//})
		//if err != nil || !token.Valid {
		//	utils.ResponseHandler(w, http.StatusUnauthorized, `Invalid token`, nil, nil, err)
		//	return
		//}
		//ctx := context.WithValue(r.Context(), enums.UserID, claims.UserID)
		//ctx2 := context.WithValue(ctx, enums.RoleID, claims.RoleID)
		//r = r.WithContext(ctx2)
		//next.ServeHTTP(w, r)
	})
}
