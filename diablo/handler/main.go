package main

import (
	"context"
	"time"

	proto "github.com/gr4yha7/maikuro/diablo/proto"
	micro "github.com/micro/go-micro/v2"
	logger "github.com/micro/go-micro/v2/logger"
	"google.golang.org/grpc"
)

type Users struct{}

const (
	addr        = "127.0.0.1:9967"
	serviceName = "go.micro.service.users"
)

func (u *Users) Register(ctx context.Context, req *proto.RegisterRequest, resp *proto.RegisterResponse) error {
	logger.Info("Received Users.Register request")
	// TODO Fetch user input, persist to DB and return jwt token
	resp.Token = ""
	return nil
}

func (u *Users) Login(ctx context.Context, req *proto.LoginRequest, resp *proto.LoginResponse) error {
	logger.Info("Received Users.Login request")
	// TODO Fetch user input, persist to DB and return jwt token
	resp.Token = ""
	return nil
}

func main() {
	go func() {
		for {
			grpc.DialContext(context.TODO(), addr)
			time.Sleep(time.Second)
		}
	}()

	service := micro.NewService(
		micro.Name(serviceName),
	)

	service.Init()

	proto.RegisterUsersHandler(service.Server(), new(Users))

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
