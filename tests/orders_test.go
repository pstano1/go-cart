package tests

import (
	"testing"

	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/pkg"
)

func TestGetOrders(t *testing.T) {
	tests := []struct {
		name   string
		in     *api.InstanceAPI
		filter *pkg.OrderFilter
		want   []pkg.Order
	}{
		{
			name: "non existent customerId",
			in:   API,
			filter: &pkg.OrderFilter{
				CustomerId: "test",
			},
			want: make([]pkg.Order, 0),
		},
		{
			name: "query all orders for one customer",
			in:   API,
			filter: &pkg.OrderFilter{
				CustomerId: customerId,
			},
			want: make([]pkg.Order, 1),
		},
		{
			name: "query order by id (order doesn't exist)",
			in:   API,
			filter: &pkg.OrderFilter{
				Id:         "test",
				CustomerId: customerId,
			},
			want: make([]pkg.Order, 0),
		},
		{
			name: "query order by id (order exists)",
			in:   API,
			filter: &pkg.OrderFilter{
				Id:         "4709fbb1-6462-4518-859c-f804b43b0d2e",
				CustomerId: customerId,
			},
			want: make([]pkg.Order, 1),
		},
	}

	for _, test := range tests {
		coupons, _ := test.in.GetOrders(test.filter)
		if len(coupons) != len(test.want) {
			t.Errorf("%s - got %d length, want %d length", test.name, len(coupons), len(test.want))
		}
	}
}

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		name    string
		in      *api.InstanceAPI
		payload *pkg.OrderCreate
		want    error
	}{
		{
			name: "invalid basket value",
			in:   API,
			payload: &pkg.OrderCreate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				TotalCost:  50.00,
				Currency:   "EUR",
				Country:    "PL",
				City:       "Warszawa",
				PostalCode: "00-902",
				Address:    "ul. Wiejska 4",
				Basket: pkg.JSONB{
					"price":    50.00,
					"currency": "EUR",
					"name":     "Example Product",
					"quantity": 1.0,
				},
				TaxId: "",
			},
			want: pkg.ErrInvalidBasketValue,
		},
		{
			name: "proper order create request",
			in:   API,
			payload: &pkg.OrderCreate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				TotalCost:  50.00,
				Currency:   "EUR",
				Country:    "PL",
				City:       "Warszawa",
				PostalCode: "00-902",
				Address:    "ul. Wiejska 4",
				Basket: pkg.JSONB{
					"0883981f-dd7e-436f-a753-be8172324e28": map[string]interface{}{
						"price":    50.00,
						"currency": "EUR",
						"name":     "Example Product",
						"quantity": 1.0,
					},
				},
				TaxId: "",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		_, err := test.in.CreateOrder(test.payload)
		if err != test.want {
			t.Errorf("%s - got %d, want %d", test.name, err, test.want)
		}
	}
}

func TestUpdateOrder(t *testing.T) {
	tests := []struct {
		name    string
		in      *api.InstanceAPI
		payload *pkg.OrderUpdate
		want    error
	}{
		{
			name: "non existent customerId in request",
			in:   API,
			payload: &pkg.OrderUpdate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				Id:         "test",
				TotalCost:  50.00,
				Currency:   "EUR",
				Country:    "PL",
				City:       "Warszawa",
				PostalCode: "00-902",
				Address:    "ul. Wiejska 4",
				Status:     "paid",
				Basket: pkg.JSONB{
					"price":    50.00,
					"currency": "EUR",
					"name":     "Example Product",
					"quantity": 1,
				},
				TaxId: "",
			},
			want: pkg.ErrOrderNotFound,
		},
		{
			name: "proper order update request",
			in:   API,
			payload: &pkg.OrderUpdate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				TotalCost:  50.00,
				Currency:   "EUR",
				Country:    "PL",
				City:       "Warszawa",
				PostalCode: "00-902",
				Address:    "ul. Wiejska 4",
				Status:     "paid",
				Basket: pkg.JSONB{
					"price":    50.00,
					"currency": "EUR",
					"name":     "Example Product",
					"quantity": 1,
				},
				TaxId: "",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		err := test.in.UpdateOrder(test.payload)
		if err != test.want {
			t.Errorf("%s - got %d, want %d", test.name, err, test.want)
		}
	}
}
