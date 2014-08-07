package server

import (
	"fmt"
	"net/http"

	"github.com/mix3/go-ikusan/irc"
)

type checkConnection struct{}

func NewCheckConnection() *checkConnection {
	return &checkConnection{}
}

func (cc *checkConnection) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	conn, err := irc.Conn()
	if r.Method == "POST" && (err != nil || !conn.IsConnected()) {
		rw.WriteHeader(503)
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(rw, `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>ikusan</title>
<head>
<body>
can not connect to irc server
</body>
</html>
`)
		return
	}
	next(rw, r)
}
