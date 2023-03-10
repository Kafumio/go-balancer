package balancer

import "hash/crc32"

/**
  @author: kafmio
  @since: 2023/3/9
  @desc: //TODO Implement the ip_hash algorithm
**/

func init() {
	factories[IPHashBalancer] = NewIPHash
}

type IPHash struct {
	BaseBalancer
}

func NewIPHash(hosts []string) Balancer {
	return &IPHash{
		BaseBalancer: BaseBalancer{hosts: hosts},
	}
}

func (i *IPHash) Balance(key string) (string, error) {
	i.RLock()
	defer i.RUnlock()
	if len(i.hosts) == 0 {
		return "", NoHostError
	}
	value := crc32.ChecksumIEEE([]byte(key)) % uint32(len(i.hosts))
	return i.hosts[value], nil
}
