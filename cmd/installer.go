package cmd

import (
	"github.com/spf13/cobra"
)

var installer = &cobra.Command{
	Use:   "install",
	Short: "Install different scripts(bash,fish,powershell) to run wintree-cd. A command that changes directory base on the dir/file you clicked\n wintree install bash\n wintree-cd {path} ",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return installShell(args[0])

	},
}

func init() {
	rootCmd.AddCommand(installer)
}
