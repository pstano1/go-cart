package api

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

type Price struct {
	currency string
	value    interface{}
}

func TestPriceValidator(t *testing.T) {
	var tests = []struct {
		name  string
		price Price
		want  bool
	}{
		{
			name: "too short currency symbol",
			price: Price{
				currency: "EU",
				value:    20,
			},
			want: false,
		},
		{
			name: "too long currency symbol",
			price: Price{
				currency: "EURO",
				value:    20,
			},
			want: false,
		},
		{
			name: "price value as wrong type (string)",
			price: Price{
				currency: "EUR",
				value:    "20",
			},
			want: false,
		},
		{
			name: "price value as wrong type (interface{})",
			price: Price{
				currency: "EUR",
				value: struct {
					price int
				}{price: 20},
			},
			want: false,
		},
		{
			name: "correct values with value as int",
			price: Price{
				currency: "EUR",
				value:    20,
			},
			want: true,
		},
		{
			name: "correct values with value as float",
			price: Price{
				currency: "EUR",
				value:    20.05,
			},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok := isValidPrice(test.price.currency, test.price.value)
			if ok != test.want {
				t.Errorf("%s - got %t, want %t", test.name, ok, test.want)
			}
		})
	}
}

type Description struct {
	key   string
	value interface{}
}

func TestDescriptionValidator(t *testing.T) {
	var tests = []struct {
		name        string
		description Description
		want        bool
	}{
		{
			name: "too short key",
			description: Description{
				key:   "E",
				value: "Lorem ipsum...",
			},
			want: false,
		},
		{
			name: "too long key",
			description: Description{
				key:   "ENG",
				value: "Lorem ipsum...",
			},
			want: false,
		},
		{
			name: "value as wrong type (interface)",
			description: Description{
				key: "EN",
				value: struct {
					content string
				}{content: "Lorem ipsum..."},
			},
			want: false,
		},
		{
			name: "correct values",
			description: Description{
				key:   "EN",
				value: "Lorem ipsum...",
			},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok := isValidNameOrDescription(test.description.key, test.description.value)
			if ok != test.want {
				t.Errorf("%s - got %t, want %t", test.name, ok, test.want)
			}
		})
	}
}

type ProductDescription struct {
	key   string
	value interface{}
}

func TestProductDescriptionValidator(t *testing.T) {
	var tests = []struct {
		name        string
		basketEntry ProductDescription
		want        bool
	}{
		{
			name: "price as proper type (int)",
			basketEntry: ProductDescription{
				key:   "price",
				value: 100,
			},
			want: true,
		},
		{
			name: "price as proper type (float)",
			basketEntry: ProductDescription{
				key:   "price",
				value: 100.00,
			},
			want: true,
		},
		{
			name: "price as wrong type (string)",
			basketEntry: ProductDescription{
				key:   "price",
				value: "100",
			},
			want: false,
		},
		{
			name: "quantity as proper type (int)",
			basketEntry: ProductDescription{
				key:   "quantity",
				value: 1,
			},
			want: true,
		},
		{
			name: "quantity as wrong type (string)",
			basketEntry: ProductDescription{
				key:   "quantity",
				value: "100",
			},
			want: false,
		},
		{
			name: "currency as wrong type (int)",
			basketEntry: ProductDescription{
				key:   "currency",
				value: 100,
			},
			want: false,
		},
		{
			name: "currency with too short value",
			basketEntry: ProductDescription{
				key:   "currency",
				value: "PL",
			},
			want: false,
		},
		{
			name: "currency with too long value",
			basketEntry: ProductDescription{
				key:   "currency",
				value: "polski złoty",
			},
			want: false,
		},
		{
			name: "currency with digits as value",
			basketEntry: ProductDescription{
				key:   "currency",
				value: "123",
			},
			want: false,
		},
		{
			name: "proper currency value",
			basketEntry: ProductDescription{
				key:   "currency",
				value: "PLN",
			},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok := isValidBasketProductDescription(test.basketEntry.key, test.basketEntry.value)
			if ok != test.want {
				t.Errorf("%s - got %t, want %t", test.name, ok, test.want)
			}
		})
	}
}

type BasketEntry struct {
	key   string
	value interface{}
}

func TestBasketEntryValidator(t *testing.T) {
	var tests = []struct {
		name        string
		basketEntry BasketEntry
		want        bool
	}{
		{
			name: "value as wrong type (string)",
			basketEntry: BasketEntry{
				key:   "test-id",
				value: "string",
			},
			want: false,
		},
		{
			name: "value as wrong type (int)",
			basketEntry: BasketEntry{
				key:   "test-id",
				value: 100,
			},
			want: false,
		},
		{
			name: "incomplete value",
			basketEntry: BasketEntry{
				key: "test-id",
				value: map[string]interface{}{
					"price":    100,
					"currency": "PLN",
					"name":     "Product #0",
				},
			},
			want: false,
		},
		{
			name: "complete value with wrong data type",
			basketEntry: BasketEntry{
				key: "test-id",
				value: map[string]interface{}{
					"price":    100,
					"currency": "PLN",
					"quantity": 5.5,
					"name":     "Product #0",
				},
			},
			want: false,
		},
		{
			name: "proper payload",
			basketEntry: BasketEntry{
				key: "test-id",
				value: map[string]interface{}{
					"price":    100,
					"currency": "PLN",
					"quantity": 1,
					"name":     "Product #0",
				},
			},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok := isValidBasketEntry(test.basketEntry.key, test.basketEntry.value)
			if ok != test.want {
				t.Errorf("%s - got %t, want %t", test.name, ok, test.want)
			}
		})
	}
}
