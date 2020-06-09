package proxy

import (
	"net/http"
	"time"
)

func (s *Server) setCookie(wr http.ResponseWriter, r *http.Request, token string) {
	expired := last(time.Now().Unix(), s.interval)
	cookie := http.Cookie{
		Name:     tokenKey,
		Value:    token,
		Domain:   r.URL.Host,
		Expires:  time.Unix(expired, 0),
		HttpOnly: true,
	}
	http.SetCookie(wr, &cookie)

}
