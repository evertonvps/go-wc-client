package models

// Product represent WooCommerce Product
// https://woocommerce.github.io/woocommerce-rest-api-docs/#product-properties
type Product struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	Slug              string        `json:"slug"`
	Permalink         string        `json:"permalink"`
	DateCreated       string        `json:"date_created"`
	DateCreatedGmt    string        `json:"date_created_gmt"`
	DateModified      string        `json:"date_modified"`
	DateModifiedGmt   string        `json:"date_modified_gmt"`
	Type              string        `json:"type"`
	Status            string        `json:"status"`
	Featured          bool          `json:"featured"`
	CatalogVisibility string        `json:"catalog_visibility"`
	Description       string        `json:"description"`
	ShortDescription  string        `json:"short_description"`
	Sku               string        `json:"sku"`
	Price             StringFloat   `json:"price"`
	RegularPrice      StringFloat   `json:"regular_price"`
	SalePrice         StringFloat   `json:"sale_price"`
	DateOnSaleFrom    interface{}   `json:"date_on_sale_from"`
	DateOnSaleFromGmt interface{}   `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      interface{}   `json:"date_on_sale_to"`
	DateOnSaleToGmt   interface{}   `json:"date_on_sale_to_gmt"`
	OnSale            bool          `json:"on_sale"`
	Purchasable       bool          `json:"purchasable"`
	TotalSales        int           `json:"total_sales"`
	Virtual           bool          `json:"virtual"`
	Downloadable      bool          `json:"downloadable"`
	Downloads         []interface{} `json:"downloads"`
	DownloadLimit     int           `json:"download_limit"`
	DownloadExpiry    int           `json:"download_expiry"`
	ExternalURL       string        `json:"external_url"`
	ButtonText        string        `json:"button_text"`
	TaxStatus         string        `json:"tax_status"`
	TaxClass          string        `json:"tax_class"`
	ManageStock       bool          `json:"manage_stock"`
	StockQuantity     int           `json:"stock_quantity"`
	Backorders        string        `json:"backorders"`
	BackordersAllowed bool          `json:"backorders_allowed"`
	Backordered       bool          `json:"backordered"`
	LowStockAmount    interface{}   `json:"low_stock_amount"`
	SoldIndividually  bool          `json:"sold_individually"`
	Weight            string        `json:"weight"`
	Dimensions        Dimensions    `json:"dimensions"`
	ShippingRequired  bool          `json:"shipping_required"`
	ShippingTaxable   bool          `json:"shipping_taxable"`
	ShippingClass     string        `json:"shipping_class"`
	ShippingClassID   int           `json:"shipping_class_id"`
	ReviewsAllowed    bool          `json:"reviews_allowed"`
	AverageRating     string        `json:"average_rating"`
	RatingCount       int           `json:"rating_count"`
	UpsellIds         []interface{} `json:"upsell_ids"`
	CrossSellIds      []interface{} `json:"cross_sell_ids"`
	ParentID          int           `json:"parent_id"`
	PurchaseNote      string        `json:"purchase_note"`
	Categories        []Category    `json:"categories"`
	Tags              []interface{} `json:"tags"`
	Images            []Image       `json:"images"`
	Attributes        []Attribute   `json:"attributes"`
	DefaultAttributes []interface{} `json:"default_attributes"`
	Variations        []interface{} `json:"variations"`
	GroupedProducts   []interface{} `json:"grouped_products"`
	MenuOrder         int           `json:"menu_order"`
	PriceHTML         string        `json:"price_html"`
	RelatedIds        []int         `json:"related_ids"`
	//MetaData          []MetaData    `json:"meta_data"`
	StockStatus string `json:"stock_status"`
	HasOptions  bool   `json:"has_options"`
}
type StringFloat string