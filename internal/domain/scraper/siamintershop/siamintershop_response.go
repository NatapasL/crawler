package siamintershop

import "encoding/json"

type SiamintershopResponse struct {
	Total    string                 `json:"total,omitempty"`
	Limit    string                 `json:"limit,omitempty"`
	Offset   int                    `json:"offset,omitempty"`
	Products []SiamintershopProduct `json:"products,omitempty"`
}

func NewSiamintershopResponseFromBytes(data []byte) *SiamintershopResponse {
	var siamintershopResponse SiamintershopResponse
	json.Unmarshal(data, &siamintershopResponse)

	return &siamintershopResponse
}
