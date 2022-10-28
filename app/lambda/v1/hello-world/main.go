package main

import (
	"context"
	"fmt"

	"github.com/Mahamadou828/AOAC/app/tools/config"
	"github.com/Mahamadou828/AOAC/business/core/v1/hello"
	"github.com/Mahamadou828/AOAC/business/web/v1"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	sdklambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	sdklambda.Start(web.NewHandler(handler))
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest, cfg *lambda.Config) (events.APIGatewayProxyResponse, error) {
	data := struct {
		Env       string `conf:"env:ENV,required"`
		SecretEnv string `conf:"env:SECRET_ENV,required"`
	}{}

	if err := config.ParseLambdaCfg(&data); err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("can't parse config: %v", err)
	}

	resp, err := hello.Hello(data.Env, data.SecretEnv, r.Path)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return lambda.Response(ctx, 200, resp)
}
