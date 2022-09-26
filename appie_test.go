package appie

import "testing"

func TestProductValues(t *testing.T) {
	a, _ := New()
	product, _ := a.GetProduct("130874")
	if product.Name == "" {
		t.Error("Product Name is EMPTY")
	}
	if product.Description == "" {
		t.Error("Product description is EMPTY")
	}
	if product.Price == "" {
		t.Error("Product price is EMPTY")
	}
	if product.Image == "" {
		t.Error("Product image is EMPTY")
	}
}