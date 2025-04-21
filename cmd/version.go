package cmd

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// GitWSVersion is the version of the cli to be overwritten by goreleaser in the CI run with the version of the release in github
var GitWSVersion string

func getGitWSVersion() string {
	noVersionAvailable := "No version info available for this build, run 'git-ws help version' for additional info"
	
	if len(GitWSVersion) != 0 {
		return GitWSVersion
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return noVersionAvailable
	}

	// If no main version is available, Go defaults it to (devel)
	if bi.Main.Version != "(devel)" {
		return bi.Main.Version
	}

	var vcsRevision string
	var vcsTime time.Time
	for _, setting := range bi.Settings {
		switch setting.Key {
		case "vcs.revision":
			vcsRevision = setting.Value
		case "vcs.time":
			vcsTime, _ = time.Parse(time.RFC3339, setting.Value)
		}
	}

	if vcsRevision != "" {
		return fmt.Sprintf("%s, (%s)", vcsRevision, vcsTime)
	}

	return noVersionAvailable
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display application version information.",
	Long: `
The version command provides information about the application's version.

Git WS requires version information to be embedded at compile time.
For detailed version information, Git WS needs to be built as specified in the README installation instructions.
If Git WS is built within a version control repository and other version info isn't available,
the revision hash will be used instead.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		version := getGitWSVersion()
		fmt.Printf("GitWS CLI version: %v\n", version)
	},
}
