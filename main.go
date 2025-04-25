package main

import (
	"context"
	"git-ws/cmd"
	"log"
	"os"
)

func main() {
	wsCmd := cmd.RootCmd()
    if err := wsCmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
