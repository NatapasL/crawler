package siamintershop

type SiamintershopProduct struct {
	ProductId        string   `json:"product_id,omitempty"`
	ProductShortUrl  string   `json:"product_short_url,omitempty"`
	ProductPrice     string   `json:"product_price,omitempty"`
	ProductFullPrice string   `json:"product_full_price,omitempty"`
	ProductName      string   `json:"product_name,omitempty"`
	ProductTitle     string   `json:"product_title,omitempty"`
	ProductTags      []string `json:"product_tags,omitempty"`
}
