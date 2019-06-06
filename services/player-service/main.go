package main

import (
	"context"
	"fmt"
	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"
	log "github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"github.com/srleyva/turbine/user-service/pkg/logging"
	micrologrus "github.com/tudurom/micro-logrus"
)

// CLI args
var playerService = "go.micro.srv.user"

// PlayerService defines a player service
type PlayerService struct {
	playerStore []*proto.Request
}

// NewPlayerService will create a cache and return player service
func NewPlayerService() *PlayerService {
	playerStore := []*proto.Request{}
	return &PlayerService{playerStore}
}

// Hello recieves a request and returns response
func (p *PlayerService) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	p.playerStore = append(p.playerStore, req)
	rsp.Msg = fmt.Sprintf("New List is %v", p.playerStore)
	return nil
}

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

	proto.RegisterPlayerHandler(playerService.Server(), NewPlayerService())
	if err := playerService.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
