package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hassan-alidoost/oms-project/config"
	orderHandler "github.com/hassan-alidoost/oms-project/internal/handlers/http/order"
	"github.com/hassan-alidoost/oms-project/internal/infra/postgres"
	"github.com/hassan-alidoost/oms-project/internal/infra/postgres/repository"
	"github.com/hassan-alidoost/oms-project/internal/services"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/hassan-alidoost/oms-project/docs"
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

	orderRepository := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepository)
	orderHandler := orderHandler.NewOrderHandler(orderService)
	
	mux := http.NewServeMux()
	orderHandler.RegisterRoutes(mux)

	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	log.Println("OMS API starting on :8080")
    http.ListenAndServe(":8080", mux)
}