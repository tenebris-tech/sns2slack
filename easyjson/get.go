// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package easyjson

import (
	"errors"
)

func (j *EasyJSON) Get(keys ...string) (interface{}, error) {
	var p interface{}

	// Get number of keys in search
	numKeys := len(keys)

	// Starting point
	p = j.Data

	// Iterate through keys
	for i, k := range keys {
		if val, ok := p.(map[string]interface{})[k]; ok {
			if (i + 1) >= numKeys {
				// This is the element we want
				return val, nil
			}

			// Update our pointer
			p = val

		} else {
			return "", errors.New("key not found")
		}
	}
	return "", errors.New("key not found")
}
