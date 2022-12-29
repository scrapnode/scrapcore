package cmd

import (
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/spf13/cobra"
)

func NewShow() *cobra.Command {
	command := &cobra.Command{
		Use:               "show",
		Short:             "show information of your app",
		PersistentPreRunE: ChainPreRunE(),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			logger := xlogger.FromContext(ctx).With("fn", "cli")

			var cfg xconfig.Configs
			if err := cfg.Unmarshal(xconfig.FromContext(ctx)); err != nil {
				logger.Fatal(err)
			}

			PrintObj("---***---", cfg)
		},
	}

	return command
}
