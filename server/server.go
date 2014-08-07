package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mix3/go-ikusan/args"
	ch "github.com/mix3/go-ikusan/channel"
	"github.com/mix3/go-ikusan/irc"
)

func New(config *args.Result) *negroni.Negroni {
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.Use(NewReverseProxy(config.ReverseProxy()...))
	n.Use(NewCheckConnection())
	n.UseHandler(NewRouter(config.Mount()))
	return n
}

func NewRouter(mount string) http.Handler {
	router := mux.NewRouter()
	if mount != "" {
		router.HandleFunc("/"+mount, IndexHandler).Methods("GET")
	}
	subrouter := router.PathPrefix("/" + mount).Subrouter()
	subrouter.HandleFunc("/", IndexHandler).Methods("GET")
	subrouter.HandleFunc("/channel_list", ChannelListHandler).Methods("GET")
	subrouter.HandleFunc("/join", JoinHandler).Methods("POST")
	subrouter.HandleFunc("/leave", LeaveHandler).Methods("POST")
	subrouter.HandleFunc("/part", PartHandler).Methods("POST")
	subrouter.HandleFunc("/notice", NoticeHandler).Methods("POST")
	subrouter.HandleFunc("/privmsg", PrivmsgHandler).Methods("POST")
	return router
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	Config := args.GetConfig()
	mount := Config.Mount()
	path := ""
	if mount != "" {
		path = "/" + mount
	}
	indexTemplate.Execute(w, path)
}

func ChannelListHandler(w http.ResponseWriter, req *http.Request) {
	message := strings.Join(ch.List(), "\n") + "\n"
	render(w, 200, message)
}

func JoinHandler(w http.ResponseWriter, req *http.Request) {
	conn, _ := irc.Conn()
	channel := req.FormValue("channel")
	channelKeyword := req.FormValue("channel_keyword")
	if _, ok := ch.Get(channel); ok {
		render(w, 403, "joinned channel: %s\n", channel)
		return
	}
	conn.Join(channel, channelKeyword)
	render(w, 200, "join success channel: %s\n", channel)
}

func LeaveHandler(w http.ResponseWriter, req *http.Request) {
	PartHandler(w, req)
}

func PartHandler(w http.ResponseWriter, req *http.Request) {
	conn, _ := irc.Conn()
	channel := req.FormValue("channel")
	if _, ok := ch.Get(channel); !ok {
		render(w, 404, "not joinned channel: %s\n", channel)
		return
	}
	conn.Part(channel)
	render(w, 200, "leave success channel: %s\n", channel)
}

func NoticeHandler(w http.ResponseWriter, req *http.Request) {
	conn, _ := irc.Conn()
	channel := req.FormValue("channel")
	config := args.GetConfig()
	if config.NoPostWithJoin() {
		if _, ok := ch.Get(channel); !ok {
			render(w, 404, "not joinned channel: %s\n", channel)
			return
		}
	} else {
		conn.Join(channel)
	}
	message := truncateMessage(req.FormValue("message"), config.MaxLength())
	conn.Notice(channel, message)
	render(w, 200, "message sent channel: %s %s\n", channel, message)
}

func PrivmsgHandler(w http.ResponseWriter, req *http.Request) {
	conn, _ := irc.Conn()
	channel := req.FormValue("channel")
	config := args.GetConfig()
	if config.NoPostWithJoin() {
		if _, ok := ch.Get(channel); !ok {
			render(w, 404, "not joinned channel: %s\n", channel)
			return
		}
	} else {
		conn.Join(channel)
	}
	message := truncateMessage(req.FormValue("message"), config.MaxLength())
	conn.Privmsg(channel, message)
	render(w, 200, "message sent channel: %s %s\n", channel, message)
}

func render(w http.ResponseWriter, status int, format string, data ...interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, format, data...)
}

func truncateMessage(message string, maxLength uint) string {
	l := int(maxLength)
	if l <= 0 {
		return message
	}
	runes := []rune(message)
	if len(runes) < l {
		l = len(runes)
	}
	return string(runes[:l])
}
