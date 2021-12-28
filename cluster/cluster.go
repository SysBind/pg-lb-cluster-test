package cluster

type Cluster interface {
	Init() (<-chan string, <-chan error)
}
