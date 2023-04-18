package stripe

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

func createProduct(sc *client.API, name, description string) (*stripe.Product, error) {
	params := &stripe.ProductParams{
		Name:        stripe.String(name),
		Description: stripe.String(description),
	}
	return sc.Products.New(params)
}

// Create a price for a product:
func createPrice(sc *client.API, productID string, unitAmount int64, currency, interval string) (*stripe.Price, error) {
	params := &stripe.PriceParams{
		Product:    stripe.String(productID),
		UnitAmount: stripe.Int64(unitAmount),
		Currency:   stripe.String(currency),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(interval),
		},
	}
	return sc.Prices.New(params)
}

func getPrice(sc *client.API, priceID string) (*stripe.Price, error) {
	return sc.Prices.Get(priceID, nil)
}

func listPrices(sc *client.API) ([]*stripe.Price, error) {
	var prices []*stripe.Price
	params := &stripe.PriceListParams{}
	i := sc.Prices.List(params)

	for i.Next() {
		p := i.Price()
		prices = append(prices, p)
	}

	if err := i.Err(); err != nil {
		return nil, err
	}
	return prices, nil
}

func listProducts(sc *client.API) ([]*stripe.Product, error) {
	var products []*stripe.Product
	params := &stripe.ProductListParams{}
	i := sc.Products.List(params)

	for i.Next() {
		p := i.Product()
		products = append(products, p)
	}

	if err := i.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func getProduct(sc *client.API, productID string) (*stripe.Product, error) {
	return sc.Products.Get(productID, nil)
}

func updateProduct(sc *client.API, productID, name, description string) (*stripe.Product, error) {
	params := &stripe.ProductParams{
		Name:        stripe.String(name),
		Description: stripe.String(description),
	}
	return sc.Products.Update(productID, params)
}

func deleteProduct(sc *client.API, productID string) (*stripe.Product, error) {
	return sc.Products.Del(productID, nil)
}
