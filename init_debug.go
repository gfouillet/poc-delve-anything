//go:build debug
// +build debug

package main

import (
	"github.com/gfouillet/poc-delve-anything/delve"
	"github.com/gfouillet/poc-delve-anything/delve/config"
)

func init() {
	runMain = delve.Delve(config.Default(),
		config.WithPort(1122),
		config.WaitDebugger())(runMain)
}
