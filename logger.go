// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/tenebris-tech/glog"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := statusWriter{ResponseWriter: w}
		inner.ServeHTTP(&sw, r)
		src := getIP(r)

		// Don't log health checks to reduce log noise
		if name != "health" {
			glog.Infof("%s %s %s %s %d %d %d",
				src, name, r.Method, r.RequestURI, sw.status, sw.length, time.Since(start))
		}
	})
}

// getIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func getIP(r *http.Request) string {
	var s = ""
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		s = forwarded
	} else {
		s = r.RemoteAddr
	}

	// Clean up, remove port number
	if len(s) > 0 {
		if strings.HasPrefix(s, "[") {
			// IPv6 address
			t := strings.Split(s, "]")
			s = t[0][1:]
		} else {
			// IPv4 - just hack off port number
			t := strings.Split(s, ":")
			s = t[0]
		}
	}
	return s
}
