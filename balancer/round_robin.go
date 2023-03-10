package balancer

/**
  @author: kafmio
  @since: 2023/3/9
  @desc: //TODO Implement the Round-Robin algorithm
**/

func init() {
	factories[RoundRobinBalancer] = NewRoundRobin
}

type RoundRobin struct {
	BaseBalancer
	i uint64
}

// NewRoundRobin create new RoundRobin balancer
func NewRoundRobin(hosts []string) Balancer {
	return &RoundRobin{
		i: 0,
		BaseBalancer: BaseBalancer{
			hosts: hosts,
		},
	}
}

// Balance selects a suitable host according
func (r *RoundRobin) Balance(_ string) (string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.hosts) == 0 {
		return "", NoHostError
	}
	host := r.hosts[r.i%uint64(len(r.hosts))]
	r.i++
	return host, nil
}
