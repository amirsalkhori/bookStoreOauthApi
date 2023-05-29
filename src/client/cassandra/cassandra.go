package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

// var (
// 	cluster *gocql.ClusterConfig
// )

func GetSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	fmt.Println("Cassandra connection successfully...")

	// defer session.Close()

	return session, nil
}

// func GetSession() (*gocql.Session, error) {
// 	cluster, err := cluster.CreateSession()
// 	if err != nil {
// 		fmt.Println("My error is:", err.Error())
// 	}
// 	return cluster, nil
// }
