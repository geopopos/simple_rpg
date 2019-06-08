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
	requests := []*proto.PlayerRequest{
		&proto.PlayerRequest{PlayerGUID: "Geopopos"},
		&proto.PlayerRequest{PlayerGUID: "Sleyva"},
		&proto.PlayerRequest{PlayerGUID: "turdmongler"},
	}
	for _, request := range requests {
		rsp, err := cli.GetPlayer(context.TODO(), request)
		if err != nil {
			log.Fatal(err)
		}
		if rsp.Player == nil {
		   log.Print("No player exists with this GUID");
		} else {
		  log.Print(rsp)
		}
	}
}
