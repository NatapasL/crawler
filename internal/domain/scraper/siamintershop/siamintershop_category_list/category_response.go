package siamintershopcategorylist

import "encoding/json"

type CategoryResponse struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func NewCategoryResponseList(data []byte) []CategoryResponse {
	var responses []CategoryResponse
	json.Unmarshal(data, &responses)

	return responses
}
