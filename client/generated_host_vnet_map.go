package client

const (
	HOST_VNET_MAP_TYPE = "hostVnetMap"
)

type HostVnetMap struct {
	Resource
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    HostId string `json:"hostId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    VnetId string `json:"vnetId,omitempty"`
    
}

type HostVnetMapCollection struct {
	Collection
	Data []HostVnetMap `json:"data,omitempty"`
}

type HostVnetMapClient struct {
	rancherClient *RancherClient
}

type HostVnetMapOperations interface {
	List(opts *ListOpts) (*HostVnetMapCollection, error)
	Create(opts *HostVnetMap) (*HostVnetMap, error)
	Update(existing *HostVnetMap, updates interface{}) (*HostVnetMap, error)
	ById(id string) (*HostVnetMap, error)
	Delete(container *HostVnetMap) error
}

func newHostVnetMapClient(rancherClient *RancherClient) *HostVnetMapClient {
	return &HostVnetMapClient{
		rancherClient: rancherClient,
	}
}

func (c *HostVnetMapClient) Create(container *HostVnetMap) (*HostVnetMap, error) {
	resp := &HostVnetMap{}
	err := c.rancherClient.doCreate(HOST_VNET_MAP_TYPE, container, resp)
	return resp, err
}

func (c *HostVnetMapClient) Update(existing *HostVnetMap, updates interface{}) (*HostVnetMap, error) {
	resp := &HostVnetMap{}
	err := c.rancherClient.doUpdate(HOST_VNET_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HostVnetMapClient) List(opts *ListOpts) (*HostVnetMapCollection, error) {
	resp := &HostVnetMapCollection{}
	err := c.rancherClient.doList(HOST_VNET_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *HostVnetMapClient) ById(id string) (*HostVnetMap, error) {
	resp := &HostVnetMap{}
	err := c.rancherClient.doById(HOST_VNET_MAP_TYPE, id, resp)
	return resp, err
}

func (c *HostVnetMapClient) Delete(container *HostVnetMap) error {
	return c.rancherClient.doResourceDelete(HOST_VNET_MAP_TYPE, &container.Resource)
}
