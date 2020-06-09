package proxy

import (
	"net/http"
)

func (s *Server) token(r *http.Request) string {
	token := r.URL.Query().Get(tokenKey)
	if token == "" {
		cookie, err := r.Cookie(tokenKey)
		if err == nil {
			token = cookie.Value
		}
	}
	return token
}
