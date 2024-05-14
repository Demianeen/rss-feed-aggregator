package utils

import (
	"io"
	"net/http"
	"time"
)

func FetchDataFromUrl(url string) ([]byte, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	return data, nil
}
