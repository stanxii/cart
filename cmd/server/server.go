package main

import (
	"encoding/json"
	"net/http"

	cart "github.com/marcusolsson/coverage-cravings"
)

type server struct {
	orders cart.OrderRepository
}

func (s *server) register(mux *http.ServeMux) {
	mux.HandleFunc("/orders", s.listOrders)
}

func (s *server) listOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("order_id")

		os, err := s.orders.FindByID(id)
		if err != nil {
			if err == cart.ErrOrderNotFound {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(os)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	case "POST":
		order := new(cart.Order)
		if err := json.NewDecoder(r.Body).Decode(order); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.orders.Save(order); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
