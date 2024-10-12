package app

import (
	"github.com/spf13/cobra"

	"github.com/brpaz/raindrop-images-dl/internal/cmd"
	"github.com/brpaz/raindrop-images-dl/internal/version"
)

// App represents an instance of the CLI application.
type App struct {
	rootCmd *cobra.Command
}

// New creates a new instance of the application.
func New() *App {
	app := &App{
		rootCmd: cmd.NewRootCmd(),
	}

	app.registerCommands()

	return app
}

func (a *App) registerCommands() {
	versionCmd := cmd.NewVersionCmd(cmd.VersionInfo{
		Version:   version.Version,
		GitCommit: version.GitCommit,
		BuildDate: version.BuildDate,
	})

	downloadCmd := cmd.NewDownloadCmd()

	a.rootCmd.AddCommand(
		versionCmd,
		downloadCmd,
	)
}

// Run executes the CLI application.
func (a *App) Run() error {
	return a.rootCmd.Execute()
}
