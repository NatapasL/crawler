package siamintershop

import "manga-crawler/internal/infrastructure/api_gateway/siamintershop"

type siamintershopApiModule struct {
	ApiRequest          *siamintershop.ApiRequest
	ProductSearchApi    *siamintershop.ProductSearchApi
	GetProductDetailApi *siamintershop.GetProductDetailApi
	GetCategoryListApi  *siamintershop.GetCategoryListApi
}

type siamintershopApiModuleDependency struct{}

func initializeSiamintershopApiModule(deps siamintershopApiModuleDependency) siamintershopApiModule {
	apiRequest := siamintershop.ApiRequest{}
	productSearchApi := siamintershop.NewProductSearchApi(
		siamintershop.ProductSearchApiDependencies{Request: apiRequest},
	)
	getProductDetailApi := siamintershop.NewGetProductDetailApi(
		siamintershop.GetProductDetailApiDependencies{Request: apiRequest},
	)
	getCategoryListApi := siamintershop.NewGetCategoryListApi(
		siamintershop.GetCategoryListApiDependencies{Request: apiRequest},
	)

	return siamintershopApiModule{
		ApiRequest:          &apiRequest,
		ProductSearchApi:    productSearchApi,
		GetProductDetailApi: getProductDetailApi,
		GetCategoryListApi:  getCategoryListApi,
	}
}
