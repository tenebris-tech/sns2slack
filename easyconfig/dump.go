//
// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.
//

package easyconfig

import "fmt"

// dump the configuration to stdout
func dump(c *EasyConfig) {
	fmt.Printf("Current configuration:\n")
	for n, v := range c.Data {
		fmt.Printf("\t%s = %s\n", n, v)
	}
}
