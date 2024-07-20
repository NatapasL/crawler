package siamintershopproductdetail

import "encoding/json"

type ProductDetailResponse struct {
	ProductDetailResponseSubProductList

	ProductId    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductTitle string `json:"product_title"`

	ProductMinPrice  float64 `json:"product_min_price"`
	ProductMaxPrice  float64 `json:"product_max_price"`
	ProductFullPrice float64 `json:"product_full_price"`

	ProductSalePercent float64 `json:"product_sale_percent"`
}

type ProductDetailResponseSubProductList struct {
	Flag       string                          `json:"subproduct_flag"`
	Suboptions ProductDetailResponseSubOptions `json:"suboptions"`
}

type ProductDetailResponseSubOptions struct {
	Options map[string]ProductDetailResponseSubOptionsOptions `json:"options"`
}

type ProductDetailResponseSubOptionsOptions struct {
	Key   string                                `json:"key"`
	Name  string                                `json:"name"`
	Items []ProductDetailResponseSubOptionsItem `json:"items"`
}

type ProductDetailResponseSubOptionsItem struct {
	Key        string   `json:"key"`
	Text       string   `json:"text"`
	ProductIds []string `json:"product_ids"`
}

func NewProductDetailResponse(data []byte) *ProductDetailResponse {
	var response ProductDetailResponse
	json.Unmarshal(data, &response)

	return &response
}
