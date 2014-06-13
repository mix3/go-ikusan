package server

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/codegangsta/negroni"
)

type Test struct {
	HeaderKey  string
	RemoteAddr string
	HeaderAddr string
	ExpectAddr string
}

var tests = []Test{
	Test{"X-Real-IP", "127.0.0.1:19301", "192.168.0.1", "192.168.0.1"},
	Test{"X-Forwarded-For", "127.0.0.2:19302", "192.168.0.2", "192.168.0.2"},
	Test{"X-Forwarded-For", "127.0.0.3:19303", "192.168.0.3", "127.0.0.3:19303"},
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestReverseProxy(t *testing.T) {
	for _, test := range tests {
		header := http.Header{}
		header.Set(test.HeaderKey, test.HeaderAddr)
		req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
		req.RemoteAddr = test.RemoteAddr
		req.Header = header
		requestTest(t, req, test)
	}
}

func requestTest(t *testing.T, r *http.Request, test Test) {
	n := negroni.New()
	n.Use(NewReverseProxy([]string{"127.0.0.1", "127.0.0.2"}...))
	n.UseHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		expect(t, r.RemoteAddr, test.ExpectAddr)
	}))
	n.ServeHTTP(httptest.NewRecorder(), r)
}
