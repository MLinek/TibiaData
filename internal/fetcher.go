package internal

import (
	"io"
	"net/http"
)

type Fetcher struct {
}

func (f Fetcher) fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte("{}"), err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte("{}"), err
	}
	return body, nil
}
