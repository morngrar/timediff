// Copyright 2022 timediff Authors
// SPDX-License-Identifier: Apache-2.0

package timediff

import (
	Z "github.com/rwxrob/bonzai/z"
)

func ExampleSub() {
	Z.ExitOn()
	Sub.Call(nil, "08:00", "15:30")
	Sub.Call(nil, "20", "10")
	Z.ExitOff()

	// Output:
	// 7.50 hours
	// 14.00 hours
}
