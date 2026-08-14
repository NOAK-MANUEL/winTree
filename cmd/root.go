package cmd

import (
	"log"
	"winTree/bubble_modal"
	globarvar "winTree/globarVar"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wintree",
	Args:  cobra.ExactArgs(1),
	Short: "winTree is a terminal file explorer in a tree form",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := validPath(args[0])
		err := bubble_modal.ExecuteBubble(path)
		return err
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&globarvar.JustDir, "dir", "d", false, "Get only directories in the tree")
	rootCmd.Flags().BoolVarP(&globarvar.JustFile, "file", "f", false, "Get only files in the tree")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
