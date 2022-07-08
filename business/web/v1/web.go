package web

import (
	"context"
	"github.com/Mahamadou828/AOAC/business/sys/aws"
	"github.com/Mahamadou828/AOAC/business/web/v1/middleware"
	"os"
	"time"

	"github.com/Mahamadou828/AOAC/foundation/web"
	"github.com/aws/aws-lambda-go/events"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
)

const service = "aoac"

type LambdaHandler func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

//NewHandler create a new lambda handler
func NewHandler(handler web.Handler) LambdaHandler {
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
	client, err := aws.New(service, os.Getenv("ENV"))
	if err != nil {
		panic(err)
	}

	//Create context for the handler
	hub := sentry.CurrentHub().Clone()
	v := web.RequestTrace{
		ID:         uuid.NewString(),
		Now:        time.Now().UTC(),
		StatusCode: 0,
		Hub:        hub,
	}

	ctx := context.WithValue(context.Background(), web.CtxKey, &v)

	//wrap the handler with all the middlewares
	h := web.WrapMiddleware(handler, middleware.Errors())

	lambda := func(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		defer sentry.Flush(2 * time.Second)
		return h(ctx, r, client)
	}

	//call the handler
	return lambda
}
