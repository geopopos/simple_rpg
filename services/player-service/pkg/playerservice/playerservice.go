package playerservice

import (
	"context"

	"github.com/geopopos/simple_rpg/services/player-service/pkg/config"
	"github.com/geopopos/simple_rpg/services/player-service/pkg/playerstore"
	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"

	"github.com/sirupsen/logrus"
)

// PlayerService defines a player service
type PlayerService struct {
	Players playerstore.PlayerStore
	Logger  *logrus.Logger
}

// NewPlayerService will create a cache and return player service
func NewPlayerService(logger *logrus.Logger, conf *config.Configuration) *PlayerService {
	var store playerstore.PlayerStore
	switch conf.Database.Type {
	case config.DBTypeMemory:
		store = playerstore.NewMemoryStore()
	case config.DBTypeCassandra:
		cass, err := playerstore.NewCassandraStore(conf)
		if err != nil {
			logger.Fatalf("err initilizing store: %s", err)
		}
		store = cass
	default:
		logger.Fatalf("%s not supported", conf.Database.Type)
	}
	logger.Infof("Database %s Initialized Successfully for service", conf.Database.Type)

	return &PlayerService{store, logger}
}

// GetPlayer recieves request for player information and reads from datastore and returns response
func (p *PlayerService) GetPlayer(ctx context.Context, req *proto.PlayerRequest, rsp *proto.PlayerResponse) (err error) {
	rsp.Player, err = p.Players.GetPlayer(req.PlayerGUID)
	if err != nil && err.Error() == "player not found" {
		rsp.Player = nil
		return nil
	}
	return
}
