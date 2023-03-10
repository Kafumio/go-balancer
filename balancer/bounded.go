package balancer

import "github.com/lafikl/consistent"

/**
  @author: kafmio
  @since: 2023/3/9
  @desc: //TODO Implement Consistent Hashing with Bounded Loads algorithm
**/

func init() {
	factories[BoundedBalancer] = NewBounded
}

// Bounded refers to consistent hash with bounded
type Bounded struct {
	ch *consistent.Consistent
}

// NewBounded create new Bounded balancer
func NewBounded(hosts []string) Balancer {
	c := &Bounded{consistent.New()}
	for _, h := range hosts {
		c.ch.Add(h)
	}
	return c
}

// AddHost new host to the balancer
func (b *Bounded) AddHost(host string) {
	b.ch.Add(host)
}

// RemoveHost new host from the balancer
func (b *Bounded) RemoveHost(host string) {
	b.ch.Remove(host)
}

// Balance selects a suitable host according to the key value
func (b *Bounded) Balance(key string) (string, error) {
	if len(b.ch.Hosts()) == 0 {
		return "", NoHostError
	}
	return b.ch.GetLeast(key)
}

// IncConn refers to the number of connections to the server `+1`
func (b *Bounded) IncConn(host string) {
	b.ch.Inc(host)
}

// DoneConn refers to the number of connections to the server `-1`
func (b *Bounded) DoneConn(host string) {
	b.ch.Done(host)
}
