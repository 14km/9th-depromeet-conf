package main

import (
	"encoding/json"
	"github.com/14km/9th-depromeet-conf/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type response struct {
	Documents []document `json:"documents"`
}

type document struct {
	AddressName string `json:"address_name"`
	AddressType string `json:"address_type"`
	RoadAddress string `json:"road_address"`
	X           string `json:"x"`
	Y           string `json:"y"`
}

type responseDto struct {
	Weather      WeatherDto      `json:"weather"`
	AirPollution AirPollutionDto `json:"airPollution"`
}

type responseWeekDto struct {
	Weather      WeatherWeekDto  `json:"weather"`
	AirPollution AirPollutionDto `json:"airPollution"`
}

type WeatherDto struct {
	List interface{} `json:"list"`
	City interface{} `json:"city"`
}

type WeatherWeekDto struct {
	Lat      interface{} `json:"lat"`
	Lon      interface{} `json:"lon"`
	Timezone interface{} `json:"timezone"`
	Current  interface{} `json:"current"`
	Daily    interface{} `json:"daily"`
}

type AirPollutionDto struct {
	List interface{} `json:"list"`
}

func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.CORS())

	// Routes
	e.GET("/ping", hello)

	e.GET("/kakao", kakao)

	e.GET("/weather-map", weather)

	e.GET("/weather", addressAfterWeather)

	e.GET("/weather-air", weatherAndAir)

	e.GET("/weather-week-air", weatherForWeekAndAir)

	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!!! V3")
}

func addressAfterWeather(c echo.Context) error {
	address := c.QueryParam("address")
	api.SettingByKakao("https://dapi.kakao.com", "1234")
	api.SettingByWeatherMap("https://api.openweathermap.org", "1234")

	var info response
	kakaoApi, _ := api.GetAddressContentsByKakaoApi(address)
	if err := json.Unmarshal(kakaoApi, &info); err != nil {
		panic(err)
	}

	lat := info.Documents[0].Y
	lon := info.Documents[0].X

	weather, _ := api.GetWeatherByWeatherMapApi(lat, lon)
	airPollution, _ := api.GetAirPollutionByWeatherMapApi(lat, lon)

	var weatherDto WeatherDto
	var airPollutionDto AirPollutionDto
	_ = json.Unmarshal(weather, &weatherDto)
	_ = json.Unmarshal(airPollution, &airPollutionDto)

	u := &responseDto{
		Weather:      weatherDto,
		AirPollution: airPollutionDto,
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(u)
}

func weatherAndAir(c echo.Context) error {
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")

	api.SettingByWeatherMap("https://api.openweathermap.org", "1234")

	weather, _ := api.GetWeatherByWeatherMapApi(lat, lon)
	airPollution, _ := api.GetAirPollutionByWeatherMapApi(lat, lon)

	var weatherDto WeatherDto
	var airPollutionDto AirPollutionDto
	_ = json.Unmarshal(weather, &weatherDto)
	_ = json.Unmarshal(airPollution, &airPollutionDto)

	u := &responseDto{
		Weather:      weatherDto,
		AirPollution: airPollutionDto,
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(u)
}

func weatherForWeekAndAir(c echo.Context) error {
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")

	api.SettingByWeatherMap("https://api.openweathermap.org", "1234")

	weather, _ := api.GetWeatherForWeekByWeatherMapApi(lat, lon)
	airPollution, _ := api.GetAirPollutionByWeatherMapApi(lat, lon)

	var weatherDto WeatherWeekDto
	var airPollutionDto AirPollutionDto
	_ = json.Unmarshal(weather, &weatherDto)
	_ = json.Unmarshal(airPollution, &airPollutionDto)

	u := &responseWeekDto{
		Weather:      weatherDto,
		AirPollution: airPollutionDto,
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(u)
}

func kakao(c echo.Context) error {
	address := c.QueryParam("address")
	api.SettingByKakao("https://dapi.kakao.com", "1234")

	kakaoApi, _ := api.GetAddressContentsByKakaoApi(address)

	return c.String(http.StatusOK, string(kakaoApi))
}

func weather(c echo.Context) error {
	//lat string, lon string
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")

	api.SettingByWeatherMap("https://api.openweathermap.org", "1234")

	weather, _ := api.GetWeatherByWeatherMapApi(lat, lon)

	return c.String(http.StatusOK, string(weather))
}
