package client

const (
	CONTAINER_TYPE = "instanceConsoleInput"
)

type InstanceConsoleInput struct {
	Resource
    
}

type InstanceConsoleInputCollection struct {
	Collection
	Data []InstanceConsoleInput `json:"data,omitempty"`
}

type InstanceConsoleInputClient struct {
	rancherClient *RancherClient
}

type InstanceConsoleInputOperations interface {
	List(opts *ListOpts) (*InstanceConsoleInputCollection, error)
	Create(opts *InstanceConsoleInput) (*InstanceConsoleInput, error)
	Update(existing *InstanceConsoleInput, updates interface{}) (*InstanceConsoleInput, error)
	ById(id string) (*InstanceConsoleInput, error)
	Delete(container *InstanceConsoleInput) error
}

func newInstanceConsoleInputClient(rancherClient *RancherClient) *InstanceConsoleInputClient {
	return &InstanceConsoleInputClient{
		rancherClient: rancherClient,
	}
}

func (self *InstanceConsoleInputClient) Create(container *InstanceConsoleInput) (*InstanceConsoleInput, error) {
	resp := &InstanceConsoleInput{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *InstanceConsoleInputClient) Update(existing *InstanceConsoleInput, updates interface{}) (*InstanceConsoleInput, error) {
	resp := &InstanceConsoleInput{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *InstanceConsoleInputClient) List(opts *ListOpts) (*InstanceConsoleInputCollection, error) {
	resp := &InstanceConsoleInputCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *InstanceConsoleInputClient) ById(id string) (*InstanceConsoleInput, error) {
	resp := &InstanceConsoleInput{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *InstanceConsoleInputClient) Delete(container *InstanceConsoleInput) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
