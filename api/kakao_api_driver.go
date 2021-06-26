package api

import (
	"github.com/go-resty/resty/v2"
)

var kakaoHost string
var kakaoApiKey string

func SettingByKakao(host string, apiKey string) {
	kakaoHost = host
	kakaoApiKey = apiKey
}

/**
url : https://dapi.kakao.com/v2/local/search/address?query=서초동
*/
func GetAddressContentsByKakaoApi(address string) (body []byte, err error) {
	url := kakaoHost + "/v2/local/search/address.json"

	client := resty.New()

	response, err := client.R().
		SetQueryParams(map[string]string{
			"query": address,
		}).
		SetAuthScheme("KakaoAK").
		SetAuthToken(kakaoApiKey).
		SetHeader("Content-Type", "application/json").
		Get(url)

	if err != nil {
		panic(err)
	}

	return response.Body(), nil
}
