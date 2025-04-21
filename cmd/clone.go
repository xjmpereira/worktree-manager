package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"encoding/json"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cloneCmd)
}

// User struct which contains a name
// a type and a list of social links
type GitwsConfig struct {
    RemoteUrl  string `json:"gitws_remote_url"`
    RootDomain string `json:"gitws_root_domain"`
    RootOwner  string `json:"gitws_root_owner"`
    RootRepo   string `json:"gitws_root_repo"`
	RootBranch string `json:"gitws_root_branch"`
	RootDir    string `json:"gitws_root_dir"`
	GitDir     string `json:"gitws_git_dir"`
	
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

	fmt.Println(config.RootDomain)
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

func matchedGroups(regEx, url string) (paramsMap map[string]string) {
    var compRegEx = regexp.MustCompile(regEx)
    match := compRegEx.FindStringSubmatch(url)

    paramsMap = make(map[string]string)
    for i, name := range compRegEx.SubexpNames() {
        if i > 0 && i <= len(match) {
            paramsMap[name] = match[i]
        }
    }
    return paramsMap
}

func cloneRepository(repository string, directory string, branch string) {
	if _, err := os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(directory, os.ModePerm); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Creating directory: %s\n", directory)
			ps := exec.Command("git", "clone", repository, directory, "--branch", branch)
			out, err := ps.CombinedOutput()
			if err != nil {
				log.Fatal(string(out))
			}
		} else {
			fmt.Printf("Already exists: %s\n", directory)
		}
	}
}
var cloneCmd = &cobra.Command{
	Use:   "clone <repository>",
	Short: "Clone a repository and prepare it with GitWS structure.",
	Run: func(cmd *cobra.Command, args []string) {
		repository := args[0]
		user_home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal( err )
		}

		repo_groups := matchedGroups(`(?P<method>https?|git|ssh)(?::\/\/(?:\w+@)?|@)(?P<domain>.*?)(?:\.org|\.com)(?:\/|:)(?P<fullname>(?P<owner>.+?)\/(?P<repo>.+?))(?:\.git|\/)?$`, repository)
		if len(repo_groups) == 0 {
			err := fmt.Errorf("invalid repository: %s", repository)
			fmt.Println(err)
			os.Exit(1)
		}
		gitws_remote_url := strings.Trim(repository, " \t")
		gitws_root_domain := strings.Trim(repo_groups["domain"], " \t")
		gitws_root_owner := strings.Trim(repo_groups["owner"], " \t")
		gitws_root_repo := strings.Trim(repo_groups["repo"], " \t")

		ps := exec.Command("git", "ls-remote", "--symref", repository, "HEAD")
		out, err := ps.CombinedOutput()
		if err != nil {
			log.Fatal(string(out))
		}
		branch_groups := matchedGroups(`refs\/heads\/(?P<branch>.*) *HEAD.*`, string(out))
		gitws_root_branch := strings.Trim(branch_groups["branch"], " \t")
		gitws_root_dir := filepath.Join(user_home, gitws_root_owner, gitws_root_repo)
		gitws_git_dir := filepath.Join(gitws_root_dir, gitws_root_branch)

		config := GitwsConfig{
			RemoteUrl: gitws_remote_url,
			RootDomain: gitws_root_domain,
			RootOwner: gitws_root_owner,
			RootRepo: gitws_root_repo,
			RootBranch: gitws_root_branch,
			RootDir: gitws_root_dir,
			GitDir: gitws_git_dir,
		}
		writeGitwsConfig(config.RootDir, config)
		fmt.Printf("%+v\n", config)

		cloneRepository(config.RemoteUrl, config.GitDir, config.RootBranch)
	},
}
