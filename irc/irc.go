package irc

import (
	"crypto/tls"
	"time"

	"github.com/mix3/go-ikusan/args"
	"github.com/mix3/go-irc"
)

var conn *Conn

type Channel struct {
	ChannelKeyword string
	JoinAt         time.Time
}

type Conn struct {
	*irc.Conn
	joinChannels map[string]Channel
	quit         chan struct{}
}

func (conn *Conn) Callback(e *irc.Event) {
	conn.DefaultCallback(e)
	switch e.Code {
	case "001":
		conn_ := GetConn()
		for channel, channelInfo := range conn_.joinChannels {
			conn_.Join(channel, channelInfo.ChannelKeyword)
		}
	}
}

func Init(config *args.Result) error {
	cfg := &irc.Config{
		Nick:     config.IrcNickname(),
		User:     config.IrcUser(),
		SSL:      config.EnableSsl(),
		Interval: time.Duration(config.IrcPostInterval()) * time.Second,
	}

	if config.EnableSsl() && config.InsecureSkipVerify() {
		cfg.SSLConfig = &tls.Config{InsecureSkipVerify: true}
	}

	ircconn, err := irc.New(cfg)
	if err != nil {
		return err
	}
	conn = &Conn{
		ircconn,
		make(map[string]Channel),
		nil,
	}
	conn.SetEmbed(conn)

	var quit chan struct{}
	quit, err = conn.Connect(
		config.IrcServer(),
		config.IrcPort(),
		config.IrcKeyword(),
	)
	if err != nil {
		return err
	}

	conn.quit = quit

	return nil
}

func GetConn() *Conn {
	return conn
}

func (conn *Conn) GetQuitChan() chan struct{} {
	return conn.quit
}

func (conn *Conn) IsJoined(channel string) bool {
	_, ok := conn.joinChannels[channel]
	return ok
}

func (conn *Conn) Join(channel string, option ...string) {
	keyword := ""
	if 0 < len(option) {
		keyword = option[0]
	}
	if keyword != "" {
		conn.Conn.Join(channel + " " + keyword)
	} else {
		conn.Conn.Join(channel)
	}
	conn.joinChannels[channel] = Channel{keyword, time.Now()}
}

func (conn *Conn) Part(channel string) {
	conn.Conn.Part(channel)
	delete(conn.joinChannels, channel)
}

func (conn *Conn) ChannelList() []string {
	list := []string{}
	for k, _ := range conn.joinChannels {
		list = append(list, k)
	}
	return list
}
