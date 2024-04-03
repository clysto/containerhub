package containerhub

import "testing"

func TestConnectHub(t *testing.T) {
	hub, err := ConnectHub()
	if err != nil {
		t.Errorf("ConnectHub() failed: %v", err)
	}
	containers, err := hub.ListContainers("maoyachen")
	if err != nil {
		t.Errorf("ListContainers() failed: %v", err)
	}
	for _, container := range containers {
		t.Logf("Container: %v", container)
	}
}

func TestCreateContainer(t *testing.T) {
	hub, err := ConnectHub()
	if err != nil {
		t.Errorf("ConnectHub() failed: %v", err)
	}
	id, err := hub.CreateContainer("containerhub-basic:latest", "maoyachen", "test", nil)
	if err != nil {
		t.Errorf("CreateContainer() failed: %v", err)
	}
	t.Logf("Container created: %v", id)
}
