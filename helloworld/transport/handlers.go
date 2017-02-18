package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	ep "github.com/anarcher/go-kit-simple/helloworld/endpoint"
)

type Handlers struct {
	SayHello      grpctransport.Handler
	SayHelloAgain grpctransport.Handler
}

func MakeHandlers(ctx context.Context, endpoints ep.Endpoints) Handlers {
	hs := Handlers{
		SayHello:      MakeHandler(ctx, endpoints.SayHello),
		SayHelloAgain: MakeHandler(ctx, endpoints.SayHelloAgain),
	}

	return hs
}

func MakeHandler(ctx context.Context, e endpoint.Endpoint) grpctransport.Handler {
	s := grpctransport.NewServer(
		ctx,
		e,
		decodeRequestFunc,
		encodeResponseFunc,
	)

	return s
}

func decodeRequestFunc(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponseFunc(_ context.Context, res interface{}) (interface{}, error) {
	return res, nil
}
