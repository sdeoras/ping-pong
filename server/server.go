package server

import (
	"context"
	"fmt"

	"github.com/sdeoras/ping-pong/pb"
)

type server struct{}

// Mesg receives input from client and sends pong back to the client with inverted counter
func (s *server) Mesg(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	fmt.Printf("server: %4d, %s\n", in.Packet.Counter, in.Packet.Mesg)

	var mesg string
	switch in.Packet.Mesg {
	case "ping":
		mesg = "pong"
	case "pong":
		mesg = "ping"
	default:
		mesg = "undefined"
	}

	return &pb.SendResponse{
		Packet: &pb.Packet{
			Mesg:    mesg,
			Counter: -in.Packet.Counter,
		},
	}, nil
}

func NewPingPongServer() pb.PingPongServer {
	return new(server)
}
