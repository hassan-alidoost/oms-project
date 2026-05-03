package main

import (
	"log"

	"github.com/hassan-alidoost/oms-project/config"
	"github.com/hassan-alidoost/oms-project/internal/app"
)

const configPath string = "./../../config"

// @title OMS API
// @version 1.0
// @description Order Management System with Hexagonal Architecture.
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.LoadConfig(configPath)

	if err != nil {
		log.Fatalf("could not load the config %v", err)
	}

	app.Initialize(cfg)
}