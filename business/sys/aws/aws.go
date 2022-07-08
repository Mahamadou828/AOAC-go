package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

//Client provides an api to interact with all aws services.
type Client struct {
	env     string
	service string
	SSM     *SSM
}

func New(sv, env string) (*Client, error) {
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
		env:     env,
		service: sv,
		SSM:     NewSSM(sv, env, sess),
	}, nil
}
