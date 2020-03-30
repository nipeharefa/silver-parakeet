package controller

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/nipeharefa/silver-parakeet/model"
	"github.com/stretchr/testify/assert"
)

func TestMovieController(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("TestGetMovies", func(t *testing.T) {
		a := model.NewMockMovieServiceClient(ctrl)

		movies := &model.Movies{}
		in := &model.GetMoviesParam{SearchWord: "", Page: 0}
		a.EXPECT().
			GetMovies(
				context.Background(),
				in).
			Return(movies, nil).
			AnyTimes()

		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := NewMovieController(a)

		if assert.NoError(t, handler.Search(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("TestErrGetMovies", func(t *testing.T) {
		a := model.NewMockMovieServiceClient(ctrl)
		movieJSON := "[]\n"

		in := &model.GetMoviesParam{SearchWord: "", Page: 0}
		a.EXPECT().
			GetMovies(context.Background(), in).
			Return(nil, errors.New("s")).
			AnyTimes()

		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := NewMovieController(a)

		if assert.NoError(t, handler.Search(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, movieJSON, rec.Body.String())
		}
	})
}
