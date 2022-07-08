package hello

import (
	"context"
	"errors"
	"fmt"
	"github.com/Mahamadou828/AOAC/app/tools/config"
	"github.com/Mahamadou828/AOAC/business/sys/aws"
	"github.com/Mahamadou828/AOAC/foundation/web"
	"github.com/aws/aws-lambda-go/events"
	"io/ioutil"
	"net/http"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func Hello(ctx context.Context, r events.APIGatewayProxyRequest, client *aws.Client) (events.APIGatewayProxyResponse, error) {
	cfg := struct {
		Env       string `conf:"env:ENV,required"`
		SecretEnv string `conf:"env:SECRET_ENV,required"`
	}{}

	if err := config.ParseLambdaCfg(&cfg); err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("can't parse config: %v", err)
	}

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

	data := struct {
		IP        string `json:"ip"`
		Path      string `json:"path"`
		Env       string `json:"env"`
		SecretEnv string `json:"secretEnv"`
	}{
		IP:        string(ip),
		Path:      r.Path,
		Env:       cfg.Env,
		SecretEnv: cfg.SecretEnv,
	}

	return web.Response(ctx, 200, data)
}
