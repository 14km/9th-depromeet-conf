package api

import (
	"github.com/go-resty/resty/v2"
)

var weatherMapHost string
var weatherMapApiKey string

func SettingByWeatherMap(host string, apiKey string) {
	weatherMapHost = host
	weatherMapApiKey = apiKey
}

func GetWeatherByWeatherMapApi(lat string, lon string) (body []byte, err error) {
	url := weatherMapHost + "/data/2.5/forecast"

	client := resty.New()

	response, err := client.R().
		SetQueryParams(map[string]string{
			"lat":   lat,
			"lon":   lon,
			"appid": weatherMapApiKey,
		}).
		Get(url)

	if err != nil {
		panic(err)
	}

	return response.Body(), nil
}

func GetAirPollutionByWeatherMapApi(lat string, lon string) (body []byte, err error) {
	url := weatherMapHost + "/data/2.5/air_pollution"

	client := resty.New()

	response, err := client.R().
		SetQueryParams(map[string]string{
			"lat":   lat,
			"lon":   lon,
			"appid": weatherMapApiKey,
		}).
		Get(url)

	if err != nil {
		panic(err)
	}

	return response.Body(), nil
}

func GetWeatherForWeekByWeatherMapApi(lat string, lon string) (body []byte, err error) {
	url := weatherMapHost + "/data/2.5/onecall"

	client := resty.New()

	response, err := client.R().
		SetQueryParams(map[string]string{
			"lat":     lat,
			"lon":     lon,
			"appid":   weatherMapApiKey,
			"exclude": "minutely,hourly",
		}).
		Get(url)

	if err != nil {
		panic(err)
	}

	return response.Body(), nil
}
