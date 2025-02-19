package adapters

import (
	"io"
	"net/http"
	"time"

	"github.com/WarisLi/golang-backend-challenge-3/core"
)

type APIClient struct {
	client *http.Client
}

func NewAPIClient() core.BeefRepository {
	return &APIClient{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (a *APIClient) GetData() ([]byte, error) {
	resp, err := a.client.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
