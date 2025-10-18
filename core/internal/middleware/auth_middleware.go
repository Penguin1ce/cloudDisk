// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"cloudDisk/core/utils"
	"log"
	"net/http"
	"strconv"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		uc, err := utils.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		log.Printf("获取到的认证信息 %d", uc.Id)
		r.Header.Set("UserId", strconv.Itoa(uc.Id))
		r.Header.Set("UserName", uc.Name)
		r.Header.Set("UserIdentity", uc.Identity)
		next(w, r)
	}
}
