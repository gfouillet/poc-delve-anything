// Package delve allows to run any binaries with latest version of delve.
//
// To make it works:
//  1. compile your code with
//     -gcflags "all=-N -l"
//  2. Use [Delve] function to encapsulate your main function.
//
// # Tips
//
// Don't bring delve in production:
//
// Add next to your main.go a init_debug.go file which will be compiled only if the tags debug is
// passed to the compiler:
//
//	go build -gcflags "all=-N -l" -tags debug path/to/my/package
//
// This file will contains an init function which will do encapsulate the main function into Delve,
// only if the tags is setted. This way, delve dependencies won't be shipped in production.
//
// # Example
//
// debug_init.go
//
//	//go:build debug
//	// +build debug
//
//	package main
//
//	import "github.com/gfouillet/poc-delve-anything/delve"
//
//	func init() {
//	   runMain = delve.Delve(/*options*/)(runMain)
//	}
//
// main.go
//
//	package main
//
//	import (
//	   "os"
//	)
//
//	var runMain = mainArgs
//
//	func main() {
//	   os.Exit(runMain(os.Args))
//	}
//
//	func mainArgs(args []string) int { /* ... */ }
package delve
