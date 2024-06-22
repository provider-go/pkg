package dockerapi

import (
	"context"
	"encoding/binary"
	"sort"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var containerStaleStatus map[string]string

const (
	StaleStatusProcessing = "processing"
	StaleStatusYes        = "yes"
	StaleStatusNo         = "no"
	StaleStatusError      = "error"
)

// ContainerList 容器列表
func ContainerList(req *DockerContainerList) (*DockerContainerListResponse, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	dcontainers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: req.All})
	if err != nil {
		return nil, err
	}

	containers := make([]Container, len(dcontainers))
	for i, c := range dcontainers {
		ports := make([]Port, len(c.Ports))
		for j, port := range c.Ports {
			ports[j] = Port(port)
		}

		image := strings.Split(c.Image, "@")[0]
		stale, ok := containerStaleStatus[c.ID]
		if !ok {
			stale = StaleStatusProcessing
		}

		containers[i] = Container{
			Id:     c.ID[0:10],
			Name:   c.Names[0][1:],
			Image:  image,
			Status: c.Status,
			State:  c.State,
			Ports:  ports,
			Stale:  stale,
		}
	}

	sort.Slice(containers, func(i, j int) bool {
		return containers[i].Name < containers[j].Name
	})

	return &DockerContainerListResponse{Items: containers}, nil
}

// ContainerStart 启动容器
func ContainerStart(req *DockerContainerStart) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	err = cli.ContainerStart(context.Background(), req.Id, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	return nil
}

// ContainerStop 关闭容器
func ContainerStop(req *DockerContainerStop) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	err = cli.ContainerStop(context.Background(), req.Id, container.StopOptions{})
	if err != nil {
		return err
	}

	return nil
}

// ContainerRestart 重启容器
func ContainerRestart(req *DockerContainerRestart) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	err = cli.ContainerRestart(context.Background(), req.Id, container.StopOptions{})
	if err != nil {
		return err
	}

	return nil
}

// ContainerRemove 删除容器
func ContainerRemove(req *DockerContainerRemove) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	err = cli.ContainerRemove(context.Background(), req.Id, types.ContainerRemoveOptions{Force: req.Force})
	if err != nil {
		return err
	}

	return nil
}

// ContainerLogs 查看容器日志
func ContainerLogs(req *DockerContainerLogs, num string) ([]string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	o := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Tail:       num,
	}

	reader, err := cli.ContainerLogs(context.Background(), req.Id, o)
	if err != nil {
		return nil, err
	}

	defer reader.Close()

	var logs []string
	hdr := make([]byte, 8)
	for {
		_, err = reader.Read(hdr)

		if err != nil {
			return logs, err
		}

		count := binary.BigEndian.Uint32(hdr[4:])
		dat := make([]byte, count)
		_, err = reader.Read(dat)
		logs = append(logs, string(dat))

		if err != nil {
			return logs, err
		}
	}
}
