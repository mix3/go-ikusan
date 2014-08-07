package main

import (
	"fmt"
	"os"

	"github.com/mix3/go-ikusan/args"
	"github.com/mix3/go-ikusan/irc"
	"github.com/mix3/go-ikusan/server"
)

func init() {
	args.Parse(os.Args[1:])
	config := args.GetConfig()
	if config.Help() {
		fmt.Print(args.Help())
		os.Exit(0)
	}
	if !config.IsValid() {
		fmt.Fprint(os.Stderr, args.Usage())
		os.Exit(0)
	}
}

func main() {
	config := args.GetConfig()

	// server
	n := server.New(config)
	go func() {
		n.Run(fmt.Sprintf("%s:%d", config.HttpHost(), config.HttpPort()))
	}()

	// irc
	irc.Run(config)
}
