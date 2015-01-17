package client

const (
	INSTANCE_HOST_MAP_TYPE = "instanceHostMap"
)

type InstanceHostMap struct {
	Resource
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    HostId string `json:"hostId,omitempty"`
    
    InstanceId string `json:"instanceId,omitempty"`
    
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

type InstanceHostMapCollection struct {
	Collection
	Data []InstanceHostMap `json:"data,omitempty"`
}

type InstanceHostMapClient struct {
	rancherClient *RancherClient
}

type InstanceHostMapOperations interface {
	List(opts *ListOpts) (*InstanceHostMapCollection, error)
	Create(opts *InstanceHostMap) (*InstanceHostMap, error)
	Update(existing *InstanceHostMap, updates interface{}) (*InstanceHostMap, error)
	ById(id string) (*InstanceHostMap, error)
	Delete(container *InstanceHostMap) error
}

func newInstanceHostMapClient(rancherClient *RancherClient) *InstanceHostMapClient {
	return &InstanceHostMapClient{
		rancherClient: rancherClient,
	}
}

func (c *InstanceHostMapClient) Create(container *InstanceHostMap) (*InstanceHostMap, error) {
	resp := &InstanceHostMap{}
	err := c.rancherClient.doCreate(INSTANCE_HOST_MAP_TYPE, container, resp)
	return resp, err
}

func (c *InstanceHostMapClient) Update(existing *InstanceHostMap, updates interface{}) (*InstanceHostMap, error) {
	resp := &InstanceHostMap{}
	err := c.rancherClient.doUpdate(INSTANCE_HOST_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *InstanceHostMapClient) List(opts *ListOpts) (*InstanceHostMapCollection, error) {
	resp := &InstanceHostMapCollection{}
	err := c.rancherClient.doList(INSTANCE_HOST_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *InstanceHostMapClient) ById(id string) (*InstanceHostMap, error) {
	resp := &InstanceHostMap{}
	err := c.rancherClient.doById(INSTANCE_HOST_MAP_TYPE, id, resp)
	return resp, err
}

func (c *InstanceHostMapClient) Delete(container *InstanceHostMap) error {
	return c.rancherClient.doResourceDelete(INSTANCE_HOST_MAP_TYPE, &container.Resource)
}
