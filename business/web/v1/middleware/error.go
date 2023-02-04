package middleware

import (
	"context"
	"github.com/Mahamadou828/AOAC/business/web/v1/sentryfmt"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
)

// Errors send error handler error to the client after formatting them into the ErrorResponse.
func Errors() lambda.Middelware {
	m := func(handler lambda.Handler) lambda.Handler {
		h := func(ctx context.Context, r events.APIGatewayProxyRequest, cfg *lambda.Config) (events.APIGatewayProxyResponse, error) {
			rqsResp, rqsErr := handler(ctx, r, cfg)
			v, err := lambda.GetRequestTrace(ctx)
			if rqsErr != nil || err != nil {
				sentryfmt.CaptureError(v, r, rqsErr)
				rsp := struct {
					Message string `json:"Message"`
				}{
					Message: rqsErr.Error(),
				}

				return lambda.SendResponse(ctx, 400, rsp)
			}
			return rqsResp, nil
		}

		return h
	}
	return m
}
