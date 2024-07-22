package container

import siamintershopgapi "manga-crawler/internal/infrastructure/api_gateway/siamintershop_api"

type siamintershopApiModule struct {
	ApiRequest          *siamintershopgapi.ApiRequest
	ProductSearchApi    *siamintershopgapi.ProductSearchApi
	GetProductDetailApi *siamintershopgapi.GetProductDetailApi
	GetCategoryListApi  *siamintershopgapi.GetCategoryListApi
}

type siamintershopApiModuleDependency struct{}

func initializeSiamintershopApiModule(deps siamintershopApiModuleDependency) siamintershopApiModule {
	apiRequest := siamintershopgapi.ApiRequest{}
	productSearchApi := siamintershopgapi.NewProductSearchApi(
		siamintershopgapi.ProductSearchApiDependencies{Request: apiRequest},
	)
	getProductDetailApi := siamintershopgapi.NewGetProductDetailApi(
		siamintershopgapi.GetProductDetailApiDependencies{Request: apiRequest},
	)
	getCategoryListApi := siamintershopgapi.NewGetCategoryListApi(
		siamintershopgapi.GetCategoryListApiDependencies{Request: apiRequest},
	)

	return siamintershopApiModule{
		ApiRequest:          &apiRequest,
		ProductSearchApi:    productSearchApi,
		GetProductDetailApi: getProductDetailApi,
		GetCategoryListApi:  getCategoryListApi,
	}
}
