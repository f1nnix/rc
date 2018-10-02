package main

type Remote struct {
	Remote     string
	RemotePath string `yaml:"remote_path"`
	LocalPath  string `yaml:"local_path"`
}

type Config struct {
	BaseDir     string `yaml:"base_dir"`
	SyncFolders bool   `yaml:"sync_folders"`
	Remotes     map[string]Remote
}
