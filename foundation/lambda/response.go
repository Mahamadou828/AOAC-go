package lambda

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

// Response send a response to the client in json
func Response(ctx context.Context, status int, data any) (events.APIGatewayProxyResponse, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"TraceID":      GetTraceID(ctx),
		},
		Body: string(b),
	}, nil
}
