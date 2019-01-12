package rip

import (
	packet "github.com/coreswitch/zebra/pkg/packet/rip"
)

// Server RIP server structure.
type Server struct {
	RTE packet.RTE
}

// NewServer Create a new RIP server.
func NewServer() *Server {
	return &Server{}
}
