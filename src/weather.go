package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	W_URL = `http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric`
	W_CMD = `/weather`
)

type Weather struct {
	Name string `json:"name"`
	Data []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
	} `json:"wind"`
}

func (w Weather) String() string {
	if len(w.Data) < 1 {
		return "error"
	}
	return fmt.Sprintf("City: %s\nDesc: %s\nTemp: %.2fC (feels: %.2fC)\nWind: %.2f",
		w.Name,
		w.Data[0].Main+" | "+w.Data[0].Description,
		w.Main.Temp,
		w.Main.FeelsLike,
		w.Wind.Speed,
	)
}

func getWeather(city string) (*Weather, error) {
	rsp, err := httpCli.Get(fmt.Sprintf(W_URL, strings.TrimSpace(city), weatherToken))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	var w Weather
	if err := json.Unmarshal(b, &w); err != nil {
		return nil, err
	}
	return &w, nil
}
