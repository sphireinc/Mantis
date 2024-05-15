package stripe

import (
	"testing"

	"github.com/stripe/stripe-go/v72"
)

func TestInitStripeClient(t *testing.T) {
	// Test case 1: valid API key
	apiKey := "sk_test_123"
	sc := initStripeClient(apiKey)
	if stripe.Key != apiKey {
		t.Errorf("Expected stripe.Key to be %v, but got %v", apiKey, stripe.Key)
	}
	if sc == nil {
		t.Error("Expected sc to be non-nil, but got nil")
	}

	// Test case 2: empty API key
	apiKey = ""
	sc = initStripeClient(apiKey)
	if stripe.Key != "" {
		t.Errorf("Expected stripe.Key to be empty, but got %v", stripe.Key)
	}
	if sc == nil {
		t.Error("Expected sc to be non-nil, but got nil")
	}
}
