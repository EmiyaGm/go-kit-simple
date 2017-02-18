package endpoint

import (
	"context"

	pb "github.com/anarcher/go-kit-simple/helloworld/pb"
	"github.com/anarcher/go-kit-simple/helloworld/service"
)

type Endpoints struct {
	service service.HelloService
}

func MakeEndpoints(svc service.HelloService) Endpoints {
	es := Endpoints{
		service: svc,
	}

	return es
}

func (es Endpoints) SayHello(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(*pb.HelloRequest)

	ret, err := es.service.Hello(ctx, request.Name)
	if err != nil {
		return nil, err
	}

	return &pb.HelloReply{
		Message: ret,
	}, nil
}

func (es Endpoints) SayHelloAgain(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(*pb.HelloRequest)

	ret, err := es.service.HelloAgain(ctx, request.Name)
	if err != nil {
		return nil, err
	}

	response := &pb.HelloReply{
		Message: ret,
	}
	return response, nil
}
