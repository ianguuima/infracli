package project

import "github.com/spf13/cobra"

var uploadDockerComposeCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload docker compose structure",
	Long:  "Upload docker compose structure",
	Run: func(cmd *cobra.Command, args []string) {
		println("Ok!!")
	},
}
