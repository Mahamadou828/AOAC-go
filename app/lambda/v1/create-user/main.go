package main

import (
	"context"
	"fmt"
	"net/http"

	core "github.com/Mahamadou828/AOAC/business/core/v1/user"
	"github.com/Mahamadou828/AOAC/business/web/v1"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	sdklambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	sdklambda.Start(web.NewHandler(handler))
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest, cfg *lambda.Config) (events.APIGatewayProxyResponse, error) {
	//Get request trace
	v, err := lambda.GetRequestTrace(ctx)
	if err != nil {
		return lambda.SendResponse(ctx, http.StatusInternalServerError, fmt.Errorf("unable to get request trace: %v", err))
	}

	//Create a new user
	newUser, err := core.Create(ctx, cfg, r, v.Now)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("can't create user: %v", err))
	}

	return lambda.SendResponse(ctx, http.StatusOK, newUser)
}
