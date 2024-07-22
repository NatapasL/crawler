package container

import siamintershopgapi "manga-crawler/internal/infrastructure/api_gateway/siamintershop_api"

type SiamintershopApiModule struct {
	ApiRequest          *siamintershopgapi.ApiRequest
	ProductSearchApi    *siamintershopgapi.ProductSearchApi
	GetProductDetailApi *siamintershopgapi.GetProductDetailApi
	GetCategoryListApi  *siamintershopgapi.GetCategoryListApi
}

type SiamintershopApiModuleDependency struct{}

func InitializeSiamintershopApiModule(deps SiamintershopApiModuleDependency) SiamintershopApiModule {
	apiRequest := siamintershopgapi.ApiRequest{}
	productSearchApi := siamintershopgapi.NewProductSearchApi(apiRequest)
	getProductDetailApi := siamintershopgapi.NewGetProductDetailApi(apiRequest)
	getCategoryListApi := siamintershopgapi.NewGetCategoryListApi(apiRequest)

	return SiamintershopApiModule{
		ApiRequest:          &apiRequest,
		ProductSearchApi:    productSearchApi,
		GetProductDetailApi: getProductDetailApi,
		GetCategoryListApi:  getCategoryListApi,
	}
}
