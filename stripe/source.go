package stripe

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/source"
)

func getSource(sourceID string) (*stripe.Source, error) {
	params := &stripe.SourceObjectParams{}
	return source.Get(sourceID, params)
}
