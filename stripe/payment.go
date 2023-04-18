package stripe

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

func getPaymentMethod(sc *client.API, paymentMethodID string) (*stripe.PaymentMethod, error) {
	return sc.PaymentMethods.Get(paymentMethodID, nil)
}

func refundCharge(sc *client.API, chargeID string, amount int64, reason stripe.RefundReason) (*stripe.Refund, error) {
	params := &stripe.RefundParams{
		Charge: stripe.String(chargeID),
		Amount: stripe.Int64(amount),
		Reason: (*string)(&reason),
	}
	return sc.Refunds.New(params)
}

func createPaymentSource(sc *client.API, customerID, token string) (*stripe.PaymentSource, error) {
	params := &stripe.CustomerSourceParams{
		Customer: stripe.String(customerID),
		Source: &stripe.SourceParams{
			Token: stripe.String(token),
		},
	}
	return sc.CustomerSources.New(params), nil
}

func chargeCustomer(sc *client.API, customerID, sourceID string, amount int64, currency, description string) (*stripe.Charge, error) {
	src, err := getSource(sourceID)
	if err != nil {
		panic(err)
	}
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String(currency),
		Description: stripe.String(description),
		Customer:    stripe.String(customerID),
		Source:      src,
	}
	return sc.Charges.New(params)
}
