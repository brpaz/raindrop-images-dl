package cmd

import (
	"runtime"

	"github.com/spf13/cobra"
)

type VersionInfo struct {
	Version   string
	GitCommit string
	BuildDate string
}

// NewVersionCmd initializes a command that prints the version information.
func NewVersionCmd(versionInfo VersionInfo) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display version information",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("Version: %s\n", versionInfo.Version)
			cmd.Printf("Git Commit: %s\n", versionInfo.GitCommit)
			cmd.Printf("Build Date: %s\n", versionInfo.BuildDate)
			cmd.Printf("Go Version: %s\n", runtime.Version())
			cmd.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		},
	}

	return cmd
}
