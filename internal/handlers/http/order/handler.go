package order

import (
	"net/http"
	"strings"

	"github.com/hassan-alidoost/oms-project/internal/services"
)

type OrderHandler struct {
	svc *services.OrderService
}

func NewOrderHandler(svc *services.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

// createOrder handles POST /orders.
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var dto CreateDTO
	if err := decode(r, &dto); err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	result, err := h.svc.Create(r.Context(), dto)
	if err != nil {
		respondError(w, err)
		return
	}
	respond(w, http.StatusCreated, result)
}

// listOrders handles GET /orders.
func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	results, err := h.svc.GetAll(r.Context())
	if err != nil {
		respondError(w, err)
		return
	}
	respond(w, http.StatusOK, results)
}

// getOrder handles GET /orders/{id}.
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/orders/")
	if rest == "" {
		respond(w, http.StatusBadRequest, map[string]string{"error": "missing order id"})
		return
	}

	parts := strings.SplitN(rest, "/", 2)
	id := parts[0]

	result, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		respondError(w, err)
		return
	}
	respond(w, http.StatusOK, result)
}

// updateStatus handles PUT /orders/{id}/status.
func (h *OrderHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/orders/")
	if rest == "" {
		respond(w, http.StatusBadRequest, map[string]string{"error": "missing order id"})
		return
	}

	parts := strings.SplitN(rest, "/", 2)
	id := parts[0]

	var dto UpdateStatusDTO
	if err := decode(r, &dto); err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	result, err := h.svc.UpdateStatus(r.Context(), id, dto)
	if err != nil {
		respondError(w, err)
		return
	}
	respond(w, http.StatusOK, result)
}
