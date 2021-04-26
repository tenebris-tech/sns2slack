//
// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.
//

// Package easyconfig reads a simple config file and allows access to each parameter.
package easyconfig

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// EasyConfig is used enable class-like behaviour
type EasyConfig struct {
	Data   map[string]string
	Loaded bool
}

// Private configured flag

// New EasyConfig instance
func New() EasyConfig {
	// Instantiate, initialize, and return
	var c EasyConfig
	c.Loaded = false
	c.Data = make(map[string]string)
	return c
}

// Load new EasyConfig instance from file
func Load(fileName string) (EasyConfig, error) {
	c := New()
	err := c.Load(fileName)
	return c, err
}

// Load configuration from file
func (c *EasyConfig) Load(fileName string) error {
	return load(c, fileName)
}

// Dump configuration to stdout for diagnostics
func (c *EasyConfig) Dump() {
	dump(c)
}

// Exists checks if the specified configuration key exits
func (c *EasyConfig) Exists(key string) bool {
	_, ok := c.Data[key]
	return ok
}

// GetStr retrieves a string value or returns an error
func (c *EasyConfig) GetStr(key string) (string, error) {
	value, ok := c.Data[strings.ToLower(key)]
	if ok == false {
		tmp := fmt.Sprintf("key %s does not exist", key)
		return "", errors.New(tmp)
	}
	return value, nil
}

// GetStrDef retrieves a string value or returns the specified default
func (c *EasyConfig) GetStrDef(key string, defaultValue string) string {
	value, err := c.GetStr(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetInt retrieves an integer value or returns an error
func (c *EasyConfig) GetInt(key string) (int, error) {
	value, ok := c.Data[strings.ToLower(key)]
	if ok == false {
		tmp := fmt.Sprintf("key %s does not exist", key)
		return 0, errors.New(tmp)
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("error converting value to integer")
	}
	return i, nil
}

// GetIntDef retrieves an integer value or returns the specified default
func (c *EasyConfig) GetIntDef(key string, defaultValue int) int {
	value, err := c.GetInt(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetInt64 retrieves an int64 value or returns an error
func (c *EasyConfig) GetInt64(key string) (int64, error) {
	value, ok := c.Data[strings.ToLower(key)]
	if ok == false {
		return 0, errors.New("key does not exist")
	}

	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, errors.New("error converting value to integer")
	}

	return i, nil
}

// GetInt64Def retrieves an int64 value or returns the specified default
func (c *EasyConfig) GetInt64Def(key string, defaultValue int64) int64 {
	value, err := c.GetInt64(key)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetBool retrieves a boolean value or returns an error
func (c *EasyConfig) GetBool(key string) (bool, error) {
	value, ok := c.Data[strings.ToLower(key)]
	if ok == false {
		return false, errors.New("key does not exist")
	}

	if strings.EqualFold(value, "true") {
		return true, nil
	}

	if strings.EqualFold(value, "yes") {
		return true, nil
	}

	if strings.EqualFold(value, "1") {
		return true, nil
	}

	if strings.EqualFold(value, "false") {
		return false, nil
	}

	if strings.EqualFold(value, "no") {
		return false, nil
	}

	if strings.EqualFold(value, "0") {
		return false, nil
	}

	return false, errors.New("error converting value to boolean")
}

// GetBoolDef retrieves a boolean value or the specified default
func (c *EasyConfig) GetBoolDef(key string, defaultValue bool) bool {
	value, err := c.GetBool(key)
	if err != nil {
		return defaultValue
	}
	return value
}
