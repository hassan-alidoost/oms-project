package order

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hassan-alidoost/oms-project/internal/domain"
	"github.com/hassan-alidoost/oms-project/internal/ports"
)

type OrderHandler struct {
	service ports.OrderService
}

func NewOrderHandler(service ports.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// CreateOrder handles the creation of a new order
// @Summary Create a new order
// @Description Creates an order with a random customer ID and specific price
// @Tags orders
// @Accept json
// @Produce json
// @Param order body CreateOrderRequest true "Order Details"
// @Success 201 {object} domain.Order
// @Failure 400 {string} string "Invalid request"
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderDto CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&orderDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	items := make([]domain.OrderItem, 0, len(orderDto.Items))
	for _, item := range orderDto.Items {
		items = append(items, domain.NewOrderItem(item.ProductID, item.Quantity, item.Price))
	}


	err := h.service.Create(r.Context(), items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// GetOrder retrieves an order by ID
// @Summary Get an order
// @Description Returns a single order by its uint64 ID
// @Tags orders
// @Produce json
// @Param id path uint64 true "Order ID"
// @Success 200 {object} domain.Order
// @Failure 404 {string} string "Order not found"
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	order, err := h.service.FindByID(r.Context(), domain.ID(id))
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
    orders, err := h.service.GetOrders(r.Context())

    if err != nil {
        http.Error(w, "failed to fetch orders", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(orders); err != nil {
        http.Error(w, "failed to encode response", http.StatusInternalServerError)
    }
}
