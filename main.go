/*
 *  Copyright (c) 2017 Open Baton (http://openbaton.org)
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/openbaton/go-openbaton/vnfm"
	_ "github.com/openbaton/go-openbaton/vnfm/amqp" // import needed to load the driver
	"github.com/openbaton/go-openbaton/vnfm/config"
	log "github.com/sirupsen/logrus"
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

	l := svc.Logger()

	h.Logger = l

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	join := make(chan struct{})

	go func() {
		<-sigChan

		l.WithFields(log.Fields{
			"tag": "dummy-main-sigint_callback",
		}).Warn("interrupt signal received, quitting")

		if err := svc.Stop(); err != nil {
			l.WithFields(log.Fields{
				"tag": "dummy-main-sigint_callback",
				"err": err,
			}).Fatal("stopping service failed")
		}

		close(join)
	}()

	if err = svc.Serve(); err != nil {
		l.WithFields(log.Fields{
			"tag": "dummy-main",
			"err": err,
		}).Fatal("VNFM failed during execution")
	}

	<-join

	l.WithFields(log.Fields{
		"tag": "dummy-main",
	}).Info("exiting cleanly")
}
