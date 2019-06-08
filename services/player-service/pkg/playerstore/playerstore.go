package playerstore

import (
	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"
)

// PlayerStore in the interface for our interactions
// with the underlying storage for our players
type PlayerStore interface {
	GetPlayer(GUID string) (*proto.Player, error)
}
