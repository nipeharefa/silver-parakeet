package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	proto "github.com/nipeharefa/silver-parakeet/model"
)

// controller will connecto to grpc client

type (
	MovieController interface {
		Search(c echo.Context) error
	}

	movieController struct {
		movieClient proto.MovieServiceClient
	}
)

func NewMovieController(movieClient proto.MovieServiceClient) MovieController {

	return &movieController{movieClient}
}

func (m *movieController) Search(c echo.Context) error {

	var s string
	ctx := context.Background()

	page, _ := strconv.Atoi(c.QueryParam("pagination"))
	s = c.QueryParam("searchword")
	in := &proto.GetMoviesParam{
		Page:       int32(page),
		SearchWord: s,
	}

	movies, _ := m.movieClient.GetMovies(ctx, in)
	return c.JSON(http.StatusOK, movies.List)
}
