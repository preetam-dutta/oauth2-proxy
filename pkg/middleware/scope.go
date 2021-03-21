package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/justinas/alice"
	middlewareapi "github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/middleware"
)

func NewScope(reverseProxy bool, idHeader string) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			scope := &middlewareapi.RequestScope{
				ReverseProxy: reverseProxy,
				RequestID:    genRequestID(req, idHeader),
			}
			req = middlewareapi.AddRequestScope(req, scope)
			next.ServeHTTP(rw, req)
		})
	}
}

func genRequestID(req *http.Request, idHeader string) string {
	rid := req.Header.Get(idHeader)
	if rid != "" {
		return rid
	}
	return uuid.New().String()
}
