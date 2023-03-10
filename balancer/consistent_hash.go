package balancer

import "github.com/lafikl/consistent"

/**
  @author: kafmio
  @since: 2023/3/9
  @desc: //TODO Implement the consistent_hash algorithm
**/

func init() {
	factories[ConsistentHashBalancer] = NewConsistent
}

// Consistent refers to consistent hash
type Consistent struct {
	BaseBalancer
	ch *consistent.Consistent
}

// NewConsistent create new Consistent balancer
func NewConsistent(hosts []string) Balancer {
	c := &Consistent{
		ch: consistent.New(),
	}
	for _, h := range hosts {
		c.ch.Add(h)
	}
	return c
}

// AddHost new host to the balancer
func (c *Consistent) AddHost(host string) {
	c.ch.Add(host)
}

// RemoveHost new host from the balancer
func (c *Consistent) RemoveHost(host string) {
	c.ch.Remove(host)
}

// Balance selects a suitable host according to the key value
func (c *Consistent) Balance(key string) (string, error) {
	if len(c.ch.Hosts()) == 0 {
		return "", NoHostError
	}
	return c.ch.Get(key)
}
