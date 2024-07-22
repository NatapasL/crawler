package siamintershop

import (
	"encoding/json"
	"fmt"
	"manga-crawler/config"
	"manga-crawler/internal/domain/scraper/siamintershop/productdetail"
	"net/http"
)

const GetProductDetailBaseUrl = "https://siamintershop.com/api/v1/products/%s/detail"

type GetProductDetailApi struct {
	request ApiRequest
}

type GetProductDetailApiDependencies struct {
	Request ApiRequest
}

func NewGetProductDetailApi(deps GetProductDetailApiDependencies) *GetProductDetailApi {
	return &GetProductDetailApi{request: deps.Request}
}

func (pdf GetProductDetailApi) Request(productId string) (*productdetail.ProductDetail, error) {
	req, err := pdf.buildRequest(productId)
	if err != nil {
		return nil, err
	}

	response, err := pdf.request.Request(req)
	if err != nil {
		return nil, err
	}

	productDetail := pdf.convertToProductDetail(response)

	return productDetail, nil
}

func (GetProductDetailApi) buildRequest(productId string) (*http.Request, error) {
	url := fmt.Sprintf(GetProductDetailBaseUrl, productId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", config.GetConfig().Http.UserAgent)

	return req, nil
}

func (GetProductDetailApi) convertToProductDetail(data []byte) *productdetail.ProductDetail {
	var productDetail productdetail.ProductDetail
	json.Unmarshal(data, &productDetail)

	return &productDetail
}
