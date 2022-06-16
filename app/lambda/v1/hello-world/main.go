package main

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if len(ip) == 0 {
		return events.APIGatewayProxyResponse{}, ErrNoIP
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string(ip)),
		StatusCode: 200,
	}, nil
}

func main() {

	//========================================= Logging
	//Configuring sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://f272b793754e449c88bd630f8ee06f05@o1236486.ingest.sentry.io/6509179",
		TracesSampleRate: 1.0,
		AttachStacktrace: true,
		Debug:            false,
	}); err != nil {
		panic(err)
	}

	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
	lambda.Start(handler)
}
