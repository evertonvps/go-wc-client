package main

import (
	"context"
	"fmt"

	"github.com/evertonvps/go-wc-client/woocommerce/rest"
	"github.com/evertonvps/go-wc-client/woocommerce/rest/api"
)

func main() {

	const WC_URL = "https://fake-store.com"

	wc, err := api.NewWoocommerceClient(WC_URL, &rest.ApiConfig{
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

		products, _ := wc.Products().FindBySKU(ctx, "SKU1020397")

		if products != nil {
			fmt.Println("Product:", products.Name)
		} else {
			println("Not found")
		}

	}
}
