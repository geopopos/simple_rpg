package main

import (
	"github.com/geopopos/simple_rpg/services/player-service/pkg/playerservice"
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
	golog.Formatter = logging.Formatter
	logger := micrologrus.NewMicroLogrus(golog)
	log.SetLogger(logger)

	playerService := micro.NewService(
		micro.Name(playerService),
		micro.WrapHandler(logging.LogWrapper),
	)
	playerService.Init()

	proto.RegisterPlayerServiceHandler(playerService.Server(), playerservice.NewPlayerService())
	if err := playerService.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
