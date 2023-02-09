package main

import (
	"context"
	"fmt"
	"net/http"

	core "github.com/Mahamadou828/AOAC/business/core/v1/user"
	"github.com/Mahamadou828/AOAC/business/data/v1/models/user"
	"github.com/Mahamadou828/AOAC/business/sys/validate"
	"github.com/Mahamadou828/AOAC/business/web/v1"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	sdklambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	sdklambda.Start(web.NewHandler(handler))
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest, cfg *lambda.Config) (events.APIGatewayProxyResponse, error) {
	//Unmarshal body request and verify it
	var data user.LoginUserDTO
	if err := lambda.DecodeBody(r.Body, &data); err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("unable to unmarshal body: %v", err))
	}
	if err := validate.Check(data); err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid body: %v", err))
	}

	//Get a new session
	session, err := core.Login(ctx, cfg, data)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("can't login an admin: %v", err))

	}

	return lambda.SendResponse(ctx, http.StatusOK, session)
}
