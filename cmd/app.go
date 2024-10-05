package cmd

import (
	"github.com/spf13/cobra"

	"github.com/brpaz/raindrop-images-dl/internal/version"
)

// App represents an instance of the CLI application.
type App struct {
	rootCmd *cobra.Command
}

// NewApp creates a new instance of the CLI application.
func NewApp() *App {
	app := &App{
		rootCmd: NewRootCmd(),
	}

	app.registerCommands()

	return app
}

func (a *App) registerCommands() {
	versionCmd := NewVersionCmd(VersionInfo{
		Version:   version.Version,
		GitCommit: version.GitCommit,
		BuildDate: version.BuildDate,
	})

	downloadCmd := NewDownloadCmd()

	a.rootCmd.AddCommand(
		versionCmd,
		downloadCmd,
	)
}

// Run executes the CLI application.
func (a *App) Run() error {
	return a.rootCmd.Execute()
}
