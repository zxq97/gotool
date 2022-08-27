package generate

import (
	"fmt"
	"hash/crc32"
	"os"

	"github.com/bwmarrin/snowflake"
)

type SnowIDGen struct {
	node *snowflake.Node
}

func NewSnowIDGen(svc string) (*SnowIDGen, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	svc = fmt.Sprintf("%s-%s-%d", svc, hostname, os.Getpid())
	node, err := snowflake.NewNode(int64(crc32.ChecksumIEEE([]byte(svc))) % 1024)
	if err != nil {
		return nil, err
	}
	return &SnowIDGen{node: node}, nil
}

func (g *SnowIDGen) Gen() int64 {
	return g.node.Generate().Int64()
}
