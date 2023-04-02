package api

import (
	"github.com/evertonvps/go-wc-client/woocommerce/rest"
)

type WcClientInterface interface {
	RESTClient() rest.Interface
}

type WcClient struct {
	restClient rest.Interface
}

func (c *WcClient) Products() ProductInterface {
	return newProducts(c)
}

// RESTClient returns interface that is used to communicate
// with API server by this client implementation.
func (c *WcClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

// NewWoocommerceClient creates a new Client for the given api config,
// where httpClient was generated with rest.NewClient(store, apiConfig).
func NewWoocommerceClient(store string, apiConfig *rest.ApiConfig) (*WcClient, error) {
	client, err := rest.NewClient(store, apiConfig)
	if err != nil {
		return nil, err
	}

	return &WcClient{
		restClient: client,
	}, nil
}
