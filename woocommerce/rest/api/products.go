package api

import (
	"context"
	"net/url"

	"github.com/evertonvps/go-wc-client/woocommerce/rest"
	"github.com/evertonvps/go-wc-client/woocommerce/rest/models"
)

// products implements ProductInterface
type productsApi struct {
	client           rest.Interface
	ProductInterface ProductInterface
}

// ProductInterface allows you to create, view, update, and delete individual, or a batch, of products
// https://woocommerce.github.io/woocommerce-rest-api-docs/#products
type ProductInterface interface {
	//Create()
	Get(ctx context.Context, params url.Values) (*[]models.Product, error)
	///FindBy(ctx context.Context, data interface{}) (*models.Product, error)
	FindBySKU(ctx context.Context, sku string) (*models.Product, error)
	//Delete()
	//Update()
	//BatchUpdate()
}

// newProducts returns a products
func newProducts(c *WcClient) *productsApi {
	return &productsApi{
		client: c.RESTClient(),
	}
}

// wp-json/wc/v3/products
func (p *productsApi) Get(ctx context.Context, params url.Values) (products *[]models.Product, err error) {

	products = &[]models.Product{}
	err = p.client.Get(ctx, ENDPOINT_PRODUCTS, params, &products)

	return
}

// Find SKU of the product, and returns the corresponding  object, and an error if there is any.
// wp-json/wc/v3/products
func (p *productsApi) FindBySKU(ctx context.Context, sku string) (product *models.Product, err error) {

	var products []models.Product
	v := url.Values{}
	v.Set("sku", sku)

	err = p.client.Get(ctx, ENDPOINT_PRODUCTS, v, &products)
	if len(products) > 0 {
		product = &products[0]
	}

	return
}
