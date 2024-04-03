package main

import (
	"fmt"
	"os"
)

var cmds = map[string]string{
	"create": "create",
}

func main() {
	args := os.Args
	switch cmd := args[1]; cmd {
	case cmds["create"]:
		fmt.Printf("create case\n")
	default:
		fmt.Printf("invalid command: %s\n", cmd)
	}
}
