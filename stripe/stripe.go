package stripe

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

func initStripeClient(apiKey string) *client.API {
	stripe.Key = apiKey
	sc := &client.API{}
	sc.Init(stripe.Key, nil)
	return sc
}
