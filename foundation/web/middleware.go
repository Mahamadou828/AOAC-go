package web

import (
	"context"
	"github.com/Mahamadou828/AOAC/business/sys/aws"
	"github.com/aws/aws-lambda-go/events"
)

type Handler func(ctx context.Context, request events.APIGatewayProxyRequest, client *aws.Client) (events.APIGatewayProxyResponse, error)

type Middelware func(h Handler) Handler

// WrapMiddleware A Middleware is a function designed to run some code before and/or after
//another Handler. It is designed to remove boilerplate or other concerns not
//direct to any given Handler
func WrapMiddleware(handler Handler, mw ...Middelware) Handler {
	// Loop backwards through the middleware invoking each one. Replace the
	// handler with the new wrapped handler. Looping backwards ensures that the
	// first middleware of the slice is the first to be executed by requests.
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}
