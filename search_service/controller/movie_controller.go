package controller

import (
	"context"

	proto "github.com/nipeharefa/silver-parakeet/model"
	"github.com/nipeharefa/silver-parakeet/search_service/usecase"
)

type (
	movieController struct {
		movieUsecase usecase.MovieUsecase
	}
)

func NewMovieController(movieUsecase usecase.MovieUsecase) proto.MovieServiceServer {
	return &movieController{movieUsecase}
}

func (m *movieController) GetMovies(ctx context.Context, param *proto.GetMoviesParam) (*proto.Movies, error) {

	results := m.movieUsecase.GetMoviesFromAPI(param.SearchWord, int(param.Page))

	moviesList := make([]*proto.Movie, 0)

	for _, mv := range results.Search {
		moviesList = append(moviesList, &proto.Movie{
			Title:  mv.Title,
			Year:   mv.Year,
			ImdbID: mv.ImdbID,
			Type:   mv.Type,
			Poster: mv.Poster,
		})
	}

	movies := &proto.Movies{}
	movies.List = moviesList
	return movies, nil
}
