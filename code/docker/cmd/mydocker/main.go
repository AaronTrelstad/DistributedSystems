package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mydocker run <command>")
		os.Exit(1)
	}

	cmd := os.Args[1];
	args := os.Args[2];

	if err := runContainer(cmd, args); err != nil {
		fmt.Printf("Error running container: %v\n", err);
		os.Exit(1);
	}
}
