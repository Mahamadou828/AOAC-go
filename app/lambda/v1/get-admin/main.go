package main

import (
	"context"
	"fmt"
	"github.com/Mahamadou828/AOAC/business/core/v1/admin"
	"github.com/Mahamadou828/AOAC/business/web/v1"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	sdklambda "github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func main() {
	fmt.Println("test")
	sdklambda.Start(web.NewHandler(handler))
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest, cfg *lambda.Config) (events.APIGatewayProxyResponse, error) {
	err := admin.Query(ctx, cfg)
	if err != nil {
		return lambda.Response(ctx, http.StatusInternalServerError, err)
	}
	return lambda.Response(ctx, http.StatusOK, "StatusOK")
}
