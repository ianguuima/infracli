package project

import "github.com/spf13/cobra"

var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create project",
	Long:  "Create project",
	Run: func(cmd *cobra.Command, args []string) {
		println("Ok!!")
	},
}
