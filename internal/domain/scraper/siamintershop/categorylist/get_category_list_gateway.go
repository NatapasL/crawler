package categorylist

type GetCategoryListGateway interface {
	Request() ([]byte, error)
}
