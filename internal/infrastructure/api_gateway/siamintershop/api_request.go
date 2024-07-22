package siamintershop

import (
	"io"
	"net/http"
)

type ApiRequest struct{}

func (ApiRequest) Request(request *http.Request) ([]byte, error) {
	var emptyByte []byte

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return emptyByte, err
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return emptyByte, err
	}

	return responseData, nil
}
