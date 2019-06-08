package playerstore

import (
	"fmt"

	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"
	"github.com/gocql/gocql"
)

// CassandraStore is the store type that stores everything
// in Cassandra Backend. Due to the fact that this won't be
// queried in realtime...cassandra is sufficient
type CassandraStore struct {
	Cassandra *gocql.ClusterConfig
}

// NewCassandraStore creates new cassandra store
func NewCassandraStore() (*CassandraStore, error) {
	cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster.Keyspace = "players"
	cluster.Consistency = gocql.Quorum
	if _, err := cluster.CreateSession(); err != nil {
		return nil, err
	}
	return &CassandraStore{cluster}, nil
}

// GetPlayer returns a player given the GUID
func (m *CassandraStore) GetPlayer(GUID string) (*proto.Player, error) {
	// TODO
	return nil, fmt.Errorf("player not found")
}

// UpdatePlayer updates player information
func (m *CassandraStore) UpdatePlayer(player *proto.Player) error {
	// TODO
	return nil
}
