package dockerapi

import (
	"context"
	"sort"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// NetworkList 网络列表
func NetworkList() (*DockerNetworkListResponse, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	dcontainers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	usedNetworks := make(map[string]interface{}, 0)
	for _, c := range dcontainers {
		if c.NetworkSettings != nil {
			for _, n := range c.NetworkSettings.Networks {
				usedNetworks[n.NetworkID] = nil
			}
		}
	}

	dnetworks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return nil, err
	}

	networks := make([]Network, len(dnetworks))
	for i, item := range dnetworks {
		_, inUse := usedNetworks[item.ID]
		networks[i] = Network{
			Id:     item.ID[0:10],
			Name:   item.Name,
			Driver: item.Driver,
			Scope:  item.Scope,
			InUse:  inUse,
		}
	}

	sort.Slice(networks, func(i, j int) bool {
		return networks[i].Name < networks[j].Name
	})

	return &DockerNetworkListResponse{Items: networks}, nil
}

// NetworkRemove 移除网络
func NetworkRemove(req *DockerNetworkRemove) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	err = cli.NetworkRemove(context.Background(), req.Id)
	if err != nil {
		return err
	}

	return nil
}

// NetworksPrune 修剪网络
func NetworksPrune(req *DockerNetworksPrune) (*DockerNetworksPruneResponse, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	report, err := cli.NetworksPrune(context.Background(), filters.NewArgs())
	if err != nil {
		return nil, err
	}

	return &DockerNetworksPruneResponse{NetworksDeleted: report.NetworksDeleted}, nil
}
