package image

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "image",
	Short: "Manage images",
	Long:  "Manage images",
}

func init() {
	Cmd.AddCommand(uploadImageCmd)
}
