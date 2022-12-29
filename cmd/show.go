package cmd

import (
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/spf13/cobra"
)

func NewShow() *cobra.Command {
	command := &cobra.Command{
		Use:               "show",
		Short:             "show information of your app",
		PersistentPreRunE: ChainPreRunE(),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			var cfg xconfig.Configs
			if err := cfg.Unmarshal(xconfig.FromContext(ctx)); err != nil {
				return err
			}

			PrintObj("---***---", &cfg)
			return nil
		},
	}

	return command
}
