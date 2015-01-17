package client

const (
	GENERIC_OBJECT_TYPE = "genericObject"
)

type GenericObject struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Key string `json:"key,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
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

func (c *GenericObjectClient) Create(container *GenericObject) (*GenericObject, error) {
	resp := &GenericObject{}
	err := c.rancherClient.doCreate(GENERIC_OBJECT_TYPE, container, resp)
	return resp, err
}

func (c *GenericObjectClient) Update(existing *GenericObject, updates interface{}) (*GenericObject, error) {
	resp := &GenericObject{}
	err := c.rancherClient.doUpdate(GENERIC_OBJECT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GenericObjectClient) List(opts *ListOpts) (*GenericObjectCollection, error) {
	resp := &GenericObjectCollection{}
	err := c.rancherClient.doList(GENERIC_OBJECT_TYPE, opts, resp)
	return resp, err
}

func (c *GenericObjectClient) ById(id string) (*GenericObject, error) {
	resp := &GenericObject{}
	err := c.rancherClient.doById(GENERIC_OBJECT_TYPE, id, resp)
	return resp, err
}

func (c *GenericObjectClient) Delete(container *GenericObject) error {
	return c.rancherClient.doResourceDelete(GENERIC_OBJECT_TYPE, &container.Resource)
}
