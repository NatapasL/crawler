package siamintershop

import (
	"manga-crawler/config"
	"net/http"
)

const CategoryListBaseUrl = "https://siamintershop.com/api/v1/categories/list"

type GetCategoryListApi struct {
	request ApiRequest
}

type GetCategoryListApiDependencies struct {
	Request ApiRequest
}

func NewGetCategoryListApi(deps GetCategoryListApiDependencies) *GetCategoryListApi {
	return &GetCategoryListApi{request: deps.Request}
}

func (gcl GetCategoryListApi) Request() ([]byte, error) {
	var emptyByte []byte

	req, err := http.NewRequest("GET", CategoryListBaseUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", config.GetConfig().Http.UserAgent)

	response, err := gcl.request.Request(req)
	if err != nil {
		return emptyByte, err
	}

	return response, nil
}
