package balancer

import (
	fibHeap "github.com/starwander/GoFibonacciHeap"
	"sync"
)

/**
  @author: kafmio
  @since: 2023/3/9
  @desc: //TODO Implement the least_load algorithm
**/

func init() {
	factories[LeastLoadBalancer] = NewLeastLoad
}

// Tag .
func (h *host) Tag() interface{} { return h.name }

// Key .
func (h *host) Key() float64 { return float64(h.load) }

type LeastLoad struct {
	sync.RWMutex
	heap *fibHeap.FibHeap
}

// NewLeastLoad create new LeastLoad balancer
func NewLeastLoad(hosts []string) Balancer {
	ll := &LeastLoad{heap: fibHeap.NewFibHeap()}
	for _, h := range hosts {
		ll.AddHost(h)
	}
	return ll
}

// AddHost new host to the balancer
func (l *LeastLoad) AddHost(hostName string) {
	l.Lock()
	defer l.Unlock()
	if ok := l.heap.GetValue(hostName); ok != nil {
		return
	}
	_ = l.heap.InsertValue(&host{hostName, 0})
}

// RemoveHost new host from the balancer
func (l *LeastLoad) RemoveHost(hostName string) {
	l.Lock()
	defer l.Unlock()
	if ok := l.heap.GetValue(hostName); ok == nil {
		return
	}
	_ = l.heap.Delete(hostName)
}

func (l *LeastLoad) Balance(_ string) (string, error) {
	l.RLock()
	defer l.RUnlock()
	if l.heap.Num() == 0 {
		return "", NoHostError
	}
	return l.heap.MinimumValue().Tag().(string), nil
}

// IncConn refers to the number of connections to the server `+1`
func (l *LeastLoad) IncConn(hostName string) {
	l.Lock()
	defer l.Unlock()
	if ok := l.heap.GetValue(hostName); ok == nil {
		return
	}
	h := l.heap.GetValue(hostName)
	h.(*host).load++
	_ = l.heap.IncreaseKeyValue(h)
}

// DoneConn refers to the number of connections to the server `-1`
func (l *LeastLoad) DoneConn(hostName string) {
	l.Lock()
	defer l.Unlock()
	if ok := l.heap.GetValue(hostName); ok == nil {
		return
	}
	h := l.heap.GetValue(hostName)
	h.(*host).load--
	_ = l.heap.DecreaseKeyValue(h)
}
