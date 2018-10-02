package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	rclone_args := []string{"--verbose", "--exclude=.DS_Store", "sync"}

	// test againts --dry-run option
	// TODO: map to separate config instance
	for i := 0; i < len(args); i++ {
		if args[i] == "--dry-run" {

			os.Exit(0)
		}
	}

	config := loadConfig()
	requestedRemote, userPath := splitRemotePath(args)
	direction := getSyncDirection(args)

	// get requested remote config. Exit, if not found
	remote := config.Remotes[requestedRemote]
	if remote.Remote == "" {
		log.Fatal("Remote ", requestedRemote, " was not found in rc config. Exiting")
	}

	localTarget := getLocalRcTarget(config, remote, userPath)
	remoteTarget := getRemoteRcTarget(config, remote, userPath)

	if direction == "up" {
		rclone_args = append(rclone_args, localTarget, remoteTarget)
	} else {
		rclone_args = append(rclone_args, remoteTarget, localTarget)
	}

	log.Println(rclone_args)

	// run task
	cmd := exec.Command("rclone", rclone_args...) //.Output()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
