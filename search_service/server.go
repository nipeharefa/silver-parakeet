package search

import (
	"net"

	proto "github.com/nipeharefa/silver-parakeet/model"
	"github.com/nipeharefa/silver-parakeet/search_service/controller"
	"github.com/nipeharefa/silver-parakeet/search_service/usecase"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type (
	searchService struct {
		grpcServer *grpc.Server
	}
)

func New() *searchService {

	ss := &searchService{}
	ss.grpcServer = grpc.NewServer()
	return ss
}

func (s *searchService) registerGRPCService() {

	// Usecase layer
	mu := usecase.NewMovieUsecase(viper.GetString("apiKey"))

	// Delivery / controller layer
	cc := controller.NewMovieController(mu)

	proto.RegisterMovieServiceServer(s.grpcServer, cc)
}

func (s *searchService) Run() error {

	s.registerGRPCService()

	address := ":6000"
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Infof("Starting GRPC server on Port %s", address)

	return s.grpcServer.Serve(l)
}
