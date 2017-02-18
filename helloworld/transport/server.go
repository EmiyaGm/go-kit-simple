package transport

import (
	"context"
	pb "github.com/anarcher/go-kit-simple/helloworld/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	oldcontext "golang.org/x/net/context"
)

// Server is implemented by pb.greeterServer
type Server struct {
	handlers Handlers
}

func MakeServer(ctx context.Context, handlers Handlers) *Server {
	s := &Server{
		handlers: handlers,
	}
	return s
}

func (s Server) SayHello(ctx oldcontext.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	res, err := s.handle(ctx, s.handlers.SayHello, req)
	return res.(*pb.HelloReply), err
}

func (s Server) SayHelloAgain(ctx oldcontext.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	res, err := s.handle(ctx, s.handlers.SayHelloAgain, req)
	return res.(*pb.HelloReply), err
}

func (s Server) handle(ctx oldcontext.Context, handler grpctransport.Handler, req interface{}) (interface{}, error) {
	_, rep, err := handler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep, nil
}
