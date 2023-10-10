package docker

import (
	"context"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type startContainerUseCase struct {}

func NewStartContainerUseCase () *startContainerUseCase {
	return &startContainerUseCase{}
}

func (t startContainerUseCase) Execute(id string) error {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithVersion("1.41"),
	)

	if err != nil {
		return err
	}

	var options dockerTypes.ContainerStartOptions

	err = cli.ContainerStart(context.Background(), id, options)
	if err != nil {
		return err
	}

	return nil
}