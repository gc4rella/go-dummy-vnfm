package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/mcilloni/go-openbaton/vnfm"
	_ "github.com/mcilloni/go-openbaton/vnfm/amqp" // import needed to load the driver
	"github.com/mcilloni/go-openbaton/vnfm/config"
)

var confPath = flag.String("cfg", "config.toml", "a TOML file to be loaded as config")

func main() {
	flag.Parse()

	cfg, err := config.LoadFile(*confPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: while loading config file %s: %v\n", *confPath, err)
		os.Exit(1)
	}

	h := &handl{}

	svc, err := vnfm.New("amqp", h, cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: while initialising VNFM %s: %v\n", *confPath, err)
		os.Exit(1)
	}

	h.Logger = svc.Logger()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		l := svc.Logger()
		l.Warnln("interrupt signal received, quitting")
		if err := svc.Stop(); err != nil {
			l.Fatalf("while stopping service: %v", err)
		}

		os.Exit(0)
	}()

	if err = svc.Serve(); err != nil {
		fmt.Fprintf(os.Stderr, "error: while VNFM was starting: %v\n", err)
		os.Exit(1)
	}

	if err = svc.Stop(); err != nil {
		fmt.Fprintf(os.Stderr, "error: while VNFM was stopping: %v\n", err)
		os.Exit(1)
	}
}
