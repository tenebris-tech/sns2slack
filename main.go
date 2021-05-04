// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sns2slack/slack"
	"strings"
	"syscall"

	"github.com/tenebris-tech/glog"
	"sns2slack/easyconfig"
)

const ProductName = "sns2slack"
const ProductVersion = "0.0.1"

var SlackQueue slack.Slack

func main() {
	var err error

	// Say hello
	glog.Infof("%s %s starting", ProductName, ProductVersion)

	// Load configuration
	// Try current working directory first, then /etc
	config, err := easyconfig.Load("sns2slack.conf")
	if err != nil {
		if strings.Contains(err.Error(), "no such file") ||
			strings.Contains(err.Error(), "cannot find the file") {

			// Attempt to load from /etc
			err = config.Load(string(os.PathSeparator) + "etc" + string(os.PathSeparator) + "sns2slack.conf")
			if err != nil {
				glog.Errorf("Unable to load config: %s", err.Error())
				fmt.Println(err.Error())
				os.Exit(1)
			}
		} else {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	config.Dump()

	// Setup signal catching
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// method invoked upon seeing signal
	go func() {
		for {
			s := <-signals
			glog.Noticef("Received signal: %v", s)
			AppCleanup()
		}
	}()

	// Get server configuration
	listenAddr := config.GetStrDef("server.listen", "0.0.0.0:8080")
	useTLS := config.GetBoolDef("server.tls", false)
	certFile := config.GetStrDef("server.certFile", "")
	keyFile := config.GetStrDef("server.keyFile", "")

	// Sanity checks
	if useTLS {
		if certFile == "" || keyFile == "" {
			// Log error, print, and bail
			tmp := "Fatal configuration error. For TLS, both certFile and keyFile must be specified."
			glog.Error(tmp)
			fmt.Println(tmp)
			AppCleanup()
		}
	}

	// Instantiate router
	router := newRouter(config)

	// Set up queue, which is global to account for HTTP server concurrency
	SlackQueue = slack.New()

	// Create server
	err = nil
	if useTLS {
		glog.Infof("Starting HTTPS server on %s", listenAddr)
		err = http.ListenAndServeTLS(listenAddr, certFile, keyFile, router)
	} else {
		glog.Infof("Starting HTTP server on %s", listenAddr)
		err = http.ListenAndServe(listenAddr, router)
	}

	if err != nil {
		glog.Errorf("HTTP server error: %s", err.Error())
	}
	AppCleanup()
}

// AppCleanup handles a graceful exit
func AppCleanup() {
	glog.Infof("%s %s stopping", ProductName, ProductVersion)
	os.Exit(0)
}
