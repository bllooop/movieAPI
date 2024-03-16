package handlers

import (
	"fmt"
	movieapi "movieAPI"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

/*
	func (h *Handler) shopIdentity(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			newError(c, http.StatusUnauthorized, "empty auth header")
			return
		}
		headerSplit := strings.Split(header, " ")
		if len(headerSplit) != 2 {
			newError(c, http.StatusUnauthorized, "invalid auth header")
			return
		}
		shopId, err := h.services.Authorization.ParseToken(headerSplit[1])
		if err != nil {
			newError(c, http.StatusUnauthorized, err.Error())
		}
		c.Set(shopCtx, shopId)
	}
*/
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}
		headerSplit := strings.Split(tokenString, " ")
		if len(headerSplit) != 2 {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(headerSplit[1], func(token *jwt.Token) (interface{}, error) {
			// Check signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func CreateToken(user movieapi.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*func (h *Handler) getRole(userId int) (int, error) {
	id, err := h.services.Authorization.GetUserRole(userId)
	if err != nil {
		return 0, nil
	}
} */
