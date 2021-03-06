package playerservice

import (
	"context"

	"github.com/geopopos/simple_rpg/services/player-service/pkg/playerstore"
	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"
)

// PlayerService defines a player service
type PlayerService struct {
	Players playerstore.PlayerStore
}

// NewPlayerService will create a cache and return player service
func NewPlayerService() *PlayerService {
	memoryStore := playerstore.NewMemoryStore() // TODO this should be configurable
	return &PlayerService{memoryStore}
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
