//go:build debug
// +build debug

package main

import "github.com/gfouillet/poc-delve-anything/delve"

func init() {
    runMain = delve.Delve(runMain)
}
