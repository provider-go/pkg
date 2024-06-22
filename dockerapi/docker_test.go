package dockerapi

import (
	"encoding/json"
	"testing"
)

func TestContainerList(t *testing.T) {
	req := DockerContainerList{All: true}
	res, err := ContainerList(&req)
	if err != nil {
		t.Log(err)
	}
	b, _ := json.Marshal(res)
	t.Log(string(b))
}

func TestContainerStart(t *testing.T) {
	m := DockerContainerStart{Id: "ab0c598f55"}
	err := ContainerStart(&m)
	if err != nil {
		t.Log(err)
	}
}

func TestContainerStop(t *testing.T) {
	m := DockerContainerStop{Id: "ab0c598f55"}
	err := ContainerStop(&m)
	if err != nil {
		t.Log(err)
	}
}

func TestContainerRestart(t *testing.T) {
	m := DockerContainerRestart{Id: "ab0c598f55"}
	err := ContainerRestart(&m)
	if err != nil {
		t.Log(err)
	}
}

func TestContainerRemove(t *testing.T) {
	m := DockerContainerRemove{Id: "ab0c598f55"}
	err := ContainerRemove(&m)
	if err != nil {
		t.Log(err)
	}
}

func TestContainerLogs(t *testing.T) {
	m := DockerContainerLogs{Id: "ab0c598f55"}
	res, err := ContainerLogs(&m, "10")
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}

func TestImageList(t *testing.T) {
	req := DockerImageList{All: true}
	res, err := ImageList(&req)
	if err != nil {
		t.Log(err)
	}
	b, _ := json.Marshal(res)
	t.Log(string(b))
}

func TestImageRemove(t *testing.T) {
	m := DockerImageRemove{Id: "7c33"}
	err := ImageRemove(&m)
	if err != nil {
		t.Log(err)
	}
}

func TestImagesPrune(t *testing.T) {
	req := DockerImagesPrune{All: true}
	res, err := ImagesPrune(&req)
	if err != nil {
		t.Log(err)
	}
	b, _ := json.Marshal(res)
	t.Log(string(b))
}

func TestNetworkList(t *testing.T) {
	res, err := NetworkList()
	if err != nil {
		t.Log(err)
	}
	b, _ := json.Marshal(res)
	t.Log(string(b))
}

func TestNetworkRemove(t *testing.T) {
	req := DockerNetworkRemove{}
	err := NetworkRemove(&req)
	if err != nil {
		t.Log(err)
	}
}

func TestNetworksPrune(t *testing.T) {
	req := DockerNetworksPrune{}
	res, err := NetworksPrune(&req)
	if err != nil {
		t.Log(err)
	}
	b, _ := json.Marshal(res)
	t.Log(string(b))
}

func TestVolumeList(t *testing.T) {
	res, err := VolumeList()
	if err != nil {
		t.Log(err)
	}
	b, _ := json.Marshal(res)
	t.Log(string(b))
}

func TestVolumeRemove(t *testing.T) {
	req := DockerVolumeRemove{}
	err := VolumeRemove(&req)
	if err != nil {
		t.Log(err)
	}
}

func TestVolumesPrune(t *testing.T) {
	req := DockerVolumesPrune{All: true}
	res, err := VolumesPrune(&req)
	if err != nil {
		t.Log(err)
	}
	b, _ := json.Marshal(res)
	t.Log(string(b))
}
