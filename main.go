package main

import (
	"fmt"
	"os"
)

func main() {
	if err := RunCli(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
