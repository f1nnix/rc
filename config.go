package main

import (
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

func loadConfig() Config {
	// load config with priority

	configPaths := [2]string{
		"config.yml",
		"/Users/user/.rc.yml",
	}

	var configData []byte
	for i := 0; i < len(configPaths); i++ {
		// if path exists, read config
		if _, err := os.Stat(configPaths[i]); err == nil {
			configData, err = ioutil.ReadFile(configPaths[i])
			break
		}
	}

	config := Config{}
	err := yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	return config
}
