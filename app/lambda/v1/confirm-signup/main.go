package main

import (
	"context"
	"fmt"
	core "github.com/Mahamadou828/AOAC/business/core/v1/admin"
	"github.com/Mahamadou828/AOAC/business/sys/validate"
	"github.com/Mahamadou828/AOAC/business/web/v1"
	"net/http"

	model "github.com/Mahamadou828/AOAC/business/data/v1/models/admin"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	sdklambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	sdklambda.Start(web.NewHandler(handler))
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest, cfg *lambda.Config) (events.APIGatewayProxyResponse, error) {
	//unmarshal the request body
	var data model.ConfirmSignupDTO
	if err := lambda.DecodeBody(r.Body, &data); err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("can't decode body: %v", err))
	}
	//validate the request body
	if err := validate.Check(data); err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
	}
	//confirm the account
	if err := core.ConfirmSignUp(data, cfg); err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("can't confirm signature: %v", err))
	}

	return lambda.SendResponse(ctx, http.StatusNoContent, nil)
}
