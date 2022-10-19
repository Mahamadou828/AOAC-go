package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Mahamadou828/AOAC/business/sys/validate"
	"net/http"

	core "github.com/Mahamadou828/AOAC/business/core/v1/admin"
	"github.com/Mahamadou828/AOAC/business/data/v1/models/admin"
	"github.com/Mahamadou828/AOAC/business/web/v1"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	sdklambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	sdklambda.Start(web.NewHandler(handler))
}

func handler(ctx context.Context, r events.APIGatewayProxyRequest, cfg *lambda.Config) (events.APIGatewayProxyResponse, error) {
	//Get request trace
	v, err := lambda.GetRequestTrace(ctx)
	if err != nil {
		return lambda.Response(ctx, http.StatusInternalServerError, fmt.Errorf("unable to get request trace: %v", err))
	}
	//Unmarshal body request and verify it
	var data admin.NewAdminDTO
	if err := json.Unmarshal([]byte(r.Body), &data); err != nil {
		return lambda.Response(ctx, http.StatusBadRequest, fmt.Errorf("unable to unmarshal body: %v", err))
	}
	if err := validate.Check(data); err != nil {
		return lambda.Response(ctx, http.StatusBadRequest, fmt.Errorf("invalid body: %v", err))
	}
	//Create new admin
	newAdmin, err := core.Create(ctx, cfg, data, v.Now)
	if err != nil {
		return lambda.Response(ctx, http.StatusBadRequest, fmt.Errorf("can't create new admin: %v", err))

	}
	return lambda.Response(ctx, http.StatusOK, newAdmin)
}
