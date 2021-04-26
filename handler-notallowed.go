// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package main

import (
	"encoding/json"
	"github.com/tenebris-tech/glog"
	"net/http"
)

type NotAllowedResp struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
}

// NotAllowedHandler is for errors
//noinspection GoUnusedParameter
func NotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	var resp NotAllowedResp

	// Set header
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	resp.Status = "error"
	resp.Code = http.StatusMethodNotAllowed
	w.WriteHeader(resp.Code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		glog.Errorf("JSON encode error in health handler: %s",err.Error())
	}
}
