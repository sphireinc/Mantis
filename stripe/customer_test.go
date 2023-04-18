package stripe

import (
	"errors"
	"github.com/stripe/stripe-go/v72/paymentmethod"
	"github.com/stripe/stripe-go/v72/sub"
	"testing"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
	"github.com/stripe/stripe-go/v72/customer"
)

// Define a mock Stripe client for testing purposes
type mockStripeClient struct {
	Customers      *customer.Client
	PaymentMethods *paymentmethod.Client
	Subscriptions  *sub.Client
}

func (m *mockStripeClient) NewCustomer(params *stripe.CustomerParams) (*stripe.Customer, error) {
	if params.Email == nil || params.Name == nil {
		return nil, errors.New("missing email or name in customer params")
	}
	return &stripe.Customer{
		ID:    "cus_test123",
		Email: *params.Email,
		Name:  *params.Name,
	}, nil
}

func (m *mockStripeClient) New(params *stripe.SubscriptionParams) (*stripe.Subscription, error) {
	if params.Customer == nil || params.Items == nil {
		return nil, errors.New("missing customer ID or items in SubscriptionParams")
	}
	var subItems []*stripe.SubscriptionItem
	return &stripe.Subscription{
		ID:       "sub_test123",
		Customer: &stripe.Customer{ID: *params.Customer},
		Status:   "active",
		Items:    &stripe.SubscriptionItemList{Data: subItems},
	}, nil
}

func (m *mockStripeClient) Attach(paymentMethodID string, params *stripe.PaymentMethodAttachParams) (*stripe.PaymentMethod, error) {
	if params.Customer == nil {
		return nil, errors.New("missing customer ID in PaymentMethodAttachParams")
	}
	return &stripe.PaymentMethod{
		ID:       paymentMethodID,
		Customer: &stripe.Customer{ID: *params.Customer},
		Type:     "card",
		BillingDetails: &stripe.BillingDetails{
			Name: "John Doe",
		},
	}, nil
}

func (m *mockStripeClient) Detach(paymentMethodID string) (*stripe.PaymentMethod, error) {
	return &stripe.PaymentMethod{
		ID: paymentMethodID,
	}, nil
}

func TestCreateCustomer(t *testing.T) {
	// Create a mock Stripe client
	sc := &client.API{
		Customers: mockStripeClient{}.Customers,
	}

	// Test case 1: valid parameters
	customer_, err := createCustomer(sc, "john.doe@example.com", "John Doe")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if customer_.ID != "cus_test123" {
		t.Errorf("Expected customer_ ID to be cus_test123, but got %v", customer_.ID)
	}
	if customer_.Email != "john.doe@example.com" {
		t.Errorf("Expected customer_ email to be john.doe@example.com, but got %v", customer_.Email)
	}
	if customer_.Name != "John Doe" {
		t.Errorf("Expected customer_ name to be John Doe, but got %v", customer_.Name)
	}

	// Test case 2: missing email
	_, err = createCustomer(sc, "", "John Doe")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if err.Error() != "missing email or name in customer_ params" {
		t.Errorf("Expected error message 'missing email or name in customer_ params', but got '%v'", err.Error())
	}

	// Test case 3: missing name
	_, err = createCustomer(sc, "john.doe@example.com", "")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if err.Error() != "missing email or name in customer_ params" {
		t.Errorf("Expected error message 'missing email or name in customer_ params', but got '%v'", err.Error())
	}
}

func TestAttachPaymentMethodToCustomer(t *testing.T) {
	// Create a mock Stripe client
	sc := &client.API{
		PaymentMethods: mockStripeClient{}.PaymentMethods,
	}

	// Test case 1: valid parameters
	paymentMethodID := "pm_test123"
	customerID := "cus_test123"
	paymentMethod, err := attachPaymentMethodToCustomer(sc, customerID, paymentMethodID)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if paymentMethod.ID != paymentMethodID {
		t.Errorf("Expected payment method ID to be %v, but got %v", paymentMethodID, paymentMethod.ID)
	}
	if paymentMethod.Customer.ID != customerID {
		t.Errorf("Expected payment method customer ID to be %v, but got %v", customerID, paymentMethod.Customer)
	}
	if paymentMethod.Type != "card" {
		t.Errorf("Expected payment method type to be 'card', but got %v", paymentMethod.Type)
	}
	if paymentMethod.BillingDetails.Name != "John Doe" {
		t.Errorf("Expected payment method billing name to be 'John Doe', but got %v", paymentMethod.BillingDetails.Name)
	}

	// Test case 2: missing customer ID
	paymentMethodID = "pm_test456"
	customerID = ""
	_, err = attachPaymentMethodToCustomer(sc, customerID, paymentMethodID)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if err.Error() != "missing customer ID in PaymentMethodAttachParams" {
		t.Errorf("Expected error message 'missing customer ID in PaymentMethodAttachParams', but got '%v'", err.Error())
	}
}

func TestDetachPaymentMethodFromCustomer(t *testing.T) {
	// Create a mock Stripe client
	sc := &client.API{
		PaymentMethods: mockStripeClient{}.PaymentMethods,
	}

	// Test case 1: valid parameters
	// Test case 1: valid parameters
	paymentMethodID := "pm_test123"
	paymentMethod, err := detachPaymentMethodFromCustomer(sc, paymentMethodID)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if paymentMethod.ID != paymentMethodID {
		t.Errorf("Expected payment method ID to be %v, but got %v", paymentMethodID, paymentMethod.ID)
	}

	// Test case 2: invalid payment method ID
	paymentMethodID = ""
	_, err = detachPaymentMethodFromCustomer(sc, paymentMethodID)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if err.Error() != "Invalid payment method ID" {
		t.Errorf("Expected error message 'Invalid payment method ID', but got '%v'", err.Error())
	}
}

func TestCreateSubscription(t *testing.T) {
	// Create a mock Stripe client
	sc := &client.API{
		Subscriptions: mockStripeClient{}.Subscriptions,
	}

	// Test case 1: valid parameters
	customerID := "cus_test123"
	priceID := "price_test123"
	subscription, err := createSubscription(sc, customerID, priceID)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if subscription.ID != "sub_test123" {
		t.Errorf("Expected subscription ID to be sub_test123, but got %v", subscription.ID)
	}
	if subscription.Customer.ID != customerID {
		t.Errorf("Expected subscription customer ID to be %v, but got %v", customerID, subscription.Customer)
	}
	if subscription.Status != "active" {
		t.Errorf("Expected subscription status to be 'active', but got %v", subscription.Status)
	}
	if len(subscription.Items.Data) != 1 {
		t.Errorf("Expected subscription to have 1 item, but got %v", len(subscription.Items.Data))
	}
	if subscription.Items.Data[0].Price.ID != priceID {
		t.Errorf("Expected subscription item price ID to be %v, but got %v", priceID, subscription.Items.Data[0].Price.ID)
	}

	// Test case 2: missing customer ID
	customerID = ""
	priceID = "price_test456"
	_, err = createSubscription(sc, customerID, priceID)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if err.Error() != "missing customer ID or items in SubscriptionParams" {
		t.Errorf("Expected error message 'missing customer ID or items in SubscriptionParams', but got '%v'", err.Error())
	}

	// Test case 3: missing items
	customerID = "cus_test789"
	priceID = ""
	_, err = createSubscription(sc, customerID, priceID)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if err.Error() != "missing customer ID or items in SubscriptionParams" {
		t.Errorf("Expected error message 'missing customer ID or items in SubscriptionParams', but got '%v'", err.Error())
	}
}
