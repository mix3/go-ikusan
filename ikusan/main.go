package main

import (
	"fmt"
	"log"
	"os"

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
	go func() {
		Config := args.GetConfig()
		n.Run(fmt.Sprintf("%s:%d", Config.HttpHost(), Config.HttpPort()))
	}()
	irc.GetConn().Loop()
}
