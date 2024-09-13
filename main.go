package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

var runMain = mainArgs

func main() {
    os.Exit(runMain(os.Args))
}

func mainArgs(args []string) int {
    fmt.Println("Hello, World!")
    println("cmd: " + strings.Join(args, " "))

    var command string
    for command != "exit" {
        fmt.Print("(exit to quit)>\t")
        input := bufio.NewScanner(os.Stdin)
        input.Scan()
        command = input.Text()
        fmt.Println(input.Text())
    }

    return 0
}
