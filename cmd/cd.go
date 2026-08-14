package cmd

import (
	"winTree/bubble_modal"
	globarvar "winTree/globarVar"

	"github.com/spf13/cobra"
)

var cd = &cobra.Command{
	Use:   "cd",
	Short: "Change current terminal directory to the folder/file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		globarvar.ON_CD_COMMAND = true
		path := validPath(args[0])
		bubble_modal.ExecuteBubble(path)
	},
}

func init() {
	rootCmd.AddCommand(cd)
	cd.Flags().BoolVarP(&globarvar.JustDir, "dir", "d", false, "Get only directories in the tree")
	cd.Flags().BoolVarP(&globarvar.JustFile, "file", "f", false, "Get only files in the tree")
}
