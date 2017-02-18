package client

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	pb "github.com/anarcher/go-kit-simple/helloworld/pb"
	"github.com/anarcher/go-kit-simple/helloworld/service"
)

type client struct {
	hello      endpoint.Endpoint
	helloAgain endpoint.Endpoint
}

func New(conn *grpc.ClientConn) service.HelloService {

	hello := grpctransport.NewClient(
		conn,
		"Greeter",
		"SayHello",
		encodeRequest,
		decodeResponse,
		pb.HelloReply{},
	)
	helloAgain := grpctransport.NewClient(
		conn,
		"Greeter",
		"SayHello",
		encodeRequest,
		decodeResponse,
		pb.HelloReply{},
	)

	c := &client{
		hello:      hello.Endpoint(),
		helloAgain: helloAgain.Endpoint(),
	}
	return c
}

func (c client) Hello(ctx context.Context, name string) (string, error) {
	req := &pb.HelloRequest{Name: name}
	res, err := c.hello(ctx, req)
	if err != nil {
		return "", err
	}

	return res.(*pb.HelloReply).Message, nil

}

func (c client) HelloAgain(ctx context.Context, name string) (string, error) {
	req := &pb.HelloRequest{Name: name}
	res, err := c.helloAgain(ctx, req)
	if err != nil {
		return "", err
	}

	return res.(*pb.HelloReply).Message, nil
}

func encodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func decodeResponse(_ context.Context, res interface{}) (interface{}, error) {
	return res, nil
}
