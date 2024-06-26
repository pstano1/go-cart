package tests

import (
	"testing"

	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/pkg"
)

func TestGetCoupons(t *testing.T) {
	tests := []struct {
		name   string
		in     *api.InstanceAPI
		filter *pkg.CouponFilter
		want   []pkg.Coupon
	}{
		{
			name: "non existent customerId",
			in:   API,
			filter: &pkg.CouponFilter{
				CustomerId: "test",
			},
			want: make([]pkg.Coupon, 0),
		},
		{
			name: "query all coupons for one customer",
			in:   API,
			filter: &pkg.CouponFilter{
				CustomerId: customerId,
			},
			want: make([]pkg.Coupon, 1),
		},
		{
			name: "query coupon by code (coupon doesn't exist)",
			in:   API,
			filter: &pkg.CouponFilter{
				Code:       "MAY25",
				CustomerId: customerId,
			},
			want: make([]pkg.Coupon, 0),
		},
		{
			name: "query coupon by code (coupon exists)",
			in:   API,
			filter: &pkg.CouponFilter{
				Code:       "MAY20",
				CustomerId: customerId,
			},
			want: make([]pkg.Coupon, 1),
		},
		{
			name: "query coupon by id (coupon doesn't exist)",
			in:   API,
			filter: &pkg.CouponFilter{
				Id:         "test",
				CustomerId: customerId,
			},
			want: make([]pkg.Coupon, 0),
		},
		{
			name: "query coupon by id (coupon exists)",
			in:   API,
			filter: &pkg.CouponFilter{
				Id:         "505892df-ac66-42f6-9fab-74fd03dbc5f3",
				CustomerId: customerId,
			},
			want: make([]pkg.Coupon, 1),
		},
	}

	for _, test := range tests {
		coupons, _ := test.in.GetCoupons(test.filter)
		if len(coupons) != len(test.want) {
			t.Errorf("%s - got %d, want %d", test.name, len(coupons), len(test.want))
		}
	}
}

func TestCreateCoupons(t *testing.T) {
	tests := []struct {
		name    string
		in      *api.InstanceAPI
		payload *pkg.CouponCreate
		want    error
	}{
		{
			name: "proper coupon setup",
			in:   API,
			payload: &pkg.CouponCreate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				PromoCode: "MAY25",
				Unit:      "percentage",
				Amount:    25,
			},
			want: nil,
		},
	}

	for _, test := range tests {
		_, err := test.in.CreateCoupon(test.payload)
		if err != test.want {
			t.Errorf("%s - got %d length, want %d length", test.name, err, test.want)
		}
	}
}

func TestUpdateCoupons(t *testing.T) {
	tests := []struct {
		name    string
		in      *api.InstanceAPI
		payload *pkg.CouponUpdate
		want    error
	}{
		{
			name: "non existent id in request",
			in:   API,
			payload: &pkg.CouponUpdate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				PromoCode: "JUNE20",
				Unit:      "percentage",
				Amount:    20,
			},
			want: nil,
		},
		{
			name: "proper coupon update request",
			in:   API,
			payload: &pkg.CouponUpdate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				PromoCode: "JUNE20",
				Unit:      "percentage",
				Amount:    20,
				Id:        "505892df-ac66-42f6-9fab-74fd03dbc5f3",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		err := test.in.UpdateCoupon(test.payload)
		if err != test.want {
			t.Errorf("%s - got %d, want %d", test.name, err, test.want)
		}
	}
}
