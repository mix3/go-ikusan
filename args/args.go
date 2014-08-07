package args

import (
	"flag"
	"fmt"
	"strings"
)

type ArrayString []string

func (a *ArrayString) String() string {
	return fmt.Sprint(*a)
}

func (a *ArrayString) Set(v string) error {
	*a = append(*a, v)
	return nil
}

type Result struct {
	httpHost             string
	httpPort             uint
	reverseProxy         ArrayString
	mount                string
	ircServer            string
	ircPort              uint
	ircKeyword           string
	ircNickname          string
	ircUser              string
	ircPostInterval      uint
	enableSsl            bool
	insecureSkipVerify   bool
	ircReconnectInterval uint
	noPostWithJoin       bool
	maxLength            uint
	help                 bool
}

func (a Result) HttpHost() string {
	return a.httpHost
}

func (a Result) HttpPort() uint {
	return a.httpPort
}

func (a Result) ReverseProxy() []string {
	return a.reverseProxy
}

func (a Result) Mount() string {
	return a.mount
}

func (a Result) IrcServer() string {
	return a.ircServer
}

func (a Result) IrcPort() uint {
	return a.ircPort
}

func (a Result) IrcKeyword() string {
	return a.ircKeyword
}

func (a Result) IrcNickname() string {
	return a.ircNickname
}

func (a Result) IrcUser() string {
	return a.ircUser
}

func (a Result) IrcPostInterval() uint {
	return a.ircPostInterval
}

func (a Result) EnableSsl() bool {
	return a.enableSsl
}

func (a Result) InsecureSkipVerify() bool {
	return a.insecureSkipVerify
}

func (a Result) IrcReconnectInterval() uint {
	return a.ircReconnectInterval
}

func (a Result) NoPostWithJoin() bool {
	return a.noPostWithJoin
}

func (a Result) MaxLength() uint {
	return a.maxLength
}

func (a Result) Help() bool {
	return a.help
}

func (a Result) IsValid() bool {
	return a.IrcServer() != ""
}

const (
	FLAG_SET_NAME = "ikusan/args"
)

type FlagSet struct {
	*flag.FlagSet
}

func newFlagSet(name string, errorHandling flag.ErrorHandling) *FlagSet {
	return &FlagSet{flag.NewFlagSet(name, errorHandling)}
}

func (f *FlagSet) boolVar(p *bool, name string, value bool, usage string) {
	for _, n := range strings.Split(name, "|") {
		f.BoolVar(p, n, value, usage)
	}
}

func (f *FlagSet) stringVar(p *string, name string, value string, usage string) {
	for _, n := range strings.Split(name, "|") {
		f.StringVar(p, n, value, usage)
	}
}

func (f *FlagSet) uintVar(p *uint, name string, value uint, usage string) {
	for _, n := range strings.Split(name, "|") {
		f.UintVar(p, n, value, usage)
	}
}

func (f *FlagSet) varVar(value flag.Value, name string, usage string) {
	for _, n := range strings.Split(name, "|") {
		f.Var(value, n, usage)
	}
}

var config *Result

func Parse(Args []string) {
	a := &Result{
		reverseProxy: []string{},
	}
	f := newFlagSet(FLAG_SET_NAME, flag.ExitOnError)
	f.stringVar(&a.httpHost, "o|host", "127.0.0.1", "")
	f.uintVar(&a.httpPort, "p|port", 19300, "")
	f.varVar(&a.reverseProxy, "r|reverse-proxy", "")
	f.stringVar(&a.mount, "m|mount", "", "")
	f.stringVar(&a.ircServer, "S|Server", "", "")
	f.uintVar(&a.ircPort, "P|Port", 6667, "")
	f.stringVar(&a.ircKeyword, "K|Keyword", "", "")
	f.stringVar(&a.ircNickname, "N|Nickname", "ikusan", "")
	f.stringVar(&a.ircUser, "U|User", "ikusan", "")
	f.uintVar(&a.ircPostInterval, "i|interval", 2, "")
	f.boolVar(&a.enableSsl, "enable-ssl", false, "")
	f.boolVar(&a.insecureSkipVerify, "insecure-skip-verify", false, "")
	f.uintVar(&a.ircReconnectInterval, "R|reconnect-interval", 3, "")
	f.boolVar(&a.noPostWithJoin, "j|no-post-with-join", false, "")
	f.uintVar(&a.maxLength, "l|max-length", 0, "")
	f.boolVar(&a.help, "h|help", false, "")
	f.Parse(Args)
	config = a
}

func GetConfig() *Result {
	if config != nil {
		return config
	}
	return nil
}

func Usage() string {
	return `Usage:
	# connect to chat.freenode.net
	ikusan -S chat.freenode.net
`
}

func Help() string {
	return `	-o, --host
		The interface a TCP based server daemon binds to. Defauts to undef,
		which lets most server backends bind the any (*) interface. This
		option doesn't mean anything if the server does not support TCP
		socket.
	-p, --port (default: 19300)
		The port number a TCP based server daemon listens on. Defaults to
		19300. This option doesn't mean anything if the server does not support
		TCP socket.
	-S, --Server
		irc server address.
	-P, --Port (default: 6667)
		irc server port.
	--enable-ssl
		use ssl connection.
	--insecure-skip-verify
		skip verify if use ssl.
	-K, --Keyword
		irc server password
	-N, --Nickname
		irc nickname
	-U, --User
		irc user name
	-r, --reverse-proxy
		treat X-Forwarded-For as REMOTE_ADDR if REMOTE_ADDR match this argument.
	-i, --interval
		irc post interval. for Excess Flood
	-R, --reconnect-interval
		interval of reconnect to irc server.
		exit application if interval == 0.
	-j, --no-post-with-join
		disable to irc message post with channel join
	-m, --mount
		provide TCP based server daemon with path. Default do nothing.
	-l, --max-length
		truncate message after a given length. Do not truncate by default.
`
}
