package siamintershop

import (
	"io"
	"log"
	"net/http"
	"time"
)

type ApiRequest struct{}

func (r ApiRequest) Request(request *http.Request) ([]byte, error) {
	var emptyByte []byte

	res, err := r.doRequest(request)
	if err != nil {
		log.Printf("%s, retrying...", err.Error())

		res, err = r.retry(request)
		if err != nil {
			return emptyByte, err
		}
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return emptyByte, err
	}

	return responseData, nil
}

func (ApiRequest) doRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(request)
}

func (r ApiRequest) retry(request *http.Request) (*http.Response, error) {
	var res *http.Response
	var err error

	retryDelays := [3]time.Duration{1 * time.Second, 3 * time.Second, 9 * time.Second}
	for _, delay := range retryDelays {
		time.Sleep(delay)

		res, err = r.doRequest(request)
		if err == nil {
			return res, err
		}
	}

	return nil, err
}
