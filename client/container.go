package client

const (
	CONTAINER_TYPE = "container"
)

type Container struct {
	Resource
	Name      string `json:"name,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	ImageId   string `json:"imageId,omitempty"`
}

type ContainerCollection struct {
	Collection
	Data []Container `json:"data,omitempty"`
}

type ContainerClient struct {
	rancherClient *RancherClient
}

type ContainerOperations interface {
	List(opts *ListOpts) (*ContainerCollection, error)
	Create(opts *Container) (*Container, error)
	Update(existing *Container, updates interface{}) (*Container, error)
	ById(id string) (*Container, error)
	Delete(container *Container) error
}

func newContainerClient(rancherClient *RancherClient) *ContainerClient {
	return &ContainerClient{
		rancherClient: rancherClient,
	}
}

func (self *ContainerClient) Create(container *Container) (*Container, error) {
	resp := &Container{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ContainerClient) Update(existing *Container, updates interface{}) (*Container, error) {
	resp := &Container{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ContainerClient) List(opts *ListOpts) (*ContainerCollection, error) {
	resp := &ContainerCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ContainerClient) ById(id string) (*Container, error) {
	resp := &Container{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ContainerClient) Delete(container *Container) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
