package client

const (
	CONTAINER_TYPE = "resource"
)

type Resource struct {
	Resource
    
    Actions map[json] `json:"Actions,omitempty"`
    
    Id string `json:"Id,omitempty"`
    
    Links map[json] `json:"Links,omitempty"`
    
    Type string `json:"Type,omitempty"`
    
}

type ResourceCollection struct {
	Collection
	Data []Resource `json:"data,omitempty"`
}

type ResourceClient struct {
	rancherClient *RancherClient
}

type ResourceOperations interface {
	List(opts *ListOpts) (*ResourceCollection, error)
	Create(opts *Resource) (*Resource, error)
	Update(existing *Resource, updates interface{}) (*Resource, error)
	ById(id string) (*Resource, error)
	Delete(container *Resource) error
}

func newResourceClient(rancherClient *RancherClient) *ResourceClient {
	return &ResourceClient{
		rancherClient: rancherClient,
	}
}

func (self *ResourceClient) Create(container *Resource) (*Resource, error) {
	resp := &Resource{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ResourceClient) Update(existing *Resource, updates interface{}) (*Resource, error) {
	resp := &Resource{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ResourceClient) List(opts *ListOpts) (*ResourceCollection, error) {
	resp := &ResourceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ResourceClient) ById(id string) (*Resource, error) {
	resp := &Resource{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ResourceClient) Delete(container *Resource) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
