package containerhub

import (
	"context"
	"encoding/base64"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/lithammer/shortuuid/v4"
)

type Hub struct {
	client *client.Client
	ctx    context.Context
}

type ContainerNotFoundError struct{}

func (e ContainerNotFoundError) Error() string {
	return "container not found"
}

func (h *Hub) ensureNetwork(networkName string) error {
	networks, err := h.client.NetworkList(h.ctx, types.NetworkListOptions{})
	if err != nil {
		return err
	}
	for _, net := range networks {
		if net.Name == networkName {
			// Network already exists
			return nil
		}
	}
	_, err = h.client.NetworkCreate(h.ctx, networkName, types.NetworkCreate{})
	if err != nil {
		return err
	}
	return nil
}

func ConnectHub() (*Hub, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	ctx := context.Background()
	if err != nil {
		return nil, err
	}
	client.NegotiateAPIVersion(ctx)
	return &Hub{client: client, ctx: ctx}, nil
}

func (h *Hub) GetContainer(containerID string) (types.Container, error) {
	filterArgs := filters.NewArgs()
	filterArgs.Add("label", "containerhub")
	filterArgs.Add("id", containerID)
	containers, err := h.client.ContainerList(h.ctx, container.ListOptions{
		All:     true,
		Filters: filterArgs,
	})
	if err != nil {
		return types.Container{}, err
	}
	if len(containers) == 0 {
		return types.Container{}, ContainerNotFoundError{}
	}
	return containers[0], nil
}

func (h *Hub) ListContainers(user string) ([]types.Container, error) {
	filterArgs := filters.NewArgs()
	filterArgs.Add("label", "containerhub")
	if user != "" {
		filterArgs.Add("label", "containerhub-user="+user)
	}
	containers, err := h.client.ContainerList(h.ctx, container.ListOptions{
		All:     true,
		Filters: filterArgs,
	})
	return containers, err
}

func (h *Hub) CreateContainer(image string, user string, customName string, caPubKeyPEM []byte) (string, error) {
	name := "containerhub-" + shortuuid.New()
	err := h.ensureNetwork("containerhub")
	if err != nil {
		return "", err
	}
	base64CaPubKey := base64.StdEncoding.EncodeToString(caPubKeyPEM)
	resp, err := h.client.ContainerCreate(h.ctx, &container.Config{
		Image: image,
		Labels: map[string]string{
			"containerhub":          "v1",
			"containerhub-user":     user,
			"containerhub-name":     customName,
			"containerhub-hostname": name,
		},
		Env: []string{
			"CONTAINERHUB_USER=" + user,
			"CONTAINERHUB_CA_PUB_KEY=" + base64CaPubKey,
		},
		ExposedPorts: nat.PortSet{
			"22/tcp": struct{}{},
		},
		Hostname: customName,
	}, &container.HostConfig{
		NetworkMode: "containerhub",
	}, nil, nil, name)
	return resp.ID, err
}

func (h *Hub) StartContainer(containerID string) error {
	return h.client.ContainerStart(h.ctx, containerID, container.StartOptions{})
}

func (h *Hub) StopContainer(containerID string) error {
	return h.client.ContainerStop(h.ctx, containerID, container.StopOptions{})
}

func (h *Hub) DestroyContainer(containerID string) error {
	return h.client.ContainerRemove(h.ctx, containerID, container.RemoveOptions{})
}
