package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(1)
	}
	//io.Copy(os.Stdout, f)
	f.WriteTo(os.Stdout)
}
