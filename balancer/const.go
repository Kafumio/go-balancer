package balancer

/**
  @author: Wow
  @since: 2023/3/9
  @desc: //TODO constant of balance algorithms
**/

const (
	IPHashBalancer         = "ip-hash"
	ConsistentHashBalancer = "consistent-hash"
	P2CBalancer            = "p2c"
	RandomBalancer         = "random"
	RoundRobinBalancer     = "round-robin"
	LeastLoadBalancer      = "least-load"
	BoundedBalancer        = "bounded"
)
