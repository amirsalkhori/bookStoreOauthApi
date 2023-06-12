package cassandra

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Timeout = 2 * time.Second // Set the timeout to 2 seconds
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}

	fmt.Println("Cassandra connection successfully...")
}

func GetSession() *gocql.Session {
	return session
}
