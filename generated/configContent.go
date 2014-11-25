package client

const (
	CONTAINER_TYPE = "configContent"
)

type ConfigContent struct {
	Resource
    
    AgentId string `json:"AgentId,omitempty"`
    
    Current string `json:"Current,omitempty"`
    
    Id string `json:"Id,omitempty"`
    
    Version string `json:"Version,omitempty"`
    
}

type ConfigContentCollection struct {
	Collection
	Data []ConfigContent `json:"data,omitempty"`
}

type ConfigContentClient struct {
	rancherClient *RancherClient
}

type ConfigContentOperations interface {
	List(opts *ListOpts) (*ConfigContentCollection, error)
	Create(opts *ConfigContent) (*ConfigContent, error)
	Update(existing *ConfigContent, updates interface{}) (*ConfigContent, error)
	ById(id string) (*ConfigContent, error)
	Delete(container *ConfigContent) error
}

func newConfigContentClient(rancherClient *RancherClient) *ConfigContentClient {
	return &ConfigContentClient{
		rancherClient: rancherClient,
	}
}

func (self *ConfigContentClient) Create(container *ConfigContent) (*ConfigContent, error) {
	resp := &ConfigContent{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ConfigContentClient) Update(existing *ConfigContent, updates interface{}) (*ConfigContent, error) {
	resp := &ConfigContent{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ConfigContentClient) List(opts *ListOpts) (*ConfigContentCollection, error) {
	resp := &ConfigContentCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ConfigContentClient) ById(id string) (*ConfigContent, error) {
	resp := &ConfigContent{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ConfigContentClient) Delete(container *ConfigContent) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
