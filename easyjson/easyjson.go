// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package easyjson

import (
	"encoding/json"
	"errors"
)

// EasyJSON is used enable class-like behaviour
type EasyJSON struct {
	Data   interface{}
	Loaded bool
}

// New EasyJSON object
func New() EasyJSON {
	var j EasyJSON
	j.Loaded = false
	return j
}

// Unmarshal string
func (j *EasyJSON) Unmarshal(str string) error {
	err := json.Unmarshal([]byte(str), &j.Data)
	j.Loaded = true
	return err
}

// GetStr retrieves an item and returns a string or error
func (j *EasyJSON) GetStr(str ...string) (string, error) {
	r, err := j.Get(str...)
	if err != nil {
		return "", err
	}
	return vToString(r)
}

// GetStrDef retrieves an item and returns a string or the supplied default
func (j *EasyJSON) GetStrDef(def string, str ...string) string {
	r, err := j.Get(str...)
	if err != nil {
		return def
	}

	rString, err := vToString(r)
	if err != nil {
		return def
	}

	return rString
}

// GetInt retrieves an item and returns an int or error
func (j *EasyJSON) GetInt(str ...string) (int, error) {
	r, err := j.Get(str...)
	if err != nil {
		return 0, err
	}
	return vToInt(r)
}

// GetFloat64 retrieves an item and returns a float64 or error
func (j *EasyJSON) GetFloat64(str ...string) (float64, error) {
	r, err := j.Get(str...)
	if err != nil {
		return 0.0, err
	}
	return vToFloat64(r)
}

// GetBool retrieves an item and returns a boolean or error
func (j *EasyJSON) GetBool(str ...string) (bool, error) {
	r, err := j.Get(str...)
	if err != nil {
		return false, err
	}
	return vToBool(r)
}

func (j *EasyJSON) Pretty() (string, error) {
	b, err := json.MarshalIndent(j.Data, "", "  ")
	return string(b), err
}

func (j *EasyJSON) Get(keys ...string) (interface{}, error) {
	var p interface{}

	// Get number of keys to search
	numKeys := len(keys)

	// Starting point
	p = j.Data

	// Iterate over keys
	for i, k := range keys {
		if val, ok := p.(map[string]interface{})[k]; ok {
			// Key k exists in map
			if (i + 1) >= numKeys {
				// This is the element requested (last key)
				return val, nil
			}

			// Continue to look for next key
			// Update the pointer
			p = val

		} else {
			return "", errors.New("key not found")
		}
	}
	return "", errors.New("key not found")
}
