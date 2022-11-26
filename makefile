SHELL := bin/bash
ENV := development

#Vendor all the project dependencies.
tidy:
	go mod tidy
	go mod vendor

#================================================================= V1
deploy-cdk-stack:
	cdk deploy --all -O ./infra.output.spec.json
	go run app/tools/cfnparser/main.go --file-path=infra.output.spec.json --service=$(SERVICE)

destroy-cdk-stack:
	cdk destroy --all

build-v1:
	sam validate -t config/v1/template.yml
	sam build -t config/v1/template.yml

start-api-v1:
	sam local start-api --env-vars env.local.json

start: build-v1 start-api-v1

build: build-v1

validate:
	sam validate -t config/v1/template.yml

admin:
	go run app/tools/admin/main.go

test:
	go test -v ./...

deploy: deploy-cdk-stack

destroy: destroy-cdk-stack

scrap:
	go run app/scraper/main.go