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

type SNSResp struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

type SNSNotification struct {
	Message          string
	MessageId        string
	Signature        string
	SignatureVersion string
	SigningCertURL   string
	Subject          string
	Timestamp        string
	TopicArn         string
	Type             string
	UnsubscribeURL   string
}

// SNSHandler handles SNS messages
//noinspection GoUnusedParameter
func SNSHandler(w http.ResponseWriter, r *http.Request) {
	var resp SNSResp

	// Set header
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Read and decode message body
	n := SNSNotification{}
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		glog.Errorf("Error reading POST body: %s", err.Error())
		resp.Status = "error"
		resp.Code = http.StatusBadRequest
		w.WriteHeader(resp.Code)
	} else {

		// Send notification for processing
		notification(n)

		// Success
		resp.Status = "ok"
		resp.Code = http.StatusOK
		w.WriteHeader(resp.Code)
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		glog.Errorf("JSON encode error in SNS handler: %s", err.Error())
	}
}
