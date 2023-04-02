package main

import (
	"context"
	"fmt"

	"github.com/evertonvps/go-wc-client/woocommerce/rest"
	"github.com/evertonvps/go-wc-client/woocommerce/rest/api"
	"github.com/evertonvps/go-wc-client/woocommerce/rest/models"
)

func main() {

	const WC_URL = "https://fake-store.com"

	woocommerce, err := api.NewWoocommerceClient(WC_URL, &rest.ApiConfig{
		API:            true,
		APIPrefix:      "/wp-json/wc",
		Version:        "v3",
		ConsumerKey:    "ck_",
		ConsumerSecret: "cs_",
	})
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		ctx := context.Background()
		var products *[]models.Product
		products, err := woocommerce.Products().Get(ctx, map[string][]string{})

		if err == nil {
			for _, product := range *products {
				fmt.Println("Product: ", product.Name)
			}

		} else {
			println("Error:", err.Error())
		}

	}
}
