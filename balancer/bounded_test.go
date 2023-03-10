package balancer

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestBounded_AddHost .
func TestBounded_AddHost(t *testing.T) {
	expect, err := Build(BoundedBalancer, []string{"192.168.1.1:1015",
		"192.168.1.1:1016", "192.168.1.1:1017", "192.168.1.1:1018"})
	assert.Equal(t, err, nil)
	bounded := NewBounded(nil)
	bounded.AddHost("192.168.1.1:1015")
	bounded.AddHost("192.168.1.1:1016")
	bounded.AddHost("192.168.1.1:1017")
	bounded.AddHost("192.168.1.1:1018")
	assert.Equal(t, true, reflect.DeepEqual(expect, bounded))
}

// TestBounded_RemoveHost .
func TestBounded_RemoveHost(t *testing.T) {
	expect, err := Build(BoundedBalancer, []string{"192.168.1.1:1015",
		"192.168.1.1:1016"})
	assert.Equal(t, err, nil)
	bounded := NewBounded([]string{"192.168.1.1:1015",
		"192.168.1.1:1016", "192.168.1.1:1017"})
	bounded.RemoveHost("192.168.1.1:1017")
	assert.Equal(t, true, reflect.DeepEqual(expect, bounded))
}

func TestBounded_Balance(t *testing.T) {
	expect, _ := Build(BoundedBalancer, []string{"192.168.1.1:1015",
		"192.168.1.1:1016", "192.168.1.1:1017", "192.168.1.1:1018"})
	expect.IncConn("192.168.1.1:1015")
	expect.IncConn("192.168.1.1:1015")
	expect.IncConn("NIL")
	expect.DoneConn("192.168.1.1:1015")
	expect.DoneConn("NIL")
	host, _ := expect.Balance("172.166.2.44")
	assert.Equal(t, "192.168.1.1:1017", host)
}
