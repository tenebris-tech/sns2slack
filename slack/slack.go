// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package slack

import "fmt"

// Slack is used enable class-like behaviour
type Slack struct {
	Enabled bool
	Queue   chan string
}

// New EasyJSON object
func New() Slack {
	var s Slack
	s.Enabled = true
	s.Queue = make(chan string)

	// Start goroutine to send queue
	go s.SendQueue()

	// return
	return s
}

// Add message to queue
func (s *Slack) Add(str string) {
	fmt.Println("QUEUED: %s", str)
	s.Queue <- str
}

func (s *Slack) SendQueue() {
	var msg string
	for {
		msg = <-s.Queue
		fmt.Printf("***GOT***: %s", msg)
	}
}
