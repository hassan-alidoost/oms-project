package app

import (
	"fmt"
	"log"

	"github.com/hassan-alidoost/oms-project/config"
	"github.com/hassan-alidoost/oms-project/internal/infra/postgres"
)


func Initialize(cfg *config.Config) {
	fmt.Printf("start running app in %s enviorment", cfg.App.Env)
	
	db, cleanup, err := postgres.NewPostgresDB(cfg.Database)

	if err != nil {
        log.Fatalf("could not set up database: %v", err)
    }

	defer cleanup()

}