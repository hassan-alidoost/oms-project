package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hassan-alidoost/oms-project/config"
	_ "github.com/hassan-alidoost/oms-project/docs"
	appHttp "github.com/hassan-alidoost/oms-project/internal/handlers/http"
	"github.com/hassan-alidoost/oms-project/internal/handlers/http/order"
	"github.com/hassan-alidoost/oms-project/internal/infra/postgres"
	"github.com/hassan-alidoost/oms-project/internal/infra/postgres/repository"
	"github.com/hassan-alidoost/oms-project/internal/services"
)

const configPath string = "./../../config"

func Initialize() {

	cfg, err := config.LoadConfig(configPath)

	if err != nil {
		log.Fatalf("could not load the config %v", err)
	}

	fmt.Printf("start running app in %s enviorment", cfg.App.Env)

	db, cleanup, err := postgres.NewPostgresDB(cfg.Database)

	if err != nil {
		log.Fatalf("could not set up database: %v", err)
	}

	defer cleanup()

	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	orderSvc := services.NewOrderService(orderRepo, productRepo)
	orderHandler := order.NewOrderHandler(orderSvc)

	router := appHttp.NewRouter(appHttp.Handlers{
		OrderHandler: orderHandler,
	})

	log.Println("OMS API starting on :8080")
	http.ListenAndServe(":8080", router)
}
