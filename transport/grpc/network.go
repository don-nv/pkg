package grpc

type Network string

const (
	NetworkTCP Network = "tcp"
	NetworkUDP Network = "udp"
)

func (n Network) String() string {
	return string(n)
}
