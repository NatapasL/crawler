package siamintershop

import (
	"encoding/json"
	"fmt"
	"manga-crawler/config"
	"manga-crawler/internal/domain/scraper/siamintershop/subproduct"
	"net/http"
)

const GetSubProductsBaseUrl = "https://siamintershop.com/api/v1/products/%s/subproducts"

type GetSubproductsApi struct {
	request ApiRequest
}

type GetSubproductsApiDependencies struct {
	Request ApiRequest
}

func NewGetSubproductsApi(deps GetSubproductsApiDependencies) *GetSubproductsApi {
	return &GetSubproductsApi{request: deps.Request}
}

func (gsp GetSubproductsApi) Request(productId string) ([]subproduct.SubProduct, error) {
	var emptySubproducts []subproduct.SubProduct

	req, err := gsp.buildRequest(productId)
	if err != nil {
		return emptySubproducts, err
	}

	response, err := gsp.request.Request(req)
	if err != nil {
		return emptySubproducts, err
	}

	subProducts := gsp.convertResponseToSubProducts(response)

	return subProducts, nil
}

func (GetSubproductsApi) buildRequest(productId string) (*http.Request, error) {
	url := fmt.Sprintf(GetProductDetailBaseUrl, productId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", config.GetConfig().Http.UserAgent)

	return req, nil
}

func (GetSubproductsApi) convertResponseToSubProducts(data []byte) []subproduct.SubProduct {
	var subProducts []subproduct.SubProduct
	json.Unmarshal(data, &subProducts)

	return subProducts
}
