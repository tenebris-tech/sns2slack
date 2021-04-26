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
