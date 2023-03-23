package cluster

import "go-redis/interface/resp"

func ping(cluster *clusterDatabase, c resp.Connection, cmdAndArgs [][]byte) resp.Reply {
	return cluster.db.Exec(c, cmdAndArgs)
}
