package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }
	currentUser, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get current user: %v\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "Running as user: %s (UID: %s)\n", currentUser.Username, currentUser.Uid)
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "--Usage: %s ENV_VAR_NAME\n", os.Args[0])
		os.Exit(1)
	}

	envName := os.Args[1]
	val := os.Getenv(envName)
	if val == "" {
		fmt.Fprintf(os.Stderr, "--Environment variable %q not set\n", envName)
		os.Exit(1)
	}

	fmt.Println(val)
}
