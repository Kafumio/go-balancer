# go-balancer


`go-balancer` 是一个支持http、https协议及七种负载均衡算法的负载均衡器, 同时也是要一个实现了 `load balancing` 算法的go library。

它目前支持以下负载均衡算法：
* `round-robin`
* `random`
* `power of 2 random choice`
* `consistent hash`
* `consistent hash with bounded`
* `ip-hash`
* `least-load`

## Install
首先下载源码：
```shell
> git clone https://github.com/Kafumio/go-balancer.git
```
编译源码：
```shell
> cd ./go-balancer

> go build
```

## Run
`go-balancer` 需要配置 `config.yaml` 文件, see [config.yaml](https://github.com/Kafumio/go-balancer/blob/main/config.yaml) :

现在，你可以执行`balancer`, 负载均衡器将打印ASCII图和配置细节:
```shell
> ./balancer

___  ____ _    ____ _  _ ____ ____ ____ 
|__] |__| |    |__| |\ | |    |___ |__/ 
|__] |  | |___ |  | | \| |___ |___ |  \                                        

Schema: http
Port: 8089
Health Check: true
Location:
        Route: /
        Proxy Pass: [http://192.168.1.1 http://192.168.1.2:1015 https://192.168.1.2 http://my-server.com]
        Mode: round-robin

```
`go-balancer` 将定期对所有代理站点执行“健康检查”。当站点不可达时，它将自动从负载均衡器中删除。 . 然而, `go-balancer` 仍然会在无法访问的站点上执行“健康检查”。当站点可达时，它将自动将其添加到平衡器。

## API Usage
`balancer` 也是一个实现了负载均衡算法的 go library，它可以单独作为API使用，你需要首先将它导入到你的项目中：
```shell
> go get github.com/zehuamama/balancer/balancer
```

通过 `balancer.Build` 构建负载均衡器:
```go
hosts := []string{
	"http://192.168.11.101",
	"http://192.168.11.102",
	"http://192.168.11.103",
	"http://192.168.11.104",
}

lb, err := balancer.Build(balancer.P2CBalancer, hosts)
if err != nil {
	return err
}
```
你也可以像这样使用负载均衡器:
```go

clientAddr := "172.160.1.5"  // request IP
	
targetHost, err := lb.Balance(clientAddr) 
if err != nil {
	log.Fatal(err)
}
	
lb.Inc(targetHost)
defer lb.Done(targetHost)

// route to target host
```
每个负载均衡器都实现了 `balancer.Balancer` 的接口:
```go
type Balancer interface {
	Add(string)
	Remove(string)
	Balance(string) (string, error)
	Inc(string)
	Done(string)
}
```
currently supports the following load balancing algorithms:
```go
const (
	IPHashBalancer         = "ip-hash"
	ConsistentHashBalancer = "consistent-hash"
	P2CBalancer            = "p2c"
	RandomBalancer         = "random"
	R2Balancer             = "round-robin"
	LeastLoadBalancer      = "least-load"
	BoundedBalancer        = "bounded"
)
```

## License

balancer 是根据条款授权的：[BSD 2-Clause License](https://github.com/Kafumio/go-balancer/blob/main/LICENSE)
