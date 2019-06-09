package playerstore

import (
	"github.com/geopopos/simple_rpg/services/player-service/pkg/config"
	proto "github.com/geopopos/simple_rpg/services/player-service/proto/player"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
)

// CassandraStore is the store type that stores everything
// in Cassandra Backend. Due to the fact that this won't be
// queried in realtime...cassandra is sufficient
type CassandraStore struct {
	Cassandra *gocql.ClusterConfig
}

// NewCassandraStore creates new cassandra store
func NewCassandraStore(conf *config.Configuration) (*CassandraStore, error) {
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
	var player *proto.Player
	session, err := m.Cassandra.CreateSession()
	if err != nil {
		return nil, err
	}
	stmt, names := qb.Select("players.player").Where(qb.Eq("guid")).ToCql()
	q := gocqlx.Query(session.Query(stmt), names).BindMap(qb.M{
		"guid": GUID,
	})

	if err := q.GetRelease(&player); err != nil {
		return nil, err
	}
	return player, nil
}

// UpdatePlayer updates player information
func (m *CassandraStore) UpdatePlayer(player *proto.Player) error {
	return nil
}
