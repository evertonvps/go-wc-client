# WooCommerce API - Golang Client

A Golang wrapper for the WooCommerce REST API. Easily interact with the WooCommerce REST API using this library.
This is a fork of `darh/wc-api-golang`. Main difference: 
Basically the Client of `darh/wc-api-golang` is the core, and the WcClient encapsulates it together with the services (Products, Orders, etc). 

This lib is **not** backward compatible with darh/wc-api-golang!

## Installation

```bash
$ go get github.com/evertonvps/go-wc-client
```

## Getting started

Generate API credentials (Consumer Key & Consumer Secret) following this instructions <http://docs.woocommerce.com/document/woocommerce-rest-api/>
.

Check out the WooCommerce API endpoints and data that can be manipulated in <https://woocommerce.github.io/woocommerce-rest-api-docs/>.

## Setup

Setup for the new WP REST API integration (WooCommerce 2.6 or later):

```golang
import (
	"github.com/evertonvps/go-wc-client/woocommerce/rest"
	"github.com/evertonvps/go-wc-client/woocommerce/rest/api"
)

	woocommerce, err := api.NewWoocommerceClient("https://fake-store.com", &rest.ApiConfig{
		API:            true,
		APIPrefix:      "/wp-json/wc",
		Version:        "v3",
		ConsumerKey:    "ck_",
		ConsumerSecret: "cs_",
	})
```

### Paramaters

|       ApiConfig   |   Type   |                Description                 |
| ----------------- | -------- | ------------------------------------------ |
| `store`           | `string` | Your Store URL, example: http://woo.dev/   |
| `apiConfig`       | `*rest.ApiConfig`  | Extra arguments (see client options table) |

#### Client options

|        Option       |   Type   |                                                      Description                                                       |
|---------------------|----------|------------------------------------------------------------------------------------------------------------------------|
| `API`            | `bool`   | Allow make requests to the new WP REST API integration (WooCommerce 2.6 or later)                                      |
| `APIPrefix`     | `string` | Custom WP REST API URL prefix, used to support custom prefixes created with the `rest_url_prefix` filter               |
| `Version`           | `string` | API version, default is `v3`                                                                                           |
| `Timeout`           | `time.Duration`    | Request timeout, default is `15`                                                                                       |
| `VerifySSL`        | `bool`   | Verify SSL when connect, use this option as `false` when need to test with self-signed certificates, default is `true` |
| `QueryStringAuth` | `bool`   | Force Basic Authentication as query string when `true` and using under HTTPS, default is `false`                       |
| `OauthTimestamp`   | `time.Time` | Custom oAuth timestamp, default is `time.Now()`                                                                            |
| `ck`              | `string` | Your API consumer key                      |
| `cs`              | `string` | Your API consumer secret                   |

## Interfaces
### Get Products
```golang
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
```
----

## Methods

|    Params    |   Type   |                         Description                          |
| ------------ | -------- | ------------------------------------------------------------ |
| `endpoint`   | `string` | WooCommerce API endpoint, example: `customers` or `order/12` |
| `data`       | `interface{}`  | Only for POST and PUT, data that will be converted to JSON   |
| `parameters` | `url.Values`  | Only for GET and DELETE, request query string                |

### GET

```golang
rc, err := woocommerce.Get(ctx, endpoint, parameters)
```

### POST

```golang
rc, err := woocommerce.Post(ctx, endpoint, data)
```

### PUT

```golang
rc, err := woocommerce.Put(ctx, endpoint, data)
```

### DELETE

```golang
rc, err := woocommerce.Delete(ctx, endpoint, parameters)
```

### OPTIONS

```golang
rc, err := woocommerce.Options(ctx, endpoint)
```

#### Response

All methods will return a `io.ReadCloser` and `nil` on success or an `error` on failure.

## Release History
