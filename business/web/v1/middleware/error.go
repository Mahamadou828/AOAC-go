package middleware

import (
	"context"
	"github.com/Mahamadou828/AOAC/business/web/v1/sentryfmt"
	"github.com/Mahamadou828/AOAC/foundation/web"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
)

//Errors send error handler error to the client after formatting them into the ErrorResponse.
func Errors() web.Middelware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, r events.APIGatewayProxyRequest, sess *session.Session, secrets map[string]string) (events.APIGatewayProxyResponse, error) {
			rqsResp, rqsErr := handler(ctx, r, sess, secrets)
			v, err := web.GetRequestTrace(ctx)
			if rqsErr != nil || err != nil {
				sentryfmt.CaptureError(v, r, rqsErr)
				rsp := struct {
					Message string `json:"Message"`
				}{
					Message: rqsErr.Error(),
				}

				return web.Response(ctx, 400, rsp)
			}
			return rqsResp, nil
		}

		return h
	}
	return m
}
