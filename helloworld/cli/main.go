package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"os"

	"github.com/anarcher/go-kit-simple/helloworld/client"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-kit-getstart"
	app.Usage = "go-kit getstart with gRPC"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "name",
			EnvVar: "NAME",
			Value:  "",
		},
		cli.StringFlag{
			Name:   "addr",
			EnvVar: "ADDR",
			Value:  ":8081",
		},
	}

	app.Action = func(c *cli.Context) error {
		ctx := context.Background()

		conn, err := grpc.Dial(c.String("addr"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		service := client.New(conn)
		reply, err := service.Hello(ctx, c.String("name"))
		if err != nil {
			return err
		}

		fmt.Println(reply)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}

}
