package balancer

import "sync"

/**
  @author: kafmio
  @since: 2023/3/9
  @desc: //TODO Base balancer for balance algorithms
**/

type BaseBalancer struct {
	sync.RWMutex
	hosts []string
}

// AddHost new host to the balancer
func (b *BaseBalancer) AddHost(host string) {
	b.Lock()
	defer b.Unlock()
	for _, h := range b.hosts {
		if h == host {
			return
		}
	}
	b.hosts = append(b.hosts, host)
}

// RemoveHost new host from the balancer
func (b *BaseBalancer) RemoveHost(host string) {
	b.Lock()
	defer b.Unlock()
	for i, h := range b.hosts {
		if h == host {
			b.hosts = append(b.hosts[:i], b.hosts[i+1:]...)
			return
		}
	}
}

// Balance selects a suitable host according
func (b *BaseBalancer) Balance(key string) (string, error) {
	return "", nil
}

// IncConn .
func (b *BaseBalancer) IncConn(_ string) {}

// DoneConn .
func (b *BaseBalancer) DoneConn(_ string) {}
