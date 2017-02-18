package service

import (
	"context"
)

type HelloService interface {
	Hello(context.Context, string) (string, error)
	HelloAgain(context.Context, string) (string, error)
}

type helloService struct {
}

func NewHelloService() HelloService {
	s := &helloService{}
	return s
}

func (s *helloService) Hello(ctx context.Context, name string) (string, error) {
	return "hello " + name, nil
}

func (s *helloService) HelloAgain(ctx context.Context, name string) (string, error) {
	return "hello again " + name, nil
}
