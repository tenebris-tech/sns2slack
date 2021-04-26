// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package main

import (
	"encoding/json"
	"github.com/tenebris-tech/glog"
	"net/http"
	"os"
)

type HealthResp struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
}

// HealthHandler is for health check
//noinspection GoUnusedParameter
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	var resp HealthResp

	// Set header
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Check for presence of /tmp/down
	if _, err := os.Stat(string(os.PathSeparator) + "tmp" + string(os.PathSeparator) + "down"); err == nil {
		// exists -- send status down and 503
		resp.Status = "down"
		resp.Code = http.StatusServiceUnavailable
		w.WriteHeader(resp.Code)
	} else {
		// does not exist - send ok and 200
		resp.Status = "ok"
		resp.Code = http.StatusOK
		w.WriteHeader(http.StatusOK)
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		glog.Errorf("JSON encode error in health handler: %s",err.Error())
	}
}
