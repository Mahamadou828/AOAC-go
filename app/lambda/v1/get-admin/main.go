package main

import (
	"context"
	"fmt"
	"github.com/Mahamadou828/AOAC/business/sys/validate"
	"net/http"
	"strconv"

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
	//Get limit parameter
	limitStr, ok := r.QueryStringParameters["limit"]
	if !ok {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("missing required query string parameter limit"))
	}

	//Convert limit to int64
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid number of rows per page"))
	}

	//Get optional parameters last evaluated key parameter
	lastEK := r.QueryStringParameters["lastEvaluatedKey"]

	//Get optional parameters email
	email := r.QueryStringParameters["email"]

	if len(email) > 0 {
		if ok := validate.CheckEmail(email); !ok {
			return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid email: %s", email))
		}
	}

	//find all admin
	res, err := core.Find(ctx, cfg, email, lastEK, limit)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("can't query admin: %v", err))
	}

	return lambda.SendResponse(ctx, http.StatusOK, res)
}
