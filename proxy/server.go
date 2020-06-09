package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/xlzd/gotp"
)

const (
	tokenKey      = "x-totp-token"
	defaultDigits = 6
)

type Server struct {
	secret   string
	upstream *url.URL
	interval int64
	totp     *gotp.TOTP
	proxy    *httputil.ReverseProxy
}

func NewServer(upstream *url.URL, secret string, interval int64) *Server {
	proxy := httputil.NewSingleHostReverseProxy(upstream)
	totp := gotp.NewTOTP(secret, defaultDigits, int(interval), nil) // refer: gotp.NewDefaultTOTP
	return &Server{
		secret:   secret,
		upstream: upstream,
		interval: interval,
		proxy:    proxy,
		totp:     totp,
	}
}

func (s *Server) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	token := s.token(r)
	if token == "" {
		wr.WriteHeader(http.StatusUnauthorized)
		log.Println("reject empty token")
		return
	}

	s.setCookie(wr, r, token)

	if !s.totp.Verify(token, int(time.Now().Unix())) {
		wr.WriteHeader(http.StatusUnauthorized)
		log.Println("reject wrong token")
		return
	}

	r.URL.Host = s.upstream.Host
	r.URL.Scheme = s.upstream.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = s.upstream.Host

	s.proxy.ServeHTTP(wr, r)
}
