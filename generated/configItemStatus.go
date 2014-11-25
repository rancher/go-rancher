package client

const (
	CONTAINER_TYPE = "configItemStatus"
)

type ConfigItemStatus struct {
	Resource
    
    AgentId string `json:"AgentId,omitempty"`
    
    AppliedUpdated date `json:"AppliedUpdated,omitempty"`
    
    AppliedVersion int `json:"AppliedVersion,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RequestedUpdated date `json:"RequestedUpdated,omitempty"`
    
    RequestedVersion int `json:"RequestedVersion,omitempty"`
    
    SourceVersion string `json:"SourceVersion,omitempty"`
    
}

type ConfigItemStatusCollection struct {
	Collection
	Data []ConfigItemStatus `json:"data,omitempty"`
}

type ConfigItemStatusClient struct {
	rancherClient *RancherClient
}

type ConfigItemStatusOperations interface {
	List(opts *ListOpts) (*ConfigItemStatusCollection, error)
	Create(opts *ConfigItemStatus) (*ConfigItemStatus, error)
	Update(existing *ConfigItemStatus, updates interface{}) (*ConfigItemStatus, error)
	ById(id string) (*ConfigItemStatus, error)
	Delete(container *ConfigItemStatus) error
}

func newConfigItemStatusClient(rancherClient *RancherClient) *ConfigItemStatusClient {
	return &ConfigItemStatusClient{
		rancherClient: rancherClient,
	}
}

func (self *ConfigItemStatusClient) Create(container *ConfigItemStatus) (*ConfigItemStatus, error) {
	resp := &ConfigItemStatus{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ConfigItemStatusClient) Update(existing *ConfigItemStatus, updates interface{}) (*ConfigItemStatus, error) {
	resp := &ConfigItemStatus{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ConfigItemStatusClient) List(opts *ListOpts) (*ConfigItemStatusCollection, error) {
	resp := &ConfigItemStatusCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ConfigItemStatusClient) ById(id string) (*ConfigItemStatus, error) {
	resp := &ConfigItemStatus{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ConfigItemStatusClient) Delete(container *ConfigItemStatus) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
