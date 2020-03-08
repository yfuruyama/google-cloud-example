package main

import (
	"context"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/yfuruyama/crzerolog"
	pb "github.com/yfuruyama/google-cloud-example/cloud_run/grpc_unary/proto"
)

type server struct {
}

func (s *server) Echo(ctx context.Context, r *pb.EchoRequest) (*pb.EchoReply, error) {
	logger := log.Ctx(ctx)

	logger.Info().Msg("gRPC Unary!")

	return &pb.EchoReply{Msg: r.GetMsg() + "!"}, nil
}

func main() {
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal().Msgf("Failed to listen: %v", err)
	}

	rootLogger := zerolog.New(os.Stdout)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(crzerolog.InjectLoggerInterceptor(&rootLogger)),
	)
	pb.RegisterHelloServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatal().Msgf("Failed to serve: %v", err)
	}
}
