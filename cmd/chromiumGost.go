package cmd

import (
	"github.com/aerokube/images/build"
	"github.com/spf13/cobra"
)

var (
	gostCmd = &cobra.Command{
		Use:   "gost",
		Short: "build Chromium-gost image",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := build.Requirements{
				BrowserSource:  build.BrowserSource(browserSource),
				BrowserChannel: browserChannel,
				DriverVersion:  driverVersion,
				NoCache:        noCache,
				TestsDir:       testsDir,
				RunTests:       test,
				IgnoreTests:    ignoreTests,
				Tags:           tags,
				PushImage:      push,
			}
			gost := &build.Gost{req}
			return gost.Build()
		},
	}
)
