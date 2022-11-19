package web

import (
	"context"
	"github.com/Mahamadou828/AOAC/business/sys/database"
	"os"
	"time"

	"github.com/Mahamadou828/AOAC/business/sys/aws"
	"github.com/Mahamadou828/AOAC/business/web/v1/middleware"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
)

const service = "aoac"

type LambdaHandler func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// NewHandler create a new lambda handler
func NewHandler(handler lambda.Handler) LambdaHandler {
	//Init sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://f272b793754e449c88bd630f8ee06f05@o1236486.ingest.sentry.io/6509179",
		TracesSampleRate: 1.0,
		AttachStacktrace: true,
		Debug:            false,
	}); err != nil {
		panic(err)
	}

	//Init a new aws sess
	awsCfg := aws.Config{
		ServiceName:       service,
		Environment:       os.Getenv("ENV"),
		CognitoUserPoolID: os.Getenv("COGNITO_USER_POOL_ID"),
		CognitoClientID:   os.Getenv("COGNITO_CLIENT_ID"),
	}
	client, err := aws.New(awsCfg)
	if err != nil {
		panic(err)
	}

	//Create context for the handler
	hub := sentry.CurrentHub().Clone()
	v := lambda.RequestTrace{
		ID:         uuid.NewString(),
		Now:        time.Now().UTC(),
		StatusCode: 0,
		Hub:        hub,
	}

	ctx := context.WithValue(context.Background(), lambda.CtxKey, &v)

	//wrap the handler with all the middlewares
	h := lambda.WrapMiddleware(handler, middleware.Errors())

	cfg := lambda.Config{
		AWSClient: client,
		Db:        database.Open(client, os.Getenv("ENV")),
		Env:       os.Getenv("ENV"),
	}

	lmHandler := func(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		defer sentry.Flush(2 * time.Second)
		return h(ctx, r, &cfg)
	}

	//call the handler
	return lmHandler
}
