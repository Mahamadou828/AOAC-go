package main

import (
	"context"
	"fmt"
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
	rowsPerPageStr, ok := r.QueryStringParameters["rowsPerPage"]
	if !ok {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("missing required query string parameter rowsPerPage"))
	}
	pageStr, ok := r.QueryStringParameters["page"]
	if !ok {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("missing required query string parameter page"))
	}
	rowsPerPage, err := strconv.Atoi(rowsPerPageStr)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid number of rows per page"))
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("invalid number of page"))
	}

	res, err := core.Query(ctx, cfg, page, rowsPerPage)
	if err != nil {
		return lambda.SendError(ctx, http.StatusBadRequest, fmt.Errorf("can't query admin: %v", err))
	}
	return lambda.Response(ctx, http.StatusOK, res)
}
