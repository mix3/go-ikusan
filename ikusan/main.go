package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mix3/go-ikusan/args"
	"github.com/mix3/go-ikusan/irc"
	"github.com/mix3/go-ikusan/server"
)

func init() {
	args.Parse(os.Args[1:])
	Config := args.GetConfig()
	if Config.Help() {
		fmt.Print(args.Help())
		os.Exit(0)
	}
	if !Config.IsValid() {
		fmt.Fprint(os.Stderr, args.Usage())
		os.Exit(0)
	}
	err := irc.Init(Config)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	n := server.New("")
	Config := args.GetConfig()
	go func() {
		n.Run(fmt.Sprintf("%s:%d", Config.HttpHost(), Config.HttpPort()))
	}()
	conn := irc.GetConn()
	quit := conn.GetQuitChan()
	var err error
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
				time.Sleep(time.Duration(Config.IrcReconnectInterval()) * time.Second)
			}
		}
	}
}
