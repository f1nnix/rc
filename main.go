package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

func syncRemote(requestedRemote string, userPath string, direction string, dryRun bool) int {
	rcloneArgs := []string{"--verbose", "--exclude=.DS_Store", "sync"}

	config := loadConfig()
	// get requested remote config. Exit, if not found
	remote := config.Remotes[requestedRemote]
	if remote.Remote == "" {
		log.Fatal("Remote ", requestedRemote, " was not found in rc config. Exiting")
	}

	localTarget := getLocalRcTarget(config, remote, userPath)
	remoteTarget := getRemoteRcTarget(config, remote, userPath)

	if direction == "up" {
		rcloneArgs = append(rcloneArgs, localTarget, remoteTarget)
	} else {
		rcloneArgs = append(rcloneArgs, remoteTarget, localTarget)
	}

	log.Println(rcloneArgs)

	if dryRun == true {
		return 0
	}

	cmd := exec.Command("rclone", rcloneArgs...) //.Output()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return 0
}

func main() {
	var DryRun bool

	var rootCmd = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long:  `echo is for echoing anything back. Echo works a lot like print, except it has a child command.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {

			//sync all vaults up
			if len(args) == 0 {
				config := loadConfig()

				for requestedRemote, _ := range config.Remotes {
					syncRemote(requestedRemote, "/", "up", DryRun)
				}

				os.Exit(0)
			}

			//sync all vaults down
			if args[0] == "down" {
				config := loadConfig()

				for requestedRemote, _ := range config.Remotes {
					syncRemote(requestedRemote, "/", "down", DryRun)
				}

				os.Exit(0)
			}

			//sync particular remote
			requestedRemote, userPath := splitRemotePath(args)
			direction := getSyncDirection(args)

			syncRemote(requestedRemote, userPath, direction, DryRun)

		},
	}

	//flags
	rootCmd.PersistentFlags().BoolVarP(&DryRun, "dry-run", "d", false, "run generator without sync")

	rootCmd.Execute()
}
