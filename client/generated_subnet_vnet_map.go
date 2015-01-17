package client

const (
	SUBNET_VNET_MAP_TYPE = "subnetVnetMap"
)

type SubnetVnetMap struct {
	Resource
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    SubnetId string `json:"subnetId,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    VnetId string `json:"vnetId,omitempty"`
    
}

type SubnetVnetMapCollection struct {
	Collection
	Data []SubnetVnetMap `json:"data,omitempty"`
}

type SubnetVnetMapClient struct {
	rancherClient *RancherClient
}

type SubnetVnetMapOperations interface {
	List(opts *ListOpts) (*SubnetVnetMapCollection, error)
	Create(opts *SubnetVnetMap) (*SubnetVnetMap, error)
	Update(existing *SubnetVnetMap, updates interface{}) (*SubnetVnetMap, error)
	ById(id string) (*SubnetVnetMap, error)
	Delete(container *SubnetVnetMap) error
}

func newSubnetVnetMapClient(rancherClient *RancherClient) *SubnetVnetMapClient {
	return &SubnetVnetMapClient{
		rancherClient: rancherClient,
	}
}

func (c *SubnetVnetMapClient) Create(container *SubnetVnetMap) (*SubnetVnetMap, error) {
	resp := &SubnetVnetMap{}
	err := c.rancherClient.doCreate(SUBNET_VNET_MAP_TYPE, container, resp)
	return resp, err
}

func (c *SubnetVnetMapClient) Update(existing *SubnetVnetMap, updates interface{}) (*SubnetVnetMap, error) {
	resp := &SubnetVnetMap{}
	err := c.rancherClient.doUpdate(SUBNET_VNET_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SubnetVnetMapClient) List(opts *ListOpts) (*SubnetVnetMapCollection, error) {
	resp := &SubnetVnetMapCollection{}
	err := c.rancherClient.doList(SUBNET_VNET_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *SubnetVnetMapClient) ById(id string) (*SubnetVnetMap, error) {
	resp := &SubnetVnetMap{}
	err := c.rancherClient.doById(SUBNET_VNET_MAP_TYPE, id, resp)
	return resp, err
}

func (c *SubnetVnetMapClient) Delete(container *SubnetVnetMap) error {
	return c.rancherClient.doResourceDelete(SUBNET_VNET_MAP_TYPE, &container.Resource)
}
