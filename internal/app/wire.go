package app

import (
	"fmt"
	"log"

	"github.com/hassan-alidoost/oms-project/config"
)

const configPath string = "./../../config"

func Initialize() {
	cfg, err := config.LoadConfig(configPath)

	if err != nil {
		log.Fatalf("could not load the config %v", err)
	}

	fmt.Printf("start running app in %s enviorment", cfg.App.Env)
}