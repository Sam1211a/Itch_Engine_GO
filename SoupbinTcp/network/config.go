package network

import (
	"context"
	"net"
	"soupbintcp/engine"
)

type Client struct {
	Conn     net.Conn
	Eng      *engine.Engine
	Ctx      context.Context
	Cancel   context.CancelFunc
	Sequence uint64
	Session  string
}

const (
	Host     = "10.61.90.20:54001"
	Username = "ITFIC1"
	Password = "password"

	Session = ""
	// Sequence = "1"
)
