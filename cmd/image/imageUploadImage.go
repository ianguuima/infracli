package image

import (
	"fmt"
	"github.com/spf13/cobra"
	"infracli/config"
	"os"
	"os/exec"
	"strings"
)

var uploadImageCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a image to the project",
	Long:  "Upload a image to the project",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Tag can be project:1.0
		imageTag := generateImage(args)
		uploadImage(imageTag)
	},
}

func uploadImage(imageTag string) {
	repository := config.GetRepository()
	token := repository.GetAuthorizationToken()

	credentials := token.Credentials

	execCommand(exec.Command("docker", "login", "-u", credentials.Username, "-p", credentials.Password, token.Url))

	version := strings.Split(imageTag, ":")[1]
	strippedUrl := strings.Replace(token.Url, "https://", "", -1)
	remoteTag := fmt.Sprintf("%s/%s:%s", strippedUrl, repository.GetRepositoryName(), version)

	execCommand(exec.Command("docker", "tag", imageTag, remoteTag))
	execCommand(exec.Command("docker", "push", remoteTag))
}

func generateImage(args []string) string {
	tag, version := getImageTagAndVersion(args[0])
	path := args[1]

	execCommand(exec.Command("docker", "build", "-t", fmt.Sprintf("%s:%s", tag, version), path))

	return strings.Join([]string{tag, version}, ":")
}

func getImageTagAndVersion(arg string) (string, string) {
	tag := strings.Split(arg, ":")
	imageName := tag[0]
	version := tag[1]

	return imageName, version
}

func execCommand(cmd *exec.Cmd) {
	command := cmd
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Println("could not run command: ", err.Error())
		os.Exit(1)
	}
}
