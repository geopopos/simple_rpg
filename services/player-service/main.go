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

type PlayerStore interface {
	GetPlayer(GUID string) (*proto.Player, error)
}

type MemoryStore struct {
	Store []*proto.Player
}

func NewMemoryStore() *MemoryStore {
	player := &proto.Player{
		GUID:      "turdmongler",
		Username:  "tmongler",
		MapGUID:   "some-guid",
		XPos:      5,
		YPos:      7,
		Health:    10,
		MaxHealth: 100,
	}
	playerStore := []*proto.Player{player}
	return &MemoryStore{playerStore}
}

func (m *MemoryStore) GetPlayer(GUID string) (*proto.Player, error) {
	for _, user := range m.Store {
		if user.GUID == GUID {
			return user, nil
		}
	}
	return nil, fmt.Errorf("player not found")
}

// PlayerService defines a player service
type PlayerService struct {
	Players PlayerStore
}

// NewPlayerService will create a cache and return player service
func NewPlayerService() *PlayerService {
	memoryStore := NewMemoryStore()
	return &PlayerService{memoryStore}
}

// Hello recieves a request and returns response
func (p *PlayerService) GetPlayer(ctx context.Context, req *proto.PlayerRequest, rsp *proto.PlayerResponse) (err error) {
	rsp.Player, err = p.Players.GetPlayer(req.PlayerGUID)
	if err != nil && err.Error() == "player not found" {
		rsp.Player = nil
		return nil
	}
	return
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

	proto.RegisterPlayerServiceHandler(playerService.Server(), NewPlayerService())
	if err := playerService.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
