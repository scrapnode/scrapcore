package cmd

import (
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/spf13/cobra"
	"strings"
)

func New() *cobra.Command {
	command := &cobra.Command{
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			dirs, err := cmd.Flags().GetStringArray("configs-dirs")
			if err != nil {
				return err
			}
			provider, err := xconfig.New(dirs...)
			sets, err := cmd.Flags().GetStringArray("set")
			if err != nil {
				return err
			}
			for _, s := range sets {
				kv := strings.Split(s, "=")
				provider.Set(kv[0], kv[1])
			}
			ctx = xconfig.WithContext(ctx, provider)

			// change context to our new context
			cmd.SetContext(ctx)
			return nil
		},
		ValidArgs: []string{},
	}

	command.PersistentFlags().StringArrayP(
		"configs-dirs", "c",
		[]string{".", "./secrets"}, "path/to/config/file",
	)
	command.PersistentFlags().StringArrayP(
		"set", "s",
		[]string{}, "override values in config file",
	)

	command.AddCommand(NewShow())
	return command
}
