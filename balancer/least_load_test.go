package balancer

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestLeastLoad_Balance .
func TestLeastLoad_Balance(t *testing.T) {
	expect, err := Build(LeastLoadBalancer, []string{"192.168.1.1:1015",
		"192.168.1.1:1016", "192.168.1.1:1017", "192.168.1.1:1018"})
	expect.RemoveHost("192.168.1.1:1018")
	assert.Equal(t, err, nil)
	expect.IncConn("192.168.1.1:1015")
	expect.IncConn("192.168.1.1:1016")
	expect.IncConn("192.168.1.1:1016")
	expect.IncConn("192.168.1.1:1018")
	expect.DoneConn("192.168.1.1:1018")
	expect.DoneConn("192.168.1.1:1016")
	ll := NewLeastLoad([]string{"192.168.1.1:1016"})
	ll.RemoveHost("192.168.1.1:1018")
	ll.AddHost("192.168.1.1:1015")
	ll.AddHost("192.168.1.1:1016")
	ll.AddHost("192.168.1.1:1017")
	ll.IncConn("192.168.1.1:1015")
	ll.IncConn("192.168.1.1:1016")
	ll.IncConn("192.168.1.1:1016")
	ll.DoneConn("192.168.1.1:1016")
	llHost, _ := ll.Balance("")
	expectHost, _ := expect.Balance("")
	assert.Equal(t, true, reflect.DeepEqual(llHost, expectHost))
}
