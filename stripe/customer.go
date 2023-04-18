package stripe

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

func createCustomer(sc *client.API, email, name string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
		Name:  stripe.String(name),
	}
	return sc.Customers.New(params)
}

func attachPaymentMethodToCustomer(sc *client.API, customerID, paymentMethodID string) (*stripe.PaymentMethod, error) {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(customerID),
	}
	return sc.PaymentMethods.Attach(paymentMethodID, params)
}

func detachPaymentMethodFromCustomer(sc *client.API, paymentMethodID string) (*stripe.PaymentMethod, error) {
	return sc.PaymentMethods.Detach(paymentMethodID, nil)
}

func createSubscription(sc *client.API, customerID, priceID string) (*stripe.Subscription, error) {
	items := []*stripe.SubscriptionItemsParams{
		{
			Price: stripe.String(priceID),
		},
	}
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerID),
		Items:    items,
	}
	return sc.Subscriptions.New(params)
}

func listSubscriptionsForCustomer(sc *client.API, customerID string) ([]*stripe.Subscription, error) {
	var subscriptions []*stripe.Subscription
	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(customerID),
	}
	i := sc.Subscriptions.List(params)

	for i.Next() {
		s := i.Subscription()
		subscriptions = append(subscriptions, s)
	}

	if err := i.Err(); err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func cancelSubscription(sc *client.API, subscriptionID string) (*stripe.Subscription, error) {
	return sc.Subscriptions.Cancel(subscriptionID, nil)
}

func getCustomer(sc *client.API, customerID string) (*stripe.Customer, error) {
	return sc.Customers.Get(customerID, nil)
}

func updateCustomer(sc *client.API, customerID, email, name string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
		Name:  stripe.String(name),
	}
	return sc.Customers.Update(customerID, params)
}

func deleteCustomer(sc *client.API, customerID string) (*stripe.Customer, error) {
	return sc.Customers.Del(customerID, nil)
}

func listCustomers(sc *client.API) ([]*stripe.Customer, error) {
	var customers []*stripe.Customer
	params := &stripe.CustomerListParams{}
	i := sc.Customers.List(params)

	for i.Next() {
		c := i.Customer()
		customers = append(customers, c)
	}

	if err := i.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}
