package irc

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/mix3/go-ikusan/args"
	"github.com/mix3/go-ikusan/channel"
	"github.com/mix3/go-irc"
)

var conn *MyConn

type MyConn struct {
	*irc.Conn
}

func callbackFunc(conn *irc.Conn, e *irc.Event) {
	irc.DefaultCallback(conn, e)
	switch e.Code {
	case "001":
		for _, k := range channel.List() {
			v, _ := channel.Get(k)
			if v.ChannelKeyword != "" {
				conn.Join(k + " " + v.ChannelKeyword)
			} else {
				conn.Join(k)
			}
		}
	}
}

func create(config *args.Result) *MyConn {
	cfg := &irc.Config{
		Nick:     config.IrcNickname(),
		User:     config.IrcUser(),
		SSL:      config.EnableSsl(),
		Interval: time.Duration(config.IrcPostInterval()) * time.Second,
	}

	if config.EnableSsl() && config.InsecureSkipVerify() {
		cfg.SSLConfig = &tls.Config{InsecureSkipVerify: true}
	}

	c, err := irc.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	c.Callbacker(callbackFunc)

	return &MyConn{c}
}

func Run(config *args.Result) {
	conn = create(config)

	quit, err := conn.Connect(
		config.IrcServer(),
		config.IrcPort(),
		config.IrcKeyword(),
	)

	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-quit:
			conn.Logger().Debugf("[INFO   ] quit")
			for {
				conn.Logger().Infof("[INFO   ] reconnecting")
				quit, err = conn.Reconnect()
				if err == nil {
					conn.Logger().Infof("[INFO   ] reconnected")
					break
				}
				conn.Logger().Warnf("[ERROR  ] fail reconnection")
				time.Sleep(time.Duration(config.IrcReconnectInterval()) * time.Second)
			}
		}
	}
}

func Conn() (*MyConn, error) {
	if conn == nil {
		return nil, fmt.Errorf("conn is null")
	}
	if !conn.IsConnected() {
		return nil, fmt.Errorf("disconnected")
	}
	return conn, nil
}

func (mc *MyConn) Join(channelName string, option ...string) {
	keyword := ""
	if 0 < len(option) {
		keyword = option[0]
	}
	if keyword != "" {
		mc.Conn.Join(channelName + " " + keyword)
	} else {
		mc.Conn.Join(channelName)
	}
	channel.Set(channelName, keyword)
}

func (mc *MyConn) Part(channelName string) {
	mc.Conn.Part(channelName)
	channel.Del(channelName)
}
