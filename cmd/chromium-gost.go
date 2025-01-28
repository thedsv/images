package cmd

import (
	"github.com/thedsv/images/build"
	"github.com/spf13/cobra"
)

var (
	chromiumGostCmd = &cobra.Command{
		Use:   "chromium-gost",
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
			chromiumGost := &build.ChromiumGost{req}
			return chromiumGost.Build()
		},
	}
)
