package main

import (
	"fmt"
	"path/filepath"
)

func getLocalRcTarget(config Config, remote Remote, userPath string) string {
	var localTarget string

	//localPath specified
	if (&remote.LocalPath != nil) && (remote.LocalPath != "") {
		localTarget = fmt.Sprintf("%s/%s/%s/", config.BaseDir, remote.LocalPath, userPath)
	} else {
		localTarget = fmt.Sprintf("%s/%s/%s/%s/", config.BaseDir, remote.Remote, remote.RemotePath, userPath)
	}

	localTarget = filepath.Clean(localTarget)

	return localTarget
}

func getRemoteRcTarget(config Config, remote Remote, userPath string) string {
	var remoteTarget, localPath string

	localPath = fmt.Sprintf("%s/%s", remote.RemotePath, userPath)
	localPath = filepath.Clean(localPath)

	remoteTarget = fmt.Sprintf("%s:%s", remote.Remote, localPath)
	return remoteTarget
}