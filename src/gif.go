package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	GIF_CMD = `/gif`
	GIF_URL = `https://api.tenor.com/v1/random?q=%s&key=%s&limit=1&ar_range=standard&media_filter=minimal&locale=en_US&contentfilter=high`
)

type Media map[string]struct {
	URL string `json:"url"`
}

type GIF struct {
	Results []struct {
		Media []Media `json:"media"`
	} `json:"results"`
}

func (g GIF) String() string {
	if len(g.Results) < 1 || len(g.Results[0].Media) < 1 || len(g.Results[0].Media[0]) < 1 {
		return "error"
	}
	media, ok := g.Results[0].Media[0]["tinygif"]
	if !ok {
		return "nok"
	}
	return strings.TrimSpace(media.URL)
}

func getGif(name string) (*GIF, error) {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		name = "random"
	}

	rsp, err := httpCli.Get(fmt.Sprintf(GIF_URL, name, gifToken))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	var gif GIF
	if err := json.Unmarshal(b, &gif); err != nil {
		return nil, err
	}
	return &gif, nil
}
