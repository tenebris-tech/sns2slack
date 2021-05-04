// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package main

import (
	"fmt"
	"github.com/tenebris-tech/glog"
	"sns2slack/easyjson"
)

// Notification from SNS
func notification(n SNSNotification) {

	// Detect subscription message
	if n.Type == "SubscriptionConfirmation" {
		confirmSubscription(n.SubscribeURL, n.TopicArn)
		return
	}

	// Detect unsubscription message
	if n.Type == "UnsubscribeConfirmation" {
		glog.Noticef("Received SNS unsubscribe confirmation for ARN %s", n.TopicArn)
		return
	}

	if n.Type != "Notification" {
		glog.Warnf("Received unknown SNS notification type for ARN %s", n.TopicArn)
		return
	}

	// We have a notification - unmarshal message
	j := easyjson.New()
	err := j.Unmarshal(n.Message)
	if err != nil {
		glog.Warnf("Error attempting to unmarshal JSON message %s ARN %s", n.MessageId, n.TopicArn)
		return
	}

	source, err := j.GetStr("source")
	if err != nil {
		fmt.Printf("get source: %s\n", err.Error())
	}

	description, err := j.GetStr("detail", "description")
	if err != nil {
		fmt.Printf("get detail/description: %s\n", err.Error())
	}

	fmt.Println("Source:", source)
	fmt.Println("Description:", description)

	p, err := j.Pretty()
	if err != nil {
		fmt.Printf("\nPretty Print Error: %s\n", err.Error())
	} else {
		fmt.Println("\n", p)
	}
	fmt.Println("---")

	// Write to SlackQueue
	SlackQueue.Add(fmt.Sprintf("Description: %s", description))

}
