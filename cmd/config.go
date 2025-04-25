package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type GitwsConfig struct {
    RemoteUrl  string `json:"remoteUrl"`
    RootDomain string `json:"rootDomain"`
    RootOwner  string `json:"rootOwner"`
    RootRepo   string `json:"rootRepo"`
	RootBranch string `json:"rootBranch"`
	RootDir    string `json:"rootDir"`
	GitDir     string `json:"gitDir"`	
}

func updateGitwsRoot() {
	var path string
	if len(gitwsDir) == 0 {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		for len(path) > 1 {
			path, _ = filepath.Split(path)
			path = filepath.Clean(path)
			fmt.Printf("Checking %s\n", path)
			configPath := filepath.Join(path, ".gitws")
			if _, err := os.Stat(configPath); ! errors.Is(err, os.ErrNotExist) {
				fmt.Printf("Found %s\n", configPath)
				break
			}
		}
	}
	gitwsDir = path
}

func readGitwsConfig(gitws_root_dir string) GitwsConfig {
	gitws_config_file := filepath.Join(gitws_root_dir, ".gitws")
	jsonFile, err := os.Open(gitws_config_file)
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var config GitwsConfig
	json.Unmarshal(byteValue, &config)

	return config
}

func writeGitwsConfig(gitws_root_dir string, config GitwsConfig) {
	byteValue, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	gitws_config_file := filepath.Join(gitws_root_dir, ".gitws")
	err = os.WriteFile(gitws_config_file, byteValue, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func NewConfigCommand() *cobra.Command {
	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Show config of the current gitws directory",
		Run: func(cmd *cobra.Command, args []string) {
			config := readGitwsConfig(gitwsDir)
			fmt.Printf("%+v\n", config)
		},
	}
	return configCmd
}