package main

import (
	"runtime/debug"
)

func init() { debug.SetGCPercent(-1) }
