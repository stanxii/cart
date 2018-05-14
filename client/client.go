package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var ErrOrderNotFound = errors.New("order not found")

type Product struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type LineItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

type Order struct {
	Status string     `json:"status"`
	Items  []LineItem `json:"items"`
}

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

func SetBaseURL(rawurl string) func(*Client) {
	return func(c *Client) {
		u, _ := url.Parse(rawurl)
		c.baseURL = u
	}
}

func New(opts ...func(*Client)) *Client {
	c := &Client{
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) Order(id string) (Order, error) {
	req, err := http.NewRequest("GET", c.baseURL.String()+"/orders", nil)
	if err != nil {
		return Order{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Order{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		if resp.StatusCode == 404 {
			return Order{}, ErrOrderNotFound
		}
		return Order{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var order Order
	if err := json.NewDecoder(resp.Body).Decode(&order); err != nil {
		return Order{}, err
	}

	return order, nil
}
