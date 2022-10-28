package main

import (
	"context"
	"fmt"
	"net/http"

	core "github.com/Mahamadou828/AOAC/business/core/v1/admin"
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
	//get and check the provided id
	id, ok := r.PathParameters["id"]
	if !ok {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("missing required parameter 'id'"))
	}
	if err := validate.CheckID(id); err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid id: %v", err))
	}

	//fetch the admin by the provided id
	admin, err := core.QueryByID(ctx, cfg, id)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("can't find admin: %v", err))
	}

	return lambda.Response(ctx, http.StatusOK, admin)
}
