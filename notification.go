// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package main

import (
	"fmt"
	"sns2slack/easyjson"
	"sns2slack/slack"
	"strings"

	"github.com/tenebris-tech/glog"
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
		// Unable to get source
		source = "unknown"
	}

	var msg slack.Message
	msg.Source = source

	switch strings.ToLower(source) {
	case "aws.guardduty":
		msg.Title = j.GetStrDef("", "detail", "description")
		if msg.Title == "" {
			msg.Title = j.GetStrDef("Unknown - Please see full message details", "detail", "eventName")
		}
		msg.Title = "AWS GuardDuty: " + msg.Title
	default:
		msg.Title = j.GetStrDef("", "detail", "description")
		if msg.Title == "" {
			j.GetStrDef("Unknown - Please see full message details", "detail", "eventName")
		}
		msg.Title = source + ": " + msg.Title
	}

	msg.Details, err = j.Pretty()
	if err != nil {
		msg.Details = fmt.Sprintf("JSON format error: %s\n", err.Error())
	}

	// Write to SlackQueue
	SlackQueue.Add(msg)
}
