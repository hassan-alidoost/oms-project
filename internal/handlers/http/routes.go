package http

import (
	"net/http"

	httporder "github.com/hassan-alidoost/oms-project/internal/handlers/http/order"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handlers struct {
	OrderHandler *httporder.OrderHandler
}

func NewRouter(h Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	//swagger
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	//order
	mux.HandleFunc("POST /orders", h.OrderHandler.CreateOrder)
	mux.HandleFunc("GET /orders", h.OrderHandler.ListOrders)
	mux.HandleFunc("GET /orders/{id}", h.OrderHandler.GetOrder)
	mux.HandleFunc("PUT /orders/{id}", h.OrderHandler.UpdateStatus)

	return mux
}
