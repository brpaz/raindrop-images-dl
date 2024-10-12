package cmd_test

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brpaz/raindrop-images-dl/internal/cmd"
)

func TestVersionCmd(t *testing.T) {
	t.Parallel()

	versionInfo := cmd.VersionInfo{
		Version:   "v0.1.0",
		GitCommit: "abcdef",
		BuildDate: "2021-01-01T00:00:00Z",
	}

	w := &bytes.Buffer{}
	cmd := cmd.NewVersionCmd(versionInfo)
	cmd.SetOut(w)
	assert.Equal(t, "version", cmd.Use)
	assert.Equal(t, "Display version information", cmd.Short)

	cmd.Run(cmd, []string{})
	expectedOutput := fmt.Sprintf("Version: %s\nGit Commit: %s\nBuild Date: %s\nGo Version: %s\nOS/Arch: %s/%s\n",
		versionInfo.Version,
		versionInfo.GitCommit,
		versionInfo.BuildDate,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)

	assert.Equal(t, expectedOutput, w.String())
}
