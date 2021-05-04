// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package slack

import "fmt"

// Slack is used enable class-like behaviour
type Slack struct {
	Enabled bool
	Queue   chan Message
}

type Message struct {
	Source  string
	Title   string
	Details string
}

// New EasyJSON object
func New() Slack {
	var s Slack
	s.Enabled = true
	s.Queue = make(chan Message)

	// Start goroutine to send queue
	go s.SendQueue()

	// return
	return s
}

// Add message to queue
func (s *Slack) Add(msg Message) {
	s.Queue <- msg
}

// SendQueue is started by New() - loop and send queue
func (s *Slack) SendQueue() {
	var msg Message
	for {
		msg = <-s.Queue
		fmt.Printf("\nTitle: %s\nSource: %s\n\n", msg.Title, msg.Source)
		//fmt.Printf("\nTitle: %s\nSource: %s\n\n%s\n\n", msg.Title, msg.Source, msg.Details)
	}
}
