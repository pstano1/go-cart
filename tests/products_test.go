package tests

import (
	"testing"

	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/pkg"
)

func TestGetProducts(t *testing.T) {
	tests := []struct {
		name   string
		in     *api.InstanceAPI
		filter *pkg.ProductFilter
		want   []pkg.Product
	}{
		{
			name: "non existent customerId",
			in:   API,
			filter: &pkg.ProductFilter{
				CustomerId: "test",
			},
			want: make([]pkg.Product, 0),
		},
		{
			name: "query all products for one customer",
			in:   API,
			filter: &pkg.ProductFilter{
				CustomerId: customerId,
			},
			want: make([]pkg.Product, 1),
		},
		{
			name: "query product by categories (product not in category)",
			in:   API,
			filter: &pkg.ProductFilter{
				Categories: []string{"books"},
				CustomerId: customerId,
			},
			want: make([]pkg.Product, 0),
		},
		{
			name: "query products by categories",
			in:   API,
			filter: &pkg.ProductFilter{
				Categories: []string{"electronics"},
				CustomerId: customerId,
			},
			want: make([]pkg.Product, 1),
		},
		{
			name: "query products by id (product doesn't exist)",
			in:   API,
			filter: &pkg.ProductFilter{
				Id:         "test",
				CustomerId: customerId,
			},
			want: make([]pkg.Product, 0),
		},
		{
			name: "query user by id (user exists)",
			in:   API,
			filter: &pkg.ProductFilter{
				Id:         "0883981f-dd7e-436f-a753-be8172324e28",
				CustomerId: customerId,
			},
			want: make([]pkg.Product, 1),
		},
	}

	for _, test := range tests {
		products, _ := test.in.GetProducts(test.filter)
		if len(products) != len(test.want) {
			t.Errorf("%s - got %d length, want %d length", test.name, len(products), len(test.want))
		}
	}
}

func TestGetProductCategories(t *testing.T) {
	tests := []struct {
		name   string
		in     *api.InstanceAPI
		filter *pkg.CategoryFilter
		want   []pkg.ProductCategory
	}{
		{
			name: "non existent customerId",
			in:   API,
			filter: &pkg.CategoryFilter{
				CustomerId: "test",
			},
			want: make([]pkg.ProductCategory, 0),
		},
		{
			name: "query all categories for one customer",
			in:   API,
			filter: &pkg.CategoryFilter{
				CustomerId: customerId,
			},
			want: make([]pkg.ProductCategory, 5),
		},
		{
			name: "query categories by id (category doesn't exist)",
			in:   API,
			filter: &pkg.CategoryFilter{
				Id:         "test",
				CustomerId: customerId,
			},
			want: make([]pkg.ProductCategory, 0),
		},
		{
			name: "query categories by id (catagory exists)",
			in:   API,
			filter: &pkg.CategoryFilter{
				Id:         "c1a2b5a1-6851-438b-a055-2ae0d1116b50",
				CustomerId: customerId,
			},
			want: make([]pkg.ProductCategory, 1),
		},
	}

	for _, test := range tests {
		categories, _ := test.in.GetCategories(test.filter)
		if len(categories) != len(test.want) {
			t.Errorf("%s - got %d length, want %d length", test.name, len(categories), len(test.want))
		}
	}
}

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		name    string
		in      *api.InstanceAPI
		payload *pkg.ProductCreate
		want    error
	}{
		{
			name: "proper product setup",
			in:   API,
			payload: &pkg.ProductCreate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				Names: pkg.JSONB{
					"EN": "Example Product",
					"PL": "Przykładowy Produkt",
				},
				Descriptions: pkg.JSONB{
					"EN": "This is an example product description.",
					"PL": "To jest przykładowy opis produktu.",
				},
				Categories: []string{"electronics"},
				Prices: pkg.JSONB{
					"USD": 50.25,
					"EUR": 50,
				},
			},
			want: nil,
		},
	}

	for _, test := range tests {
		_, err := test.in.CreateProduct(test.payload)
		if err != test.want {
			t.Errorf("%s - got %d, want %d", test.name, err, test.want)
		}
	}
}

func TestCreateProductCategory(t *testing.T) {
	tests := []struct {
		name    string
		in      *api.InstanceAPI
		payload *pkg.CategoryCreate
		want    error
	}{
		{
			name: "proper category setup",
			in:   API,
			payload: &pkg.CategoryCreate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				Name: "test-category",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		_, err := test.in.CreateCategory(test.payload)
		if err != test.want {
			t.Errorf("%s - got %d, want %d", test.name, err, test.want)
		}
	}
}

func TestUpdateProduct(t *testing.T) {
	tests := []struct {
		name    string
		in      *api.InstanceAPI
		payload *pkg.ProductUpdate
		want    error
	}{
		{
			name: "proper product setup",
			in:   API,
			payload: &pkg.ProductUpdate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				Names: pkg.JSONB{
					"en": "Example Product",
					"pl": "Przykładowy Produkt",
				},
				Descriptions: pkg.JSONB{
					"en": "This is an example product description.",
					"pl": "To jest przykładowy opis produktu.",
				},
				Categories: []string{"books"},
				Prices: pkg.JSONB{
					"USD": 50.25,
					"EUR": 50,
				},
				Id: "0883981f-dd7e-436f-a753-be8172324e28",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		err := test.in.UpdateProduct(test.payload)
		if err != test.want {
			t.Errorf("%s - got %d, want %d", test.name, err, test.want)
		}
	}
}

func TestUpdateProductCategory(t *testing.T) {
	tests := []struct {
		name    string
		in      *api.InstanceAPI
		payload *pkg.CategoryUpdate
		want    error
	}{
		{
			name: "proper category update request",
			in:   API,
			payload: &pkg.CategoryUpdate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				Name: "test-category-1",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		err := test.in.UpdateCategory(test.payload)
		if err != test.want {
			t.Errorf("%s - got %d, want %d", test.name, err, test.want)
		}
	}
}
