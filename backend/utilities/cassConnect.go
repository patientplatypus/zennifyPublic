package util


import(
	"github.com/gocql/gocql"
)


var CassCluster *gocql.ClusterConfig
var CassSession *gocql.Session