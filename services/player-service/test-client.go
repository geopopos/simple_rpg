package main

import (
	"context"
	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"
	"github.com/micro/go-grpc"
	"log"
)

func main() {
	service := grpc.NewService()
	service.Init()

	cli := proto.NewPlayerService("go.micro.srv.user", service.Client())
	// names := []string{"Geopopos", "Sleyva", "turdmongler"}
	requests := []*proto.Request{
		&proto.Request{Name: "Geopopos"},
		&proto.Request{Name: "Sleyva"},
		&proto.Request{Name: "turdmongler"},
	}
	for _, request := range requests {
		rsp, err := cli.Hello(context.TODO(), request)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(rsp)
	}
}
