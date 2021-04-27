// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package easyjson

import (
	"errors"
	"fmt"
)

func vToString(v interface{}) (string, error) {
	return fmt.Sprintf("%v", v), nil
}

func vToInt(v interface{}) (int, error) {
	i, ok := v.(int)
	if ok {
		return i, nil
	}
	return 0, errors.New("wrong type")
}

func vToFloat64(v interface{}) (float64, error) {
	i, ok := v.(float64)
	if ok {
		return i, nil
	}
	return 0, errors.New("wrong type")
}

func vToBool(v interface{}) (bool, error) {
	i, ok := v.(bool)
	if ok {
		return i, nil
	}
	return false, errors.New("wrong type")
}
