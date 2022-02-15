package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println(0)
		return
	}
	s := args[1]
	if s == "" {
		fmt.Println(0)
		return
	}
	fmt.Println(len(strings.Fields(s)))
}
