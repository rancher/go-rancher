package client

const (
	CONTAINER_TYPE = "resourceDefinition"
)

type ResourceDefinition struct {
	Resource
    
    Id string `json:"Id,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
}

type ResourceDefinitionCollection struct {
	Collection
	Data []ResourceDefinition `json:"data,omitempty"`
}

type ResourceDefinitionClient struct {
	rancherClient *RancherClient
}

type ResourceDefinitionOperations interface {
	List(opts *ListOpts) (*ResourceDefinitionCollection, error)
	Create(opts *ResourceDefinition) (*ResourceDefinition, error)
	Update(existing *ResourceDefinition, updates interface{}) (*ResourceDefinition, error)
	ById(id string) (*ResourceDefinition, error)
	Delete(container *ResourceDefinition) error
}

func newResourceDefinitionClient(rancherClient *RancherClient) *ResourceDefinitionClient {
	return &ResourceDefinitionClient{
		rancherClient: rancherClient,
	}
}

func (self *ResourceDefinitionClient) Create(container *ResourceDefinition) (*ResourceDefinition, error) {
	resp := &ResourceDefinition{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ResourceDefinitionClient) Update(existing *ResourceDefinition, updates interface{}) (*ResourceDefinition, error) {
	resp := &ResourceDefinition{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ResourceDefinitionClient) List(opts *ListOpts) (*ResourceDefinitionCollection, error) {
	resp := &ResourceDefinitionCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ResourceDefinitionClient) ById(id string) (*ResourceDefinition, error) {
	resp := &ResourceDefinition{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ResourceDefinitionClient) Delete(container *ResourceDefinition) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
