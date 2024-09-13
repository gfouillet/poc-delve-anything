package delve

import "fmt"

type MainWithArgs func(args []string) int

func Delve(main MainWithArgs) MainWithArgs {
    fmt.Println("Running in debug mode")
    return main
}
