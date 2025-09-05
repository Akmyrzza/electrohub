package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/akmyrzza/electrohub/internal/products/entity"
	"github.com/akmyrzza/electrohub/internal/products/usecase"
)

type ProductHandler struct {
	service *usecase.ProductService
}

func NewProductHandler(service *usecase.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := h.service.ListProducts()
	if err != nil {
		http.Error(w, "failed to list products", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p entity.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateProduct(p); err != nil {
		http.Error(w, "failed to create product", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, p)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
