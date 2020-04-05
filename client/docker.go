package client

import (
	"context"
	"fmt"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

func main() {
	fmt.Printf("NewEnvClient\n")
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	fmt.Printf("ContainerList\n")
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
