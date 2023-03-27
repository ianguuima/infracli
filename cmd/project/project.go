package project

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects inside your infrastructure",
	Long:  "Manage projects inside your infrastructure",
}

func init() {
	Cmd.AddCommand(uploadDockerComposeCmd)
	Cmd.AddCommand(createProjectCmd)
}
