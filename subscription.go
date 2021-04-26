// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package main

import (
	"github.com/tenebris-tech/glog"
	"net/http"
)

// confirmSubscription to SNS topic
func confirmSubscription(url string, arn string) {
	glog.Infof("Confirming SNS subscription for ARN %s", arn)

	resp, err := http.Get(url)
	if err != nil {
		glog.Errorf("Confirm SNS subscription for ARN %s failed: %s", arn, err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		glog.Errorf("Confirm SNS subscription for ARN %s failed: %s", arn, resp.Status)
	}

	glog.Infof("Confirm SNS subscription for ARN %s succeeded: %s", arn, resp.Status)
}
