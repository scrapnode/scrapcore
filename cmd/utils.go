package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
	"reflect"
)

func ChainPreRunE() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		parent := cmd.Parent()
		if parent == nil {
			return nil
		}

		err := parent.PreRunE(parent, args)
		cmd.SetContext(parent.Context())
		return err
	}
}

func PrintObj(title string, obj any) {
	fmt.Printf("%s\n", title)

	t := table.NewWriter()
	t.AppendHeader(table.Row{"key", "value"})

	col := 80

	v := reflect.ValueOf(obj).Elem()
	for _, f := range reflect.VisibleFields(v.Type()) {
		if f.IsExported() {
			value := v.FieldByName(f.Name)
			t.AppendRow([]interface{}{f.Name, value})
			col = lo.Max([]int{col, len(value.String())})
		}
	}

	t.SetOutputMirror(os.Stdout)
	t.SetAllowedRowLength(lo.Min([]int{col, 160}))
	t.Render()
}

func MustMarkFlagRequired(cmd *cobra.Command, name string) {
	if err := cmd.MarkFlagRequired(name); err != nil {
		panic(err)
	}
}

func MustGetFlagBool(cmd *cobra.Command, name string) bool {
	value, err := cmd.Flags().GetBool(name)
	if err != nil {
		panic(err)
	}
	return value
}

func MustGetFlagInt(cmd *cobra.Command, name string) int {
	value, err := cmd.Flags().GetInt(name)
	if err != nil {
		panic(err)
	}
	return value
}

func MustGetFlagString(cmd *cobra.Command, name string) string {
	value, err := cmd.Flags().GetString(name)
	if err != nil {
		panic(err)
	}
	return value
}

func MustGetFlagStringArray(cmd *cobra.Command, name string) []string {
	values, err := cmd.Flags().GetStringArray(name)
	if err != nil {
		panic(err)
	}
	return values
}
