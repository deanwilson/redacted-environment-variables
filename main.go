package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
	"path/filepath"
	"strings"
)

func DefaultConfig() *toml.Tree {

	config, _ := toml.Load(`
  [config]
  redacted = "123123123"
  `)

	return config

}

func LocateConfigFile(filename string) string {
	var paths []string     // places to search
	var config_path string // the located config file

	current, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory: ", err)
		os.Exit(1)
	}

	// local directory first
	paths = append(paths, filepath.Join(current, filename))

	// then the home directory
	homedir, err := os.UserHomeDir()
	paths = append(paths, filepath.Join(homedir, filename))

	// then fail back to system wide /etc
	paths = append(paths, filepath.Join("/etc", filename))

	for _, path := range paths {
		_, err = os.Stat(path)
		if !os.IsNotExist(err) {
			config_path = path
			break
		}
	}

	return config_path
}

func LoadConfig(filename string) *toml.Tree {
	var config *toml.Tree

	config_file := LocateConfigFile(filename)

	// If a config file is found use, otherwise defaults
	if config_file != "" {
		config, _ = toml.LoadFile(config_file)
	} else {
		config = DefaultConfig()
	}

	return config
}

func ToRedact() []string {
	var redactedNames = []string{
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"GITHUB_TOKEN",
		"GITHUB_AUTH_TOKEN",
		"_TOKEN_",
	}

	return redactedNames
}

func main() {
	config := LoadConfig(".redacted-environment.toml")
	redacted_string := config.Get("config.redacted").(string)

	redactedVariables := ToRedact()

	for _, envvar := range os.Environ() {
		pair := strings.SplitN(envvar, "=", 2)

		name, value := pair[0], pair[1]

		for _, redacted := range redactedVariables {
			if strings.Contains(name, redacted) {
				value = redacted_string
			}
		}

		fmt.Println(name + "=" + value)
	}
}
