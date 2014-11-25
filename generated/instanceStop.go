package client

const (
	CONTAINER_TYPE = "instanceStop"
)

type InstanceStop struct {
	Resource
    
    DeallocateFromHost boolean `json:"DeallocateFromHost,omitempty"`
    
    Remove boolean `json:"Remove,omitempty"`
    
}

type InstanceStopCollection struct {
	Collection
	Data []InstanceStop `json:"data,omitempty"`
}

type InstanceStopClient struct {
	rancherClient *RancherClient
}

type InstanceStopOperations interface {
	List(opts *ListOpts) (*InstanceStopCollection, error)
	Create(opts *InstanceStop) (*InstanceStop, error)
	Update(existing *InstanceStop, updates interface{}) (*InstanceStop, error)
	ById(id string) (*InstanceStop, error)
	Delete(container *InstanceStop) error
}

func newInstanceStopClient(rancherClient *RancherClient) *InstanceStopClient {
	return &InstanceStopClient{
		rancherClient: rancherClient,
	}
}

func (self *InstanceStopClient) Create(container *InstanceStop) (*InstanceStop, error) {
	resp := &InstanceStop{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *InstanceStopClient) Update(existing *InstanceStop, updates interface{}) (*InstanceStop, error) {
	resp := &InstanceStop{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *InstanceStopClient) List(opts *ListOpts) (*InstanceStopCollection, error) {
	resp := &InstanceStopCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *InstanceStopClient) ById(id string) (*InstanceStop, error) {
	resp := &InstanceStop{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *InstanceStopClient) Delete(container *InstanceStop) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
