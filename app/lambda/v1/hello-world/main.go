package main

import (
	"github.com/Mahamadou828/AOAC/business/core/v1/hello"
	"github.com/Mahamadou828/AOAC/business/web/v1"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(
		web.NewHandler(
			hello.Hello,
			"AOAC_API",
			nil,
		),
	)
}
