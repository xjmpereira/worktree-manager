package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v3"
)

var (
	currentDir     string
	wsTempCmd      string
	foundConfig    bool = false
	gitwsConfig    GitwsConfig
)

func RootCmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "ws",
		Usage: "A Program to manage git worktrees",
		Before: RootBeforeFn,
		Action: RootFn,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "currentdir",
				Value: "",
				Aliases: []string{"C"},
				Usage: "Use this as the current directory",
			},
		},
		Commands: []*cli.Command{
			ConfigCmd(),
			CloneCmd(),
			ListCmd(),
			CreateCmd(),
			RmCmd(),
			AddCmd(),
		},
	}
	return cmd
}

func RootBeforeFn(ctx context.Context, cmd *cli.Command) (context.Context, error) {
	wsTempCmd = os.Getenv("GITWS_TMP_CMD")
	os.Remove(wsTempCmd)

	currentDir = cwdOr(cmd.String("currentdir"))
	rootDir, err := searchGitwsRoot(currentDir)
	if err == nil {
		foundConfig = true
		gitwsConfig = readGitwsConfig(rootDir)
	}
	return nil, nil
}

func RootFn(ctx context.Context, cmd *cli.Command) error {
	wsList := GetWsList()
	iterChoice := iterative(wsList)
	postCmd := fmt.Sprintf(`cd %s`, iterChoice.Path)
	SetPostCmd(wsTempCmd, postCmd)
	return nil
}

func SetPostCmd(tempCmd string, cmd string) {
	err := os.WriteFile(tempCmd, []byte(cmd), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
