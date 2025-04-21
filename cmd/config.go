package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

type GitwsConfig struct {
    RemoteUrl  string `json:"remoteUrl"`
    RootDomain string `json:"rootDomain"`
    RootOwner  string `json:"rootOwner"`
    RootRepo   string `json:"rootRepo"`
	RootBranch string `json:"rootBranch"`
	RootDir    string `json:"rootDir"`
	GitDir     string `json:"gitDir"`	
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

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show config of the current gitws directory",
	Run: func(cmd *cobra.Command, args []string) {
		config := readGitwsConfig(gitwsDir)
		fmt.Printf("%+v\n", config)
	},
}
