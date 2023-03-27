package config

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"os"
	"strings"
)

var repository Repository

type RepositoryConfig struct {
	Name string
}

type Repository struct {
	config RepositoryConfig
	ecr    ecr.ECR
}

type Credentials struct {
	Username string
	Password string
}

type AuthorizationData struct {
	Credentials Credentials
	Url         string
}

func CreateRepository(config RepositoryConfig) {
	newSession, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})

	if err != nil {
		fmt.Println("Error while creating new aws session", err)
	}

	repository = Repository{
		config: config,
		ecr:    *ecr.New(newSession),
	}
}

func (repository *Repository) GetRepositoryName() string {
	return repository.config.Name
}

func (repository *Repository) GetAuthorizationToken() AuthorizationData {
	token, err := repository.ecr.GetAuthorizationToken(&ecr.GetAuthorizationTokenInput{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := token.AuthorizationData[0]
	decodeString, err := base64.StdEncoding.DecodeString(*data.AuthorizationToken)

	if err != nil {
		os.Exit(1)
	}

	credentials := strings.Split(string(decodeString), ":")

	return AuthorizationData{
		Url: *data.ProxyEndpoint,
		Credentials: Credentials{
			Username: credentials[0],
			Password: credentials[1],
		},
	}
}

func GetRepository() Repository {
	return repository
}
