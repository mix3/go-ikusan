package server

import (
	"net"
	"net/http"
)

type reverseProxy struct {
	ipMap map[string]struct{}
}

func NewReverseProxy(ips ...string) *reverseProxy {
	ipMap := map[string]struct{}{}
	for _, ip := range ips {
		ipMap[ip] = struct{}{}
	}
	return &reverseProxy{ipMap}
}

func (rp *reverseProxy) match(ip string) bool {
	_, ok := rp.ipMap[ip]
	return ok
}

func (rp *reverseProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && rp.match(ip) {
		addr := r.Header.Get("X-Real-IP")
		if addr == "" {
			addr = r.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = r.RemoteAddr
			}
		}
		r.RemoteAddr = addr
	}
	next(rw, r)
}
