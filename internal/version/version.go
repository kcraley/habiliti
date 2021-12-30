package version

import (
	"fmt"
	"runtime"
)

const (
	stringFormat = "{BuildDate: %s, GitVersion: %s, GoOSArch: %s, GoVersion: %s}"
)

var (
	buildDate  = ""
	gitVersion = ""
)

type buildInformation struct {
	BuildDate  string
	GitVersion string
	GoOSArch   string
	GoVersion  string
}

// String returns the build information as a string
func (b *buildInformation) string() string {
	return fmt.Sprintf(
		stringFormat,
		b.BuildDate,
		b.GitVersion,
		b.GoOSArch,
		b.GoVersion,
	)
}

// String returns the build information as a string
func String() string {
	version := buildInformation{
		BuildDate:  buildDate,
		GitVersion: gitVersion,
		GoOSArch:   fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH),
		GoVersion:  runtime.Version(),
	}

	return version.string()
}
