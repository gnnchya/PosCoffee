package util

import (
	"github.com/bwmarrin/snowflake"
)

//go:generate mockery --name=UUID
type UUID interface {
	Generate() (uuid string)
}

type SnowflakeUUID struct {
	generator *snowflake.Node
}

func NewUUID() (uuid UUID, err error) {
	node, err := snowflake.NewNode(1023)
	if err != nil {
		return nil, err
	}
	return &SnowflakeUUID{generator: node}, nil
}

func (u *SnowflakeUUID) Generate() (uuid string) {
	return u.generator.Generate().String()
}
