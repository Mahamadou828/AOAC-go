package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/Mahamadou828/AOAC/app/tools/config"
	"github.com/Mahamadou828/AOAC/business/sys/aws"
)

type Secret struct {
	Name  string
	Value string
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

const service = "SECRET_PARSER"

func main() {
	cfg := struct {
		Service  string `conf:"required"`
		FilePath string `conf:"required"`
	}{}
	if _, err := config.Parse(&cfg, service, nil); err != nil {
		log.Fatalf("failed to parse configuration %v", err)
	}

	jsonFile, err := os.Open(cfg.FilePath)
	if err != nil {
		log.Fatalf("failed to open file %v", err)
	}
	defer jsonFile.Close()

	b, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("failed to read file %v", err)
	}

	var out map[string]map[string]string
	if err := json.Unmarshal(b, &out); err != nil {
		log.Fatalf("failed to unmarshal file %v", err)
	}

	for env, secretMap := range out {
		var secrets []aws.Secret
		log.Println("creating client for env", "env", env)
		log.Println("starting to create secret")
		client, err := aws.New(aws.Config{
			ServiceName: cfg.Service,
			Environment: env,
		})
		if err != nil {
			log.Printf("failed to create an client for env %s: %v", env, err)
			continue
		}
		for key, val := range secretMap {
			secrets = append(secrets, aws.Secret{Name: ToSnake(key), Value: val})
		}
		log.Println("creating secret inside pool", "poolname", fmt.Sprintf("%s/%s", cfg.Service, env))
		err = client.SSM.CreateSecrets(secrets)
		if err != nil {
			log.Printf("can't create database secrets: %v", err)
			os.Exit(1)
		}
	}
}

func ToSnake(camel string) (snake string) {
	snake = matchAllCap.ReplaceAllString(matchFirstCap.ReplaceAllString(camel, "${1}_${2}"), "${1}_${2}")
	return strings.ToUpper(snake)
}
