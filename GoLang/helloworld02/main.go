package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("This is the helloworld02 program.\nArgument: %v\n", args[1:])
}
