package delve

import (
    "fmt"
    "os"

    "github.com/go-delve/delve/cmd/dlv/cmds"
)

type MainWithArgs func(args []string) int

const (
    FORCE_DELVE_PORT = "FORCE_DELVE_PORT"
    NoDebug          = "DELVE_ANYTHING_NO_DEBUG"
)

// Delve wraps a MainWithArgs function to enable debugging.
//
// It works in two phases:
// At first call, it launch the delve command `exec` with the current binary. However, before the
// call it set a environment variable DELVE_ANYTHING_NO_DEBUG.
// At the second call, DELVE_ANYTHING_NO_DEBUG is setted, so it just return the "normal" main.
func Delve(main MainWithArgs) MainWithArgs {

    if _, exists := os.LookupEnv(NoDebug); exists {
        return main
    }
    if err := os.Setenv(NoDebug, "1"); err != nil {
        fmt.Printf("Failed to set env %q: %v\n", NoDebug, err)
        fmt.Println("Starting without debug mode...")
        return main
    }
    port := "0"
    if value, exists := os.LookupEnv(FORCE_DELVE_PORT); exists {
        port = value
    }

    // Run with delve
    return func(args []string) int {
        command := args[0]
        dlvArgs := []string{
            "--headless",
            "--continue",
            "--listen=:" + port,
            "--api-version=2",
            "--accept-multiclient",
            "exec",
            command, "--",
        }
        dlvArgs = append(dlvArgs, args[1:]...)
        fmt.Printf("Starting dlv with %v\n", dlvArgs)

        dlvCmd := cmds.New(false)
        dlvCmd.SetArgs(dlvArgs)

        defer fmt.Println("dlv has stopped")

        fmt.Println("Running in debug mode")
        if err := dlvCmd.Execute(); err != nil {
            fmt.Printf("Failed to run dlv: %v\n", err)
            return 1
        }
        fmt.Printf("End debug mode\n")
        return 0
    }
}
