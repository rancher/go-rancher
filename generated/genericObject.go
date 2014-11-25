package client

const (
	CONTAINER_TYPE = "genericObject"
)

type GenericObject struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Key string `json:"Key,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type GenericObjectCollection struct {
	Collection
	Data []GenericObject `json:"data,omitempty"`
}

type GenericObjectClient struct {
	rancherClient *RancherClient
}

type GenericObjectOperations interface {
	List(opts *ListOpts) (*GenericObjectCollection, error)
	Create(opts *GenericObject) (*GenericObject, error)
	Update(existing *GenericObject, updates interface{}) (*GenericObject, error)
	ById(id string) (*GenericObject, error)
	Delete(container *GenericObject) error
}

func newGenericObjectClient(rancherClient *RancherClient) *GenericObjectClient {
	return &GenericObjectClient{
		rancherClient: rancherClient,
	}
}

func (self *GenericObjectClient) Create(container *GenericObject) (*GenericObject, error) {
	resp := &GenericObject{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *GenericObjectClient) Update(existing *GenericObject, updates interface{}) (*GenericObject, error) {
	resp := &GenericObject{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *GenericObjectClient) List(opts *ListOpts) (*GenericObjectCollection, error) {
	resp := &GenericObjectCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *GenericObjectClient) ById(id string) (*GenericObject, error) {
	resp := &GenericObject{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *GenericObjectClient) Delete(container *GenericObject) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
