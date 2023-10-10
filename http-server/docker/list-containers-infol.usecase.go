package docker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type listContainersInfoUseCase struct {}

type response struct {
	ContainerId string `json:"containerId"`
	Name []string `json:"name"`
	CpuUsage string `json:"cpuUsage"`
	MemoryUsage string `json:"memoryUsage"`
	Image string `json:"image"`
	State string `json:"state"`
	Status string `json:"status"`
	Ports []types.Port `json:"ports"`
}

func NewListContainersInfoUseCase() *listContainersInfoUseCase {
	return &listContainersInfoUseCase{}
}

func (t *listContainersInfoUseCase) Execute() []response {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithVersion("1.41"),
	)
	
	if err != nil {
		panic(err)
	}

	listOptions := types.ContainerListOptions{
		All: true,
	}

	containers, err := cli.ContainerList(context.Background(), listOptions)
	if err != nil {
		panic(err)
	}

	var responseData []response

	for _, container := range containers {
		stats, err := cli.ContainerStats(context.Background(), container.ID, false)
		if err != nil {
			panic(err)
		}
		defer stats.Body.Close()

		var statsData types.StatsJSON
		if err := json.NewDecoder(stats.Body).Decode(&statsData); err != nil {
			panic(err)
		}

		infos := response{
			ContainerId: container.ID,
			Name: container.Names,
			CpuUsage: fmt.Sprintf(" %.2f%%",calculateCPUPercentage(&statsData)),
			MemoryUsage: fmt.Sprintf(" %d",calculateMemoryUsage(&statsData)),
			Image: container.Image,
			State: container.State,
			Status: container.Status,
			Ports: container.Ports,
		}

		responseData = append(responseData, infos)
	}

	return responseData
}

func calculateCPUPercentage(stats *types.StatsJSON) float64 {
	cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage) - float64(stats.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(stats.CPUStats.SystemUsage) - float64(stats.PreCPUStats.SystemUsage)

	if systemDelta > 0 {
		cpuUsage := (cpuDelta / systemDelta) * float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) * 100.0
		
		return cpuUsage
	}

	return 0.00
}

func calculateMemoryUsage(stats *types.StatsJSON) uint64 {
	return stats.MemoryStats.Usage / (1024 * 1024)
}
