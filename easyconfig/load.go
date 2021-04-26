//
// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.
//

package easyconfig

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Load configuration from specified file
func load(c *EasyConfig, fileName string) error {

	// Have we previously loaded a config file?
	if c.Loaded {
		// Delete everything in the existing config
		for k := range c.Data {
			delete(c.Data, k)
		}
		c.Loaded = false
	}

	// Open file for read
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	//noinspection GoUnhandledErrorResult
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// Read line and increment line counter
		line := scanner.Text()
		lineCount++

		// Ignore empty lines and comments
		if len(line) < 1 ||
			strings.HasPrefix(line, "#") ||
			strings.HasPrefix(line, ";") ||
			strings.HasPrefix(line, "/") {
			continue
		}

		// Split line into name value pair
		name, value, err := nvsplit(line, "=")
		if err != nil {
			tmp := fmt.Sprintf("error parsing line %d: %s", lineCount, line)
			return errors.New(tmp)
		}

		// Store in map with key (name) forced to lower case
		c.Data[strings.ToLower(name)] = value
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Turn on configured flag
	c.Loaded = true
	return nil
}
