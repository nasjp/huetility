package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func getHue(cfg *Config) (Hue, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, hueEndPoint(cfg)+"/lights", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	hue := Hue{}

	if err := json.NewDecoder(resp.Body).Decode(&hue); err != nil {
		return nil, err
	}

	return hue, nil
}
