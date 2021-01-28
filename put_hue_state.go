package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	StatusSuccess ResponseStatus = "success"
	StatusError   ResponseStatus = "error"
)

type Response []ResponseAttribute

func (r Response) Err() error {
	for _, attr := range r {
		for status, detail := range attr {
			switch status {
			case StatusSuccess:
				continue
			case StatusError:
				return ErrPutHueState(detail["description"])
			}
		}
	}

	return nil
}

type ResponseAttribute map[ResponseStatus]ResponseDetail

type ResponseStatus string

type ResponseDetail map[string]json.RawMessage

func putHueState(cfg *Config, lightID HueLightID, state *HueState) error {
	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(state); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, hueEndPoint(cfg)+fmt.Sprintf("/lights/%s/state", lightID), body)
	if err != nil {
		return err
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	respBody := Response{}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return err
	}

	if err := respBody.Err(); err != nil {
		return err
	}

	return nil
}
