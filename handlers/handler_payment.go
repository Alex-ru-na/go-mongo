package handlers

import (
	"encoding/json"
	"go-mongodb-api/models"
	"go-mongodb-api/services"
	"log"
	"net/http"
)

type PaymentHandlers struct {
	service *services.PaymentService
}

func NewPaymentHandlers(service *services.PaymentService) *PaymentHandlers {
	return &PaymentHandlers{service: service}
}

func (h *PaymentHandlers) Payment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var payment models.Payment
		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			log.Printf("Error decoding request body: %v", err)
			http.Error(w, "Invalid request body", http.StatusUnprocessableEntity)
			return
		}

		result, err := h.service.CreatePayment(payment)
		if err != nil {
			log.Printf("Error creating payment: %v", err)
			http.Error(w, "Failed to create payment", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
