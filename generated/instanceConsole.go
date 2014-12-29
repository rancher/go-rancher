package client

const (
	CONTAINER_TYPE = "instanceConsole"
)

type InstanceConsole struct {
	Resource
    
    Kind string `json:"Kind,omitempty"`
    
    Password string `json:"Password,omitempty"`
    
    Url string `json:"Url,omitempty"`
    
}

type InstanceConsoleCollection struct {
	Collection
	Data []InstanceConsole `json:"data,omitempty"`
}

type InstanceConsoleClient struct {
	rancherClient *RancherClient
}

type InstanceConsoleOperations interface {
	List(opts *ListOpts) (*InstanceConsoleCollection, error)
	Create(opts *InstanceConsole) (*InstanceConsole, error)
	Update(existing *InstanceConsole, updates interface{}) (*InstanceConsole, error)
	ById(id string) (*InstanceConsole, error)
	Delete(container *InstanceConsole) error
}

func newInstanceConsoleClient(rancherClient *RancherClient) *InstanceConsoleClient {
	return &InstanceConsoleClient{
		rancherClient: rancherClient,
	}
}

func (self *InstanceConsoleClient) Create(container *InstanceConsole) (*InstanceConsole, error) {
	resp := &InstanceConsole{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *InstanceConsoleClient) Update(existing *InstanceConsole, updates interface{}) (*InstanceConsole, error) {
	resp := &InstanceConsole{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *InstanceConsoleClient) List(opts *ListOpts) (*InstanceConsoleCollection, error) {
	resp := &InstanceConsoleCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *InstanceConsoleClient) ById(id string) (*InstanceConsole, error) {
	resp := &InstanceConsole{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *InstanceConsoleClient) Delete(container *InstanceConsole) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
