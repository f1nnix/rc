package main

import "strings"

func splitRemotePath(args []string) (string, string) {
	// Splits requated remote-path string and parses into chunks
	var requestedRemote, userPath string

	if strings.Contains(args[0], ":") {
		// remote with path provided
		chunks := strings.Split(args[0], ":")
		requestedRemote = chunks[0]
		userPath = chunks[1]
	} else {
		// only remote name provded
		requestedRemote = args[0]
		userPath = "/"
	}
	return requestedRemote, userPath
}

func getSyncDirection(args []string) string {
	// if no direction specified, assume "up"
	var direction string
	if len(args) > 1 && args[1] == "down" {
		direction = "down"
	} else {
		direction = "up"
	}
	return direction
}
