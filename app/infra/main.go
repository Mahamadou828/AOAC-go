package main

import (
	"errors"
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"log"

	"github.com/Mahamadou828/AOAC/business/sys/aws"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	cognito "github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	s3 "github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	dynamodb "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	identitypool "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha/v2"
)

type InfraStackProps struct {
	awscdk.StackProps
	Env     string
	Service string
}

func NewInfraStack(scope constructs.Construct, id string, props *InfraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	//============================================================ SSM
	//create a pool
	client, err := aws.New(aws.Config{
		ServiceName: props.Service,
		Environment: props.Env,
	})
	if err != nil {
		log.Fatalf("cannot create an aws client: %v", err)
	}

	if err := client.SSM.CreatePool(); err != nil {
		if !errors.Is(err, aws.ErrSSMSecretPoolAlreadyExists) {
			log.Fatalf("can't create pool %v", err)
		}
	}

	//============================================================ S3
	//create the admin picture bucket
	adminBuck := s3.NewBucket(
		stack,
		jsii.String(fmt.Sprintf("admin-profile-picture-%s", props.Env)),
		&s3.BucketProps{
			BucketName: jsii.String(fmt.Sprintf("admin-profile-picture-%s", props.Env)),
		},
	)
	awscdk.NewCfnOutput(stack, jsii.String("s3AdminProfilePictureBucket"), &awscdk.CfnOutputProps{
		Value: adminBuck.BucketName(),
	})

	//create the user picture bucket
	userBucket := s3.NewBucket(
		stack,
		jsii.String(fmt.Sprintf("user-profile-picture-%s", props.Env)),
		&s3.BucketProps{
			BucketName: jsii.String(fmt.Sprintf("user-profile-picture-%s", props.Env)),
		},
	)
	awscdk.NewCfnOutput(stack, jsii.String("s3UserProfilePictureBucket"), &awscdk.CfnOutputProps{
		Value: userBucket.BucketName(),
	})
	//============================================================ Cognito
	//create a cognito pool
	c := cognito.NewUserPool(stack, jsii.String(props.Env+"cognitopool"), &cognito.UserPoolProps{
		UserPoolName:  jsii.String(fmt.Sprintf("%s-%s", props.Service, props.Env)),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		SignInAliases: &cognito.SignInAliases{
			Username:          jsii.Bool(true),
			PreferredUsername: jsii.Bool(true),
		},
		AutoVerify: &cognito.AutoVerifiedAttrs{
			Phone: jsii.Bool(true),
		},
		StandardAttributes: &cognito.StandardAttributes{
			Email: &cognito.StandardAttribute{
				Required: jsii.Bool(true),
				Mutable:  jsii.Bool(true),
			},
			PhoneNumber: &cognito.StandardAttribute{
				Required: jsii.Bool(true),
				Mutable:  jsii.Bool(true),
			},
			Fullname: &cognito.StandardAttribute{
				Required: jsii.Bool(true),
				Mutable:  jsii.Bool(true),
			},
		},
		PasswordPolicy: &cognito.PasswordPolicy{
			MinLength:        jsii.Number(12),
			RequireLowercase: jsii.Bool(true),
			RequireUppercase: jsii.Bool(true),
			RequireDigits:    jsii.Bool(true),
			RequireSymbols:   jsii.Bool(true),
		},
		CustomAttributes: &map[string]cognito.ICustomAttribute{
			"isActive": cognito.NewStringAttribute(&cognito.StringAttributeProps{
				MinLen:  jsii.Number(1),
				MaxLen:  jsii.Number(256),
				Mutable: jsii.Bool(true),
			}),
		},
		AccountRecovery:   cognito.AccountRecovery_PHONE_ONLY_WITHOUT_MFA,
		SelfSignUpEnabled: jsii.Bool(true),
	})

	//Create a new App client
	poolClient := c.AddClient(jsii.String("tgs-api"), &cognito.UserPoolClientOptions{
		AuthFlows: &cognito.AuthFlow{
			AdminUserPassword: jsii.Bool(true),
			Custom:            jsii.Bool(true),
			UserPassword:      jsii.Bool(true),
			UserSrp:           jsii.Bool(true),
		},
		GenerateSecret: jsii.Bool(false),
	})

	awscdk.NewCfnOutput(stack, jsii.String("cognitoUserPoolId"), &awscdk.CfnOutputProps{
		Value: c.UserPoolId(),
	})

	awscdk.NewCfnOutput(stack, jsii.String("cognitoClientPoolId"), &awscdk.CfnOutputProps{
		Value: poolClient.UserPoolClientId(),
	})

	identitypool.NewUserPoolAuthenticationProvider(&identitypool.UserPoolAuthenticationProviderProps{
		UserPool:       c,
		UserPoolClient: poolClient,
	})

	identitypool.NewIdentityPool(stack, jsii.String(props.Env+"identitypool"), &identitypool.IdentityPoolProps{
		AllowUnauthenticatedIdentities: jsii.Bool(true),
		AuthenticationProviders: &identitypool.IdentityPoolAuthenticationProviders{
			UserPools: &[]identitypool.IUserPoolAuthenticationProvider{
				identitypool.NewUserPoolAuthenticationProvider(&identitypool.UserPoolAuthenticationProviderProps{
					UserPool:       c,
					UserPoolClient: poolClient,
				}),
			},
		},
		IdentityPoolName: jsii.String(fmt.Sprintf("%s-%s-identity-pool", props.Service, props.Env)),
	})

	//============================================================ DynamoDB
	//Create the admin table
	adminTab := dynamodb.NewTable(stack, jsii.String(fmt.Sprintf("admin-%s", props.Env)), &dynamodb.TableProps{
		TableName: jsii.String(fmt.Sprintf("%s-admin", props.Env)),
		PartitionKey: &dynamodb.Attribute{
			Name: jsii.String("id"),
			Type: dynamodb.AttributeType_STRING,
		},
	})
	adminTab.AddGlobalSecondaryIndex(&dynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String("emailIndex"),
		PartitionKey: &dynamodb.Attribute{
			Name: jsii.String("email"),
			Type: dynamodb.AttributeType_STRING,
		},
	})
	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewInfraStack(app,
		"testing",
		&InfraStackProps{
			awscdk.StackProps{
				Env: env(),
			},
			"testing",
			"aoac",
		})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Region: jsii.String("eu-west-1"),
	}
}
