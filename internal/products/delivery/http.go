package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

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
		writeError(w, http.StatusInternalServerError, "failed to list products")
		return
	}
	writeJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	productID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	product, err := h.service.GetProductByID(productID)
	if err != nil {
		writeError(w, http.StatusNotFound, "product not found")
		return
	}
	writeJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p entity.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	product, err := h.service.CreateProduct(p)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create product")
		return
	}
	writeJSON(w, http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	productID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	var p entity.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	updated, err := h.service.UpdateProduct(productID, p)
	if err != nil {
		writeError(w, http.StatusNotFound, "product not found")
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	productID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	if err := h.service.DeleteProduct(productID); err != nil {
		writeError(w, http.StatusNotFound, "product not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    status,
			"message": message,
		},
	})
}
