package balancer

import (
	"math/rand"
	"time"
)

/**
  @author: kafmio
  @since: 2023/3/9
  @desc: //TODO Implement the Random algorithm
**/

func init() {
	factories[RandomBalancer] = NewRandom
}

type Random struct {
	BaseBalancer
	rnd *rand.Rand
}

func NewRandom(hosts []string) Balancer {
	return &Random{
		BaseBalancer: BaseBalancer{
			hosts: hosts,
		},
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *Random) Balance(_ string) (string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.hosts) == 0 {
		return "", NoHostError
	}
	return r.hosts[r.rnd.Intn(len(r.hosts))], nil
}
