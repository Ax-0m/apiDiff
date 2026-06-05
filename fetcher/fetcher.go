package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Fetch(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not fetch url %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code %d for url %s", resp.StatusCode, url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	var result interface{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, fmt.Errorf("could not parse response JSON: %w", err)
	}

	return result, nil
}
