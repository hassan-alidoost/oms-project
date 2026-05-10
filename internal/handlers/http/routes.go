package http

import (
	"net/http"

	"github.com/hassan-alidoost/oms-project/internal/handlers/http/order"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handlers struct {
	OrderHandle *order.OrderHandler
}

func NewRouter(h Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	//swagger
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	//order
	mux.HandleFunc("POST /orders", h.OrderHandle.CreateOrder)
	mux.HandleFunc("GET /orders/{id}", h.OrderHandle.GetOrder)

	return mux
}