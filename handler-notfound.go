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

type NotFoundResp struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
}

// NotFoundHandler is for errors
//noinspection GoUnusedParameter
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	var resp NotFoundResp

	// Set header
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	resp.Status = "error"
	resp.Code = http.StatusNotFound
	w.WriteHeader(resp.Code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		glog.Errorf("JSON encode error in health handler: %s",err.Error())
	}
}
