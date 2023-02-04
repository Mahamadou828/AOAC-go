package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	core "github.com/Mahamadou828/AOAC/business/core/v1/university"
	"github.com/Mahamadou828/AOAC/business/web/v1"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	sdklambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	sdklambda.Start(web.NewHandler(handler))
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest, cfg *lambda.Config) (events.APIGatewayProxyResponse, error) {
	//Get parameters limit
	limitStr, ok := r.QueryStringParameters["limit"]
	if !ok {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("missing required query string parameter limit"))
	}

	//Parse limit into a int64
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid parameter limit: %v", err))
	}

	//Extract optional parameters last evaluated key
	lastEvaluatedKey := r.QueryStringParameters["lastEvaluatedKey"]

	//Extract optional parameters country
	country := r.QueryStringParameters["country"]

	//find all university
	us, err := core.Find(ctx, cfg, country, lastEvaluatedKey, limit)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("can't retrieve university: %v", err))
	}

	return lambda.SendResponse(ctx, http.StatusOK, us)
}
