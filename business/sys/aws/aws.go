package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Config struct {
	ServiceName string
	Environment string

	CognitoUserPoolID string
	CognitoClientID   string
}

// Client provides an api to interact with all aws services.
type Client struct {
	env     string
	service string
	SSM     *SSM
	Sess    *session.Session
	Cognito *Cognito
}

func New(cfg Config) (*Client, error) {
	//Initiate a new aws session
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String("eu-west-1"),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error creating session: %v", err)
	}

	return &Client{
		env:     cfg.Environment,
		service: cfg.ServiceName,
		SSM:     NewSSM(cfg.ServiceName, cfg.Environment, sess),
		Cognito: NewCognito(sess, cfg.CognitoClientID, cfg.CognitoUserPoolID),
		Sess:    sess,
	}, nil
}
