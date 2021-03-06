package usecase

import (
	"net/url"
	"strconv"

	"github.com/nipeharefa/silver-parakeet/search_service/utils"
)

type (
	Movie struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	}

	MovieResponse struct {
		Search       []Movie `json:"Search"`
		TotalRequest string  `json:"totalResults"`
		Response     string  `json:"Response"`
	}

	MovieUsecase interface {
		GetMoviesFromAPI(string, int) MovieResponse
	}

	movieUsecase struct {
		httpClient utils.HTTPClient
		apiKey     string
	}
)

func NewMovieUsecase(apiKey string) MovieUsecase {

	httpClient := utils.NewHTTPClient()
	return &movieUsecase{httpClient, apiKey}
}

func (m *movieUsecase) GetMoviesFromAPI(s string, page int) MovieResponse {

	if page < 1 {
		page = 1
	}

	pagination := strconv.Itoa(int(page))

	u, _ := url.Parse("http://www.omdbapi.com")

	q := u.Query()
	q.Set("apiKey", m.apiKey)
	q.Set("s", s)
	q.Set("page", pagination)

	u.RawQuery = q.Encode()

	result := MovieResponse{}

	r := utils.RequestOptions{
		FullPath: u.String(),
	}

	_ = m.httpClient.Get(r, &result)

	return result
}
