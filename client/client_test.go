package client

import (
	"testing"
)

const (
	URL        = "http://localhost:8080/v1"
	ACCESS_KEY = "admin"
	SECRET_KEY = "adminpass"
)

func newClient(t *testing.T) *RancherClient {
	client, err := NewRancherClient(&ClientOpts{
		Url:       URL,
		AccessKey: ACCESS_KEY,
		SecretKey: SECRET_KEY,
	})

	if err != nil {
		t.Fatal("Failed to create client", err)
	}

	return client
}

func TestClientLoad(t *testing.T) {
	client := newClient(t)
	if client.Schemas == nil {
		t.Fatal("Failed to load schema")
	}

	if len(client.Schemas.Data) == 0 {
		t.Fatal("Schemas is empty")
	}

	if _, ok := client.Types["container"]; !ok {
		t.Fatal("Failed to find container type")
	}
}

func TestContainerList(t *testing.T) {
	client := newClient(t)

	/* Create a container to ensure list will return something */
	container, err := client.Container.Create(&Container{
		Name:      "a name",
		ImageUuid: "docker:nginx",
	})
	if err != nil {
		t.Fatal(err)
	}

	defer client.Container.Delete(container)

	containers, err := client.Container.List(nil)

	if err != nil {
		t.Fatal("Failed to list containers", err)
	}

	if len(containers.Data) == 0 {
		t.Fatal("No containers found")
	}

	if len(containers.Data[0].Id) == 0 {
		t.Fatal("Container ID is not set")
	}
}

func TestContainerCreate(t *testing.T) {
	client := newClient(t)
	container, err := client.Container.Create(&Container{
		Name:      "a name",
		ImageUuid: "docker:nginx",
	})

	if err != nil {
		t.Fatal(err)
	}

	defer client.Container.Delete(container)

	if container.Name != "a name" {
		t.Fatal("Field name is wrong [" + container.Name + "]")
	}

	if container.ImageUuid != "docker:nginx" {
		t.Fatal("Field imageUuid is wrong [" + container.ImageUuid + "]")
	}
}

func TestContainerUpdate(t *testing.T) {
	client := newClient(t)
	container, err := client.Container.Create(&Container{
		Name:      "a name",
		ImageUuid: "docker:nginx",
	})

	if err != nil {
		t.Fatal(err)
	}

	defer client.Container.Delete(container)

	if container.Name != "a name" {
		t.Fatal("Field name is wrong [" + container.Name + "]")
	}

	container, err = client.Container.Update(container, &Container{
		Name: "a different name",
	})

	if container.Name != "a different name" {
		t.Fatal("Field name is wrong [" + container.Name + "]")
	}

	by_id_container, err := client.Container.ById(string(container.Id))
	if err != nil {
		t.Fatal(err)
	}

	if by_id_container.Id != container.Id {
		t.Fatal("Container from by ID did not match")
	}

	if by_id_container.Name != container.Name {
		t.Fatal("Container from by ID did not match for name")
	}
}

func TestContainerDelete(t *testing.T) {
	client := newClient(t)
	container, err := client.Container.Create(&Container{
		Name:      "a name",
		ImageUuid: "docker:nginx",
	})

	if err != nil {
		t.Fatal(err)
	}

	err = client.Container.Delete(container)
	if err != nil {
		t.Fatal("Failed to delete", err)
	}
}

func TestContainerNotExists(t *testing.T) {
	client := newClient(t)
	_, err := client.Container.ById("badId1")
	if err == nil {
		t.Fatal("Should have received an error getting non-existent container.")
	}

	apiError, ok := err.(*ApiError)
	if !ok {
		t.Fatal("Should have received an ApiError.")
	}
	if apiError.StatusCode != 404 {
		t.Fatal("Should have received a 404 and reported it on the ApiError.")
	}
}
