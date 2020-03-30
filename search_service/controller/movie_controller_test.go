package controller

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nipeharefa/silver-parakeet/model"
	"github.com/nipeharefa/silver-parakeet/search_service/usecase"
)

func TestMovieController(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	a := usecase.NewMockMovieUsecase(ctrl)

	movieResponse := usecase.MovieResponse{}

	mc := NewMovieController(a)

	a.EXPECT().
		GetMoviesFromAPI("batman", 1).
		Return(movieResponse).AnyTimes()

	param := &model.GetMoviesParam{SearchWord: "batman", Page: 1}
	mc.GetMovies(context.Background(), param)
}
