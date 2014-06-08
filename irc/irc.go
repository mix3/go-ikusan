package irc

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/mix3/go-ikusan/args"
	ircevent "github.com/thoj/go-ircevent"
)

var conn *Connection

type Channel struct {
	ChannelKeyword string
	JoinAt         time.Time
}

type Connection struct {
	*ircevent.Connection
	joinChannels map[string]Channel
	connected    bool
}

func Init(config *args.Result) error {
	conn = &Connection{
		ircevent.IRC(
			config.IrcNickname(),
			config.IrcUser(),
		),
		make(map[string]Channel),
		false,
	}
	conn.UseTLS = config.EnableSsl()
	if config.EnableSsl() && config.InsecureSkipVerify() {
		conn.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	if config.IrcKeyword() != "" {
		conn.Password = config.IrcKeyword()
	}
	err := conn.Connect(fmt.Sprintf(
		"%s:%d",
		config.IrcServer(),
		config.IrcPort(),
	))
	if err != nil {
		return err
	}
	conn.connected = true
	return nil
}

func GetConn() *Connection {
	if conn != nil && conn.connected {
		return conn
	}
	return nil
}

func (conn *Connection) IsJoined(channel string) bool {
	_, ok := conn.joinChannels[channel]
	return ok
}

func (conn *Connection) Join(channel string) {
	conn.Connection.Join(channel)
	conn.joinChannels[channel] = Channel{"", time.Now()}
}

func (conn *Connection) Part(channel string) {
	conn.Connection.Part(channel)
	delete(conn.joinChannels, channel)
}

func (conn *Connection) ChannelList() []string {
	list := []string{}
	for k, _ := range conn.joinChannels {
		list = append(list, k)
	}
	return list
}
