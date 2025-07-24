package clients

import (
	"github.com/GeorgeTyupin/MLCinema/go_server/internal/models"
	"github.com/go-resty/resty/v2"
)

type MLClient struct {
	client  *resty.Client
	baseURL string
}

func NewMLClient(baseURL string) *MLClient {
	return &MLClient{
		client:  resty.New(),
		baseURL: baseURL,
	}
}

func (ml *MLClient) SearchMovies(query string) ([]models.Film, error) {
	var response struct {
		Status  string        `json:"status"`
		Query   string        `json:"query"`
		Found   int           `json:"found"`
		Results []models.Film `json:"results"`
	}

	_, err := ml.client.R().
		SetFormData(map[string]string{"query": query}).
		SetResult(&response).
		Post(ml.baseURL + "/recommend")

	if err != nil {
		return nil, err
	}

	return response.Results, nil
}
