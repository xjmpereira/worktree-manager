package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	cli "github.com/urfave/cli/v3"
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

func cwdOr(path string) string {
	var out string = path
	if len(out) == 0 {
		currentPath, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		out = currentPath
	}
	return out
}

func searchGitwsRoot(dir string) (string, error) {
	path := filepath.Join(cwdOr(dir), ".gitws")
	for len(path) > 1 {
		path, _ = filepath.Split(path)
		path = filepath.Clean(path)
		configPath := filepath.Join(path, ".gitws")
		if _, err := os.Stat(configPath); ! errors.Is(err, os.ErrNotExist) {
			return path, nil
		}
	}
	return "", fmt.Errorf("error: GitWS config not found")
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

func ConfigCmd() *cli.Command{
	cmd := &cli.Command{
		Name:  "config",
		Usage: "Show config of the current gitws directory",
		Action: ConfigFn,
	}
	return cmd
}

func ConfigFn(ctx context.Context, cmd *cli.Command) error {
	fmt.Printf("%+v\n", gitwsConfig)
	return nil
}
