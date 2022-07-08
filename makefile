SHELL := bin/bash
ENV := development

#Vendor all the project dependencies.
tidy:
	go mod tidy
	go mod vendor

#================================================================= V1
build-v1:
	sam validate -t config/v1/template.yml
	sam build -t config/v1/template.yml

start-api-v1:
	sam local start-api

start-v1: build-v1 start-api-v1

admin:
	go run app/tools/admin/main.go

test:
	go test -v ./...