package siamintershopgapi

import (
	"fmt"
	"manga-crawler/config"
	"net/http"
)

const GetProductDetailBaseUrl = "https://siamintershop.com/api/v1/products/%s/detail"

type GetProductDetailApi struct {
	request ApiRequest
}

func NewGetProductDetailApi(request ApiRequest) *GetProductDetailApi {
	return &GetProductDetailApi{request}
}

func (pdf GetProductDetailApi) Request(productId string) ([]byte, error) {
	var emptyByte []byte

	req, err := pdf.buildRequest(productId)
	if err != nil {
		return emptyByte, err
	}

	response, err := pdf.request.Request(req)
	if err != nil {
		return emptyByte, err
	}

	return response, nil
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
