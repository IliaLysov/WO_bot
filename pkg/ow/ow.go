package ow

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	Name string `json:"name"`

	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`

	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
	} `json:"main"`

	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

type Config struct {
	Key string `envconfig:"OW_KEY" required:"true"`
}

type OW struct {
	Config
}

func (ow *OW) Get(city string) (string, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=ru", city, ow.Key)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return "Город не найден", nil
		}
		return "", errors.New(resp.Status)
	}
	var res WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	if len(res.Weather) == 0 {
		return "", errors.New(res.Name)
	}
	return fmt.Sprintf("Погода в %s: %.1f°C, %s", res.Name, res.Main.Temp, res.Weather[0].Description), nil
}
