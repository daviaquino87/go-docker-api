package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type restartContainerUseCase struct{}

func NewRestartContainerUseCase() *restartContainerUseCase {
	return &restartContainerUseCase{}
}

func (r *restartContainerUseCase) Execute(id string) error {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithVersion("1.41"),
	)

	if err != nil {
		return err
	}

	var options container.StopOptions

	err = cli.ContainerRestart(context.Background(), id, options)
	if err != nil {
		return err
	}

	return nil
}
