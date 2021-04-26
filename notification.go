// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package main

import "fmt"

// Notification from SNS
func notification(n SNSNotification) {
	fmt.Println("Type:", n.Type)
	fmt.Println("Msg:", n.Message)

	// Detect subscription message
	if n.Type == "SubscriptionConfirmation" {
		confirmSubscription(n.SubscribeURL, n.TopicArn)
	}



}
