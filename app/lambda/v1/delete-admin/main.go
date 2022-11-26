package main

import (
	"context"
	"fmt"
	"github.com/Mahamadou828/AOAC/business/sys/validate"
	"net/http"

	core "github.com/Mahamadou828/AOAC/business/core/v1/admin"
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

	//get admin id
	id, ok := r.PathParameters["id"]
	if !ok {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("missing required parameter: id"))
	}
	if err := validate.CheckID(id); err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid id: %v", err))
	}

	//delete the admin
	newAdmin, err := core.Delete(ctx, cfg, id, v.Now)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("can't create new admin: %v", err))

	}
	return lambda.SendResponse(ctx, http.StatusOK, newAdmin)
}
