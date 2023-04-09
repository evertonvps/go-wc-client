package models

import (
	"encoding/json"
	"testing"
)

var productsJson = `[
	{
	  "id": 33182,
	  "name": "PRODUCT TEST 01",
	  "slug": "product-test-01",
	  "permalink": "https://fake-store.com/produto/product-test-01/",
	  "date_created": "2023-02-25T17:40:34",
	  "date_created_gmt": "2023-02-25T20:40:34",
	  "date_modified": "2023-02-25T17:41:22",
	  "date_modified_gmt": "2023-02-25T20:41:22",
	  "type": "simple",
	  "status": "publish",
	  "featured": false,
	  "catalog_visibility": "visible",
	  "description": "<p>PRODUCT TEST 01</p>",
	  "short_description": "<p>PRODUCT TEST 01</p>\n",
	  "sku": "SKU33182",
	  "price": "8.9925",
	  "regular_price": "8.9925",
	  "sale_price": "",
	  "date_on_sale_from": null,
	  "date_on_sale_from_gmt": null,
	  "date_on_sale_to": null,
	  "date_on_sale_to_gmt": null,
	  "on_sale": false,
	  "purchasable": true,
	  "total_sales": 0,
	  "virtual": false,
	  "downloadable": false,
	  "downloads": [],
	  "download_limit": -1,
	  "download_expiry": -1,
	  "external_url": "",
	  "button_text": "",
	  "tax_status": "taxable",
	  "tax_class": "",
	  "manage_stock": true,
	  "stock_quantity": 3,
	  "backorders": "no",
	  "backorders_allowed": false,
	  "backordered": false,
	  "low_stock_amount": null,
	  "sold_individually": false,
	  "weight": "0.000",
	  "dimensions": {
		"length": "0.00",
		"width": "0.00",
		"height": "0.00"
	  },
	  "shipping_required": true,
	  "shipping_taxable": true,
	  "shipping_class": "",
	  "shipping_class_id": 0,
	  "reviews_allowed": true,
	  "average_rating": "0.00",
	  "rating_count": 0,
	  "upsell_ids": [],
	  "cross_sell_ids": [],
	  "parent_id": 0,
	  "purchase_note": "",
	  "categories": [
		{
		  "id": 16,
		  "name": "Category 1",
		  "slug": "category1"
		}
	  ],
	  "tags": [],
	  "images": [
		{
		  "id": 33181,
		  "date_created": "2023-02-25T14:40:33",
		  "date_created_gmt": "2023-02-25T20:40:33",
		  "date_modified": "2023-02-25T14:40:33",
		  "date_modified_gmt": "2023-02-25T20:40:33",
		  "src": "https://fake-store.wp.com/wp-content/uploads/2023/02/faker.jpg",
		  "name": "faker.jpg",
		  "alt": ""
		}
	  ],
	  "attributes": [
		{
		  "id": 0,
		  "name": "test",
		  "position": 1,
		  "visible": true,
		  "variation": false,
		  "options": [
			"attributes Test"
		  ]
		}
	  ],
	  "default_attributes": [],
	  "variations": [],
	  "grouped_products": [],
	  "menu_order": 0,
	  "price_html": "<span class=\"woocommerce-Price-amount amount\"><bdi><span class=\"woocommerce-Price-currencySymbol\">&#82;&#36;</span>8,99</bdi></span>",

	  "stock_status": "instock",
	  "has_options": false,
	  "_links": {
		"self": [
		  {
			"href": "https://fake-store.com/wp-json/wc/v3/products/33182"
		  }
		],
		"collection": [
		  {
			"href": "https://fake-store.com/wp-json/wc/v3/products"
		  }
		]
	  }
	}
  ]`

func TestUnmarshalJsonList(t *testing.T) {

	var products []Product
	if err := json.Unmarshal([]byte(productsJson), &products); err == nil { // Parse []byte to the go struct pointer
		println("Product:", products[0].Name, products[0].Price)
	} else {

		t.Errorf("Unmarshal fail: %v", err.Error())
		t.FailNow()
	}

}
