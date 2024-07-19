package siamintershop

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Filter struct {
	Limit             int    `json:"limit"`
	Offset            int    `json:"offset"`
	CategoryId        string `json:"category_id"`
	CategoryWithChild bool   `json:"category_with_child"`
}

type SiamintershopResponse struct {
	Total    string                 `json:"total,omitempty"`
	Limit    string                 `json:"limit,omitempty"`
	Offset   int                    `json:"offset,omitempty"`
	Products []SiamintershopProduct `json:"products,omitempty"`
}

type SiamintershopProduct struct {
	ProductId        string `json:"product_id,omitempty"`
	ProductShortUrl  string `json:"product_short_url,omitempty"`
	ProductPrice     string `json:"product_price,omitempty"`
	ProductFullPrice string `json:"product_full_price,omitempty"`
}

// https://siamintershop.com/api/v1/products/search?filter={%22limit%22:60,%22offset%22:0,%22category_id%22:%22654%22,%22category_with_child%22:true}&include=shop_mini,dropship
const BaseUrl = "https://siamintershop.com/api/v1/products/search"
const CategoryId = "654"

func buildUrl(f Filter) (string, error) {
	req, err := http.NewRequest("GET", BaseUrl, nil)
	if err != nil {
		return "", err
	}

	filterString, err := json.Marshal(f)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("filter", string(filterString))
	q.Add("include", "shop_mini,dropship")
	req.URL.RawQuery = q.Encode()

	log.Println(req.URL.String())

	return req.URL.String(), nil
}

func callApi(url string) ([]byte, error) {
	var emptyByte []byte
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return emptyByte, nil
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return emptyByte, nil
	}
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return emptyByte, err
	}

	return responseData, nil
}

func parseJson(responseData []byte) (*SiamintershopResponse, error) {
	var restruct SiamintershopResponse
	json.Unmarshal(responseData, &restruct)

	return &restruct, nil
}

func Scrape() {

	f := Filter{
		Limit:             1,
		Offset:            0,
		CategoryId:        CategoryId,
		CategoryWithChild: true,
	}
	url, err := buildUrl(f)
	if err != nil {
		log.Println(err)
		return
	}

	responseData, err := callApi(url)
	if err != nil {
		log.Println(err)
		return
	}

	resStruct, err := parseJson(responseData)
	if err != nil {
		log.Println(err)
		return
	}

	total, _ := strconv.Atoi(resStruct.Total)

	time.Sleep(time.Second * 5)

	i := 0
	chunkSize := 59
	for (i * chunkSize) < total {
		func() {
			defer func() { i++ }()

			f := Filter{
				Limit:             chunkSize,
				Offset:            i * chunkSize,
				CategoryId:        CategoryId,
				CategoryWithChild: true,
			}

			url, err := buildUrl(f)
			if err != nil {
				log.Println(err)
				return
			}

			responseData, err := callApi(url)
			if err != nil {
				log.Println(err)
				return
			}

			_, err = parseJson(responseData)
			if err != nil {
				log.Println(err)
				return
			}
			// log.Printf("%+v", resStruct)

			time.Sleep(time.Second * 5)
		}()

	}
}
