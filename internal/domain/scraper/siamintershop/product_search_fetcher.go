package siamintershop

import (
	"encoding/json"
	"io"
	"net/http"
)

const ProductSearchBaseUrl = "https://siamintershop.com/api/v1/products/search"

type productSearchFilter struct {
	Limit             int    `json:"limit"`
	Offset            int    `json:"offset"`
	CategoryId        string `json:"category_id"`
	CategoryWithChild bool   `json:"category_with_child"`
}

type ProductSearchFetcher struct{}

func NewProductSearchFetcher() *ProductSearchFetcher {
	return &ProductSearchFetcher{}
}

func (psf ProductSearchFetcher) Fetch(filter productSearchFilter) ([]byte, error) {
	url, err := psf.buildUrl(filter)
	if err != nil {
		return nil, err
	}

	response, err := psf.fetchApi(url)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (psf ProductSearchFetcher) buildUrl(filter productSearchFilter) (string, error) {
	req, err := http.NewRequest("GET", ProductSearchBaseUrl, nil)
	if err != nil {
		return "", err
	}

	filterString, err := json.Marshal(filter)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("filter", string(filterString))
	q.Add("include", "shop_mini,dropship")
	req.URL.RawQuery = q.Encode()

	return req.URL.String(), nil
}

func (psf ProductSearchFetcher) fetchApi(url string) ([]byte, error) {
	var emptyByte []byte
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return emptyByte, nil
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return emptyByte, nil
	}
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return emptyByte, err
	}

	return responseData, nil
}
