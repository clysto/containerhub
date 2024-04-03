package containerhub

import "testing"

var caPubkeyPEM = []byte(`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC0/E3YoVHyGCq1v0DB4dkrAE5KGARx9s2Q3itT4ysigpd5dlK0AkrD7AzzxceDFmrp/kufRoWXAWjVjZMhWbacLOe4V9uczxZalUBvkcgz1kshp6yXMlFMsDZNZKG08Bu+UjWQ+Sl2aKpfY1qpvioE06+mOooeldZDVMIKF6Zm/IcZat/SnJG6U1+B2Zd9La0xyeDSHfLp9YtrXLkoFxalOlSqv7s35rAfmwW8/qmGzqFIPVtTsddqg9m4cksdvjIU65kzkLXv7+hHBE6CdxY6ExT0gOnvy3W8jdzcmcSxxQI6mytad31P8g4/6PAAg3fyK+cBifvZpcTn5z3XiBvtd3ha3VB6ryxGy5N8bBbNQLZa45C6U4CJNs3I1jghjkrjmYdDCLL3KLONOIqvdB7bJ4YsJ6nOvBq9T9/wPGgHYwSqLtyPc4RLCM6dqjZd1LlwDj37qqf1BTdWZtDzgxzLxX+WIUiVIHP33t/4SknJ9RcYMrHmQJXtJ4MwKEQ0jPisVTFScYZ9rlLWpLvr93Yx6GFU6pvmpG8PbgrWe08zzehemT+x5PUiuTWNuYzfv7vdVDWxEQ+HJaI6ug79wwxIvaSEEmbhZAOtGOVDcnUM19q2qYMPLslkTjuyugcUvJU/Ju5+sVIcnisi9C8KNEX8EMTL1JTleET94eqOf6UxzQ== user_ca`)

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
	id, err := hub.CreateContainer("containerhub-basic:latest", "maoyachen", "test", caPubkeyPEM)
	if err != nil {
		t.Errorf("CreateContainer() failed: %v", err)
	}
	t.Logf("Container created: %v", id)
}
