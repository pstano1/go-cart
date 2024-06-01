package stripeProvider

import (
	"github.com/pstano1/go-cart/internal/pkg"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	"go.uber.org/zap"
)

type IStripeProvider interface {
	CreatePayment(basket []pkg.ProductSummary) (string, error)
}

type StripeProvider struct {
	log       *zap.Logger
	stripeKey string
}

func NewProvider(key string, logger *zap.Logger) IStripeProvider {
	stripe.Key = key

	return &StripeProvider{
		log:       logger,
		stripeKey: key,
	}
}

func (s *StripeProvider) CreatePayment(basket []pkg.ProductSummary) (string, error) {
	items := make([]*stripe.CheckoutSessionLineItemParams, 0)
	for _, product := range basket {
		item := &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency:   stripe.String(product.Currency),
				UnitAmount: stripe.Int64(int64(product.Price * 100)),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: &product.Name,
				},
			},
			Quantity: stripe.Int64(int64(product.Quantity)),
		}
		items = append(items, item)
	}

	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("https://example.com/success"),
		LineItems:  items,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
	}
	result, err := session.New(params)
	return result.URL, err
}
