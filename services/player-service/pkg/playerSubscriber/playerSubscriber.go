package playersubscriber

import (
	"context"

	"github.com/geopopos/simple_rpg/services/player-service/pkg/playerstore"
	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"
	"github.com/sirupsen/logrus"
)

// PlayerSubscriber defines a player subscriber
type PlayerSubscriber struct {
	Players playerstore.PlayerStore
	Logger  *logrus.Logger
}

// NewPlayerSubscriber will create a cache and return player subscriber
func NewPlayerSubscriber(logger *logrus.Logger) *PlayerSubscriber {
	memoryStore := playerstore.NewMemoryStore() // TODO this should be configurable
	return &PlayerSubscriber{memoryStore, logger}
}

// UpdatePosition will update the users position on the map in datastorage
func (p *PlayerSubscriber) UpdatePosition(ctx context.Context, event *proto.MovementEvent) error {
	newPlayerInfo := &proto.Player{
		GUID:    event.GUID,
		MapGUID: event.MapGUID,
		XPos:    event.XPos,
		YPos:    event.YPos,
	}
	if err := p.Players.UpdatePlayer(newPlayerInfo); err != nil {
		p.Logger.Errorf("err updating playerstore: %s", err)
		return err
	}
	p.Logger.Debugf("player GUID: %s moved at %d", event.GUID, event.Timestamp)
	return nil
}
