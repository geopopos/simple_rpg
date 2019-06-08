package playerstore

import (
	"fmt"

	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"
)

// MemoryStore is the store type that stores everything
// in Memory. Due to its ephemeral nature it should not
// be used for production but rather just testing
type MemoryStore struct {
	Store []*proto.Player
}

// NewMemoryStore returns initializes an array with one player
// and creates a new MemoryStore with it
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

// GetPlayer returns a player given the GUID
func (m *MemoryStore) GetPlayer(GUID string) (*proto.Player, error) {
	for _, user := range m.Store {
		if user.GUID == GUID {
			return user, nil
		}
	}
	return nil, fmt.Errorf("player not found")
}
