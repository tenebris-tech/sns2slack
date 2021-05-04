//
// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.
//

package easyconfig

import (
	"errors"
	"strings"
)

// nvsplit splits string into name-value pair using the specified separator
func nvsplit(s string, separator string) (string, string, error) {
	if len(s) < 3 {
		return "", "", errors.New("error splitting string less than 3 characters")
	}

	tmp := strings.Split(s, separator)
	if len(tmp) < 2 {
		return "", "", errors.New("error splitting string")
	}

	// Key must not be null, but we're ok with empty string as a value
	if tmp[0] == "" {
		return "", "", errors.New("error splitting string")
	}

	return strings.TrimSpace(tmp[0]), strings.TrimSpace(tmp[1]), nil
}

// extractFromBrackets returns the string between [ and ] or nothing
func extractFromBrackets(s string) string {

	// Trim leading and trailing spaces
	s = strings.TrimSpace(s)

	// Must be at least 3 characters long
	if len(s) < 3 {
		return ""
	}

	// Must start with [
	if !strings.HasPrefix(s, "[") {
		return ""
	}

	// Trim off the [
	tmp := strings.Trim(s, "[")

	// Last character must be ]
	if tmp[len(tmp)-1:] == "]" {
		tmp = tmp[:len(tmp)-1]
	} else {
		// Delimiter is missing, return ""
		tmp = ""
	}

	// Trim leading and trailing spaces
	return strings.TrimSpace(tmp)
}
