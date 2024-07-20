package siamintershopgapi

import (
	"encoding/json"
	"manga-crawler/config"
	"net/http"
)

const ProductSearchBaseUrl = "https://siamintershop.com/api/v1/products/search"

type productSearchFilter struct {
	Limit             int    `json:"limit"`
	Offset            int    `json:"offset"`
	CategoryId        string `json:"category_id"`
	CategoryWithChild bool   `json:"category_with_child"`
}

type ProductSearchApi struct {
	request ApiRequest
}

func NewProductSearchApi(request ApiRequest) *ProductSearchApi {
	return &ProductSearchApi{request}
}

func (psf ProductSearchApi) Request(categoryId string, offset int, limit int) ([]byte, error) {
	filter := psf.buildFilter(categoryId, offset, limit)
	req, err := psf.buildRequest(filter)
	if err != nil {
		return nil, err
	}

	response, err := psf.request.Request(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (psf ProductSearchApi) buildFilter(categoryId string, offset int, limit int) productSearchFilter {
	return productSearchFilter{
		Offset:            offset,
		Limit:             limit,
		CategoryId:        categoryId,
		CategoryWithChild: true,
	}
}

func (psf ProductSearchApi) buildRequest(filter productSearchFilter) (*http.Request, error) {
	url, err := psf.buildUrl(filter)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", config.GetConfig().Http.UserAgent)

	return req, nil
}

func (ProductSearchApi) buildUrl(filter productSearchFilter) (string, error) {
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
