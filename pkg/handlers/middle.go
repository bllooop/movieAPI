package handlers

import (
	"context"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	roleCtx             = "userRole"
)

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get(authorizationHeader)
		if tokenString == "" {
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}
		headerSplit := strings.Split(tokenString, " ")
		if len(headerSplit) != 2 {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}
		userRole, err := h.services.Authorization.ParseToken(headerSplit[1])
		if err != nil {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}
		r = setValue(r, userRole)
		next.ServeHTTP(w, r)
	})
}
func setValue(r *http.Request, userRole string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), roleCtx, userRole))
}

/*func (h *Handler) getRole(userId int) (int, error) {
	id, err := h.services.Authorization.GetUserRole(userId)
	if err != nil {
		return 0, nil
	}
} */
