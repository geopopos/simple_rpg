package main

import (
	"github.com/geopopos/simple_rpg/services/player-service/pkg/playerservice"
	"github.com/geopopos/simple_rpg/services/player-service/pkg/playersubscriber"
	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"
	log "github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"github.com/srleyva/turbine/user-service/pkg/logging"
	micrologrus "github.com/tudurom/micro-logrus"
)

// CLI args
var playerService = "go.micro.srv.user"

func main() {
	golog := logrus.New()
	golog.SetLevel(logrus.DebugLevel)
	golog.Formatter = logging.Formatter
	logger := micrologrus.NewMicroLogrus(golog)
	log.SetLogger(logger)

	playerService := micro.NewService(
		micro.Name(playerService),
		micro.WrapHandler(logging.LogWrapper),
	)
	playerService.Init()

	service := playerservice.NewPlayerService(golog)
	subscriber := playersubscriber.NewPlayerSubscriber(golog)

	if err := proto.RegisterPlayerServiceHandler(playerService.Server(), service); err != nil {
		log.Fatalf("err registering servie: %s", err)
	}

	if err := micro.RegisterSubscriber("rpg.player-movement", playerService.Server(), subscriber); err != nil {
		log.Fatalf("err registering subscriber: %s", err)
	}

	if err := playerService.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
