package productdetail

type ProductDetail struct {
	ProductId    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductTitle string `json:"product_title"`

	ProductMinPrice  float64 `json:"product_min_price"`
	ProductMaxPrice  float64 `json:"product_max_price"`
	ProductFullPrice float64 `json:"product_full_price"`

	ProductSalePercent float64 `json:"product_sale_percent"`

	Flag       string    `json:"subproduct_flag"`
	Suboptions SubOption `json:"suboptions"`
}

type SubOption struct {
	Options map[string]SubProductOption `json:"options"`
}

type SubProductOption struct {
	Key   string       `json:"key"`
	Name  string       `json:"name"`
	Items []SubProduct `json:"items"`
}

type SubProduct struct {
	Key        string   `json:"key"`
	Text       string   `json:"text"`
	ProductIds []string `json:"product_ids"`
}
