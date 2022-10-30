package main

import (
	"context"
	core "github.com/Mahamadou828/AOAC/business/core/v1/admin"
	"net/http"

	"github.com/Mahamadou828/AOAC/business/web/v1"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	sdklambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	sdklambda.Start(web.NewHandler(handler))
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest, cfg *lambda.Config) (events.APIGatewayProxyResponse, error) {
	str, err := core.UploadProfilePicture(ctx, cfg, r)
	if err != nil {
		return lambda.SendError(ctx, http.StatusInternalServerError, err)
	}
	return lambda.Response(ctx, http.StatusOK, struct {
		Resp string `json:"resp"`
	}{
		Resp: str,
	})
}
