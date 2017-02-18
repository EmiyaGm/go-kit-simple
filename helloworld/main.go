package main

import (
	"context"
	"net"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/urfave/cli"
	"google.golang.org/grpc"

	"github.com/anarcher/go-kit-simple/helloworld/endpoint"
	"github.com/anarcher/go-kit-simple/helloworld/pb"
	"github.com/anarcher/go-kit-simple/helloworld/service"
	"github.com/anarcher/go-kit-simple/helloworld/transport"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-kit-getstart"
	app.Usage = "go-kit getstart with gRPC"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "addr",
			EnvVar: "ADDR",
			Value:  ":8081",
		},
	}

	app.Action = func(c *cli.Context) error {
		var logger log.Logger
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)

		logger.Log("start", 1)

		ctx := context.Background()

		//In server.go or manager.go
		ln, err := net.Listen("tcp", c.String("addr"))
		if err != nil {
			return err
		}

		helloService := service.NewHelloService()
		endpoints := endpoint.MakeEndpoints(helloService)
		handlers := transport.MakeHandlers(ctx, endpoints)

		srv := transport.MakeServer(ctx, handlers)

		s := grpc.NewServer()
		pb.RegisterGreeterServer(s, srv)
		return s.Serve(ln)

	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}

}
