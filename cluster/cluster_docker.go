package cluster

import "fmt"

type Docker struct {
	replicas int
}

func (cl Docker) Init() (status <-chan string, err <-chan error) {
	fmt.Println("DockerCluster Init " + string(cl.replicas))

	return
}
