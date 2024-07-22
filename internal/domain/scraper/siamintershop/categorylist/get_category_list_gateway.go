package categorylist

type GetCategoryListGateway interface {
	Request() ([]Category, error)
}
