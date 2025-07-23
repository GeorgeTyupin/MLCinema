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
	var films []models.Film

	_, err := ml.client.R().
		SetFormData(map[string]string{"query": query}).
		SetResult(&films).
		Post(ml.baseURL + "/recommend")

	return films, err
}
