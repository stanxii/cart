package cart

import "testing"

func TestOrder_AddLineItem(t *testing.T) {
	o := NewOrder()

	o.AddLineItem(LineItem{
		Product:  Product{"Phone", 10.0},
		Quantity: 2,
	})

	eq(t, 1, o.Size())
	eq(t, 20.0, o.Total())
}

func eq(t *testing.T, want, got interface{}) {
	t.Helper()

	if want != got {
		t.Errorf("got = %v; want = %v", got, want)
	}
}
