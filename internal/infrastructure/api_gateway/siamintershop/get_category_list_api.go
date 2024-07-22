package siamintershop

import (
	"encoding/json"
	"manga-crawler/config"
	"manga-crawler/internal/domain/scraper/siamintershop/categorylist"
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

func (gcl GetCategoryListApi) Request() ([]categorylist.Category, error) {
	req, err := http.NewRequest("GET", CategoryListBaseUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", config.GetConfig().Http.UserAgent)

	response, err := gcl.request.Request(req)
	if err != nil {
		return nil, err
	}

	catories := gcl.convertToCategoryList(response)

	return catories, nil
}

func (GetCategoryListApi) convertToCategoryList(data []byte) []categorylist.Category {
	var categories []categorylist.Category
	json.Unmarshal(data, &categories)

	return categories
}
