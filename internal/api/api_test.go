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
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok := isValidPrice(test.description.key, test.description.value)
			if ok != test.want {
				t.Errorf("%s - got %t, want %t", test.name, ok, test.want)
			}
		})
	}
}
