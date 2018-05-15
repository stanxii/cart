package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/marcusolsson/cart"
	"github.com/marcusolsson/cart/mock"
)

func TestCreateOrder(t *testing.T) {
	t.Run("Should return 202 when order was created", func(t *testing.T) {
		rec := httptest.NewRecorder()

		body := map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{
					"product": map[string]interface{}{
						"title": "Apple iPhone",
						"price": 99.0,
					},
					"quantity": 2,
				},
			},
		}
		b, _ := json.Marshal(body)

		req := httptest.NewRequest("POST", "/orders", bytes.NewReader(b))
		mux := http.NewServeMux()

		var got *cart.Order
		srv := &server{
			orders: &mock.OrderRepository{
				SaveFunc: func(o *cart.Order) error {
					got = o
					return nil
				},
			},
		}
		srv.register(mux)

		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Fatalf("status = %s; want = %s", http.StatusText(rec.Code), http.StatusText(http.StatusCreated))
		}

		if got.Total() != 198.0 {
			t.Fatalf("got.Total() = %f; want = %f", got.Total(), 198.0)
		}
	})
}

func TestListOrders(t *testing.T) {
	t.Run("Should return 200 with order when found", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/orders?order_id=123", nil)
		mux := http.NewServeMux()

		srv := &server{
			orders: &mock.OrderRepository{
				FindByIDFunc: func(id string) (*cart.Order, error) {
					return &cart.Order{
						Status: "CART",
						Items:  []cart.LineItem{},
					}, nil
				},
			},
		}
		srv.register(mux)

		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %s; want = %s", http.StatusText(rec.Code), http.StatusText(http.StatusOK))
		}

		want := map[string]interface{}{
			"status": "CART",
			"items":  []interface{}{},
		}

		equalsJSON(t, want, rec.Body.Bytes())
	})

	t.Run("Should return 404 when order not found", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/orders?order_id=123", nil)
		mux := http.NewServeMux()

		srv := &server{
			orders: &mock.OrderRepository{
				FindByIDFunc: func(id string) (*cart.Order, error) {
					return nil, cart.ErrOrderNotFound
				},
			},
		}
		srv.register(mux)

		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Fatalf("status = %s; want = %s", http.StatusText(rec.Code), http.StatusText(http.StatusNotFound))
		}
	})
}

func equalsJSON(t *testing.T, want map[string]interface{}, body []byte) {
	var got map[string]interface{}
	json.Unmarshal(body, &got)

	if !reflect.DeepEqual(want, got) {
		b, _ := json.MarshalIndent(want, "", "  ")

		var buf bytes.Buffer
		json.Indent(&buf, body, "", "  ")

		t.Fatalf("\ngot = %s\n\nwant = %s", buf.String(), string(b))
	}
}
