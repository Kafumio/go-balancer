package balancer

import "errors"

/**
  @author: kaf
  @since: 2023/3/9
  @desc: //TODO balancer interface statement
**/

var (
	NoHostError               = errors.New("no host")
	AlgorithmNoSupportedError = errors.New("algorithm not supported")
)

// Balancer interface is the load balancer for the reverse proxy
type Balancer interface {
	AddHost(string)
	RemoveHost(string)
	Balance(string) (string, error)
	IncConn(string)
	DoneConn(string)
}

// Factory is the factory that generates Balancer,
// and the factory design pattern is used here
type Factory func([]string) Balancer

var factories = make(map[string]Factory)

// Build generates the corresponding Balancer according to the algorithm
func Build(algorithm string, hosts []string) (Balancer, error) {
	factory, ok := factories[algorithm]
	if !ok {
		return nil, AlgorithmNoSupportedError
	}
	return factory(hosts), nil
}
