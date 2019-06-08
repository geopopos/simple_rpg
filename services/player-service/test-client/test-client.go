package main

import (
	"time"

	"context"

	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

// send events using the publisher
func sendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)
	i := 0
	for _ = range t.C {
		// create new event
		ev := &proto.MovementEvent{
			GUID:      "turdmongler",
			Timestamp: time.Now().Unix(),
			XPos:      int64(i),
			YPos:      int64(i),
			MapGUID:   "some-guid",
		}

		log.Logf("publishing %+v\n", ev)

		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing: %v", err)
		}
		i++
	}
}

func main() {
	// create a service
	service := micro.NewService(
		micro.Name("go.micro.cli.pubsub"),
	)
	// parse command line
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
			log.Logf("No player exists with this GUID")
		} else {
			log.Logf("Found player: %s", rsp.Player.Username)
		}
	}

	// create publisher
	pub1 := micro.NewPublisher("rpg.player-movement", service.Client())

	// pub to topic 1
	sendEv("example.topic.pubsub.1", pub1)

}
