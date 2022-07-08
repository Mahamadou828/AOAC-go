package main

import (
	"errors"
	"fmt"
	"github.com/Mahamadou828/AOAC/app/tools/admin/commands"
	"github.com/Mahamadou828/AOAC/app/tools/config"
	"github.com/Mahamadou828/AOAC/business/sys/aws"
	"os"
)

var availableCommands = []string{
	"1. SSM - create secret",
	"2. SSM - create pool",
}

func main() {
	fmt.Println("Starting admin tools")

	cfg := struct {
		Env     string `conf:"help:which env are you targeting ? [testing development staging production local]"`
		Service string `conf:"help:which services are you targeting"`
	}{}

	if err := config.ParseAdminCfg(&cfg); err != nil {
		fmt.Printf("can't parse admin config %v", err)
		os.Exit(1)
	}

	fmt.Println("Which commands do you want to run")
	for _, cmd := range availableCommands {
		fmt.Println(cmd)
	}
	var cmd int
	fmt.Printf("Choice: ")
	if _, err := fmt.Scan(&cmd); err != nil {
		fmt.Printf("stopping admin tools: %v", err)
		os.Exit(1)
	}

	client, err := aws.New(cfg.Service, cfg.Env)
	if err != nil {
		fmt.Printf("can't initialized aws session %v", err)
	}

	switch cmd {
	case 1:
		err = commands.SSMCreateSecret(client)
	case 2:
		err = commands.SSMCreatePool(client)
	default:
		err = errors.New("unknown command")
	}

	if err != nil {
		fmt.Printf("error with the choosen commands %v", err)
		os.Exit(1)
	}

	return
}
