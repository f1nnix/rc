package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRcTargets(t *testing.T) {

	config := Config{
		BaseDir:     "/Users/user",
		SyncFolders: true,
	}
	config.Remotes = make(map[string]Remote)
	config.Remotes["test1"] = Remote{
		Remote: "db",
	}
	config.Remotes["test2"] = Remote{
		Remote:     "db",
		RemotePath: "/test2",
	}

	config.Remotes["test3"] = Remote{
		Remote:    "db",
		LocalPath: "/test3",
	}

	config.Remotes["test4"] = Remote{
		Remote:     "db",
		RemotePath: "/1p",
		LocalPath:  "/test4",
	}

	userPaths := [3]string{
		"/userPath",
		"userPath/",
		"/userPath/",
	}

	//getLocalRcTarget
	assert.Equal(t, "/Users/user/db",
		getLocalRcTarget(config, config.Remotes["test1"], ""))
	for i := 0; i < len(userPaths); i++ {
		assert.Equal(t, "/Users/user/db/userPath",
			getLocalRcTarget(config, config.Remotes["test1"], userPaths[0]))
	}

	assert.Equal(t, "/Users/user/db/test2",
		getLocalRcTarget(config, config.Remotes["test2"], ""))
	for i := 0; i < len(userPaths); i++ {
		assert.Equal(t, "/Users/user/db/test2/userPath",
			getLocalRcTarget(config, config.Remotes["test2"], userPaths[0]))
	}

	assert.Equal(t, "/Users/user/test3",
		getLocalRcTarget(config, config.Remotes["test3"], ""))
	for i := 0; i < len(userPaths); i++ {
		assert.Equal(t, "/Users/user/test3/userPath",
			getLocalRcTarget(config, config.Remotes["test3"], userPaths[0]))
	}

	assert.Equal(t, "/Users/user/test4",
		getLocalRcTarget(config, config.Remotes["test4"], ""))
	for i := 0; i < len(userPaths); i++ {
		assert.Equal(t, "/Users/user/test4/userPath",
			getLocalRcTarget(config, config.Remotes["test4"], userPaths[0]))
	}

	//getRemoteRcTarget
	assert.Equal(t, "db:/",
		getRemoteRcTarget(config, config.Remotes["test1"], ""))
	for i := 0; i < len(userPaths); i++ {
		assert.Equal(t, "db:/userPath",
			getRemoteRcTarget(config, config.Remotes["test1"], userPaths[i]))
	}

	assert.Equal(t, "db:/test2",
		getRemoteRcTarget(config, config.Remotes["test2"], ""))
	for i := 0; i < len(userPaths); i++ {
		assert.Equal(t, "db:/test2/userPath",
			getRemoteRcTarget(config, config.Remotes["test2"], userPaths[i]))
	}

}
