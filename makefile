SHELL := bin/bash
ENV := development

#================================================================= V1
build-v1:
	sam validate -t config/v1/template.yml
	sam build -t config/v1/template.yml

start-api-v1:
	sam local start-api

start-v1: build-v1 start-api-v1

test:
	go test -v ./...