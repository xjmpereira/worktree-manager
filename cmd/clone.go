package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cloneCmd)
}

func getParams(regEx, url string) (paramsMap map[string]string) {

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

var cloneCmd = &cobra.Command{
	Use:   "clone <repository>",
	Short: "Clone a repository and prepare it with GitWS structure.",
	Run: func(cmd *cobra.Command, args []string) {
		repository := args[0]
		matches := getParams(`(?P<method>https?|git|ssh)(?::\/\/(?:\w+@)?|@)(?P<domain>.*?)(?:\.org|\.com)(?:\/|:)(?P<fullname>(?P<owner>.+?)\/(?P<repo>.+?))(?:\.git|\/)?$`, repository)
		if len(matches) == 0 {
			err := fmt.Errorf("invalid repository: %s", repository)
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Cloning repository:\n")
		fmt.Printf("|    url: %v\n", repository)
		fmt.Printf("| domain: %v\n", matches["domain"])
		fmt.Printf("|  owner: %v\n", matches["owner"])
		fmt.Printf("|   repo: %v\n", matches["repo"])

		fmt.Printf("<< TODO >>\n")
	},
}
