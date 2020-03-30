package gateway

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/nipeharefa/silver-parakeet/api_gateway/controller"
	proto "github.com/nipeharefa/silver-parakeet/model"
	"google.golang.org/grpc"
)

type (
	APIGateway interface {
		Run() error
	}

	apiGateway struct {
		e *echo.Echo
	}
)

func NewAPIGateway() APIGateway {

	e := echo.New()
	return &apiGateway{e}
}

func (a *apiGateway) Run() error {

	conn, err := grpc.Dial(":6000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	movieClient := proto.NewMovieServiceClient(conn)

	mc := controller.NewMovieController(movieClient)

	a.e.GET("/search", mc.Search)

	return a.e.Start(":8000")
}
