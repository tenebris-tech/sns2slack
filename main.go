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
	"strings"
	"syscall"
	"time"

	"github.com/tenebris-tech/glog"
	"main/easyconfig"
)

const ProductName = "sns2slack"
const ProductVersion = "0.0.1"

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
	listenAddr := config.GetStrDef("listen", "0.0.0.0:8080")
	HTTPTimeout := config.GetIntDef("HTTPTimeout", 60)
	HTTPIdleTimeout := config.GetIntDef("HTTPIdleTimeout", 60)

	// Instantiate router
	router := newRouter()

	// Create server
	s := &http.Server{
		Addr:              listenAddr,
		Handler:           router,
		ReadHeaderTimeout: time.Duration(HTTPTimeout) * time.Second,
		ReadTimeout:       time.Duration(HTTPTimeout) * time.Second,
		WriteTimeout:      time.Duration(HTTPTimeout) * time.Second,
		IdleTimeout:       time.Duration(HTTPIdleTimeout) * time.Second,
	}

	glog.Infof("Starting HTTP server on %s", listenAddr)

	err = s.ListenAndServe()
	if err != nil {
		glog.Errorf("HTTP server error: %s", err.Error())
	}
	AppCleanup()
}

// AppCleanup handles a graceful exit
func AppCleanup() {
	// Log and exit
	glog.Infof("%s %s stopping", ProductName, ProductVersion)
	os.Exit(0)
}