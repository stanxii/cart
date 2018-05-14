package cart

import "errors"

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

func NewOrder() *Order {
	return &Order{
		Items: make([]LineItem, 0),
	}
}

func (o *Order) AddLineItem(item LineItem) {
	o.Items = append(o.Items, item)
}

func (o *Order) Empty() {
	o.Items = make([]LineItem, 0)
}

func (o *Order) Total() float64 {
	var total float64
	for _, i := range o.Items {
		total += float64(i.Quantity) * i.Product.Price
	}
	return total
}

func (o *Order) Size() int {
	return len(o.Items)
}

var ErrOrderNotFound = errors.New("order not found")

type OrderRepository interface {
	Save(o *Order) error
	FindByID(id string) (*Order, error)
}
