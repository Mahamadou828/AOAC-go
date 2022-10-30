package hello

import (
	"errors"
	"fmt"
	"io"
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

type RspHello struct {
	IP        string `json:"ip"`
	Path      string `json:"path"`
	Env       string `json:"env"`
	SecretEnv string `json:"secretEnv"`
}

func Hello(env, secretEnv, path string) (RspHello, error) {
	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return RspHello{}, fmt.Errorf("error getting HTTP response: %v", err)
	}

	if resp.StatusCode != 200 {
		return RspHello{}, ErrNon200Response
	}

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return RspHello{}, err
	}

	if len(ip) == 0 {
		return RspHello{}, ErrNoIP
	}

	return RspHello{string(ip), path, env, secretEnv}, nil
}
