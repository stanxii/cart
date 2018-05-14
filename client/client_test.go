package client_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	cartclient "github.com/marcusolsson/coverage-cravings/client"
)

func TestGetOrder(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile(filepath.Join("testdata", "order_response.json"))
		if err != nil {
			t.Fatal(err)
		}
		w.Write(b)
	}))

	c := cartclient.New(cartclient.SetBaseURL(srv.URL))

	got, err := c.Order("ABC123")
	if err != nil {
		t.Fatal(err)
	}

	want := cartclient.Order{
		Status: "CART",
		Items: []cartclient.LineItem{
			{
				Quantity: 2,
				Product: cartclient.Product{
					Title: "iPhone 4",
					Price: 499,
				},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\ngot:\n%# v\n\nwant:\n%# v", pretty.Formatter(got), pretty.Formatter(want))
	}
}

func TestGetUnknownOrder(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	c := cartclient.New(cartclient.SetBaseURL(srv.URL))

	_, err := c.Order("ABC123")
	if err != cartclient.ErrOrderNotFound {
		t.Fatalf("unexpected error: %q", err)
	}
}
