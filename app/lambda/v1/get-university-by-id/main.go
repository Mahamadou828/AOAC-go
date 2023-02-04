package main

import (
	"context"
	"fmt"
	core "github.com/Mahamadou828/AOAC/business/core/v1/university"
	"github.com/Mahamadou828/AOAC/business/sys/validate"
	"github.com/Mahamadou828/AOAC/business/web/v1"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	sdklambda "github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func main() {
	sdklambda.Start(web.NewHandler(handler))
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest, cfg *lambda.Config) (events.APIGatewayProxyResponse, error) {
	//Extract optional parameters country
	id, ok := r.PathParameters["id"]
	if !ok {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("missing required parameter 'id'"))
	}
	if err := validate.CheckID(id); err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid id: %s", id))
	}

	u, err := core.FindByID(ctx, cfg, id)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("can't retrieve university: %v", err))
	}

	return lambda.SendResponse(ctx, http.StatusOK, u)
}
