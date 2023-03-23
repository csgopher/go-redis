package cluster

import "go-redis/interface/resp"

func execSelect(cluster *clusterDatabase, c resp.Connection, cmdAndArgs [][]byte) resp.Reply {
	return cluster.db.Exec(c, cmdAndArgs)
}
