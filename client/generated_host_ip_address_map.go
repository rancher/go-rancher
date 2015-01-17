package client

const (
	HOST_IP_ADDRESS_MAP_TYPE = "hostIpAddressMap"
)

type HostIpAddressMap struct {
	Resource
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    HostId string `json:"hostId,omitempty"`
    
    IpAddressId string `json:"ipAddressId,omitempty"`
    
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

type HostIpAddressMapCollection struct {
	Collection
	Data []HostIpAddressMap `json:"data,omitempty"`
}

type HostIpAddressMapClient struct {
	rancherClient *RancherClient
}

type HostIpAddressMapOperations interface {
	List(opts *ListOpts) (*HostIpAddressMapCollection, error)
	Create(opts *HostIpAddressMap) (*HostIpAddressMap, error)
	Update(existing *HostIpAddressMap, updates interface{}) (*HostIpAddressMap, error)
	ById(id string) (*HostIpAddressMap, error)
	Delete(container *HostIpAddressMap) error
}

func newHostIpAddressMapClient(rancherClient *RancherClient) *HostIpAddressMapClient {
	return &HostIpAddressMapClient{
		rancherClient: rancherClient,
	}
}

func (c *HostIpAddressMapClient) Create(container *HostIpAddressMap) (*HostIpAddressMap, error) {
	resp := &HostIpAddressMap{}
	err := c.rancherClient.doCreate(HOST_IP_ADDRESS_MAP_TYPE, container, resp)
	return resp, err
}

func (c *HostIpAddressMapClient) Update(existing *HostIpAddressMap, updates interface{}) (*HostIpAddressMap, error) {
	resp := &HostIpAddressMap{}
	err := c.rancherClient.doUpdate(HOST_IP_ADDRESS_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HostIpAddressMapClient) List(opts *ListOpts) (*HostIpAddressMapCollection, error) {
	resp := &HostIpAddressMapCollection{}
	err := c.rancherClient.doList(HOST_IP_ADDRESS_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *HostIpAddressMapClient) ById(id string) (*HostIpAddressMap, error) {
	resp := &HostIpAddressMap{}
	err := c.rancherClient.doById(HOST_IP_ADDRESS_MAP_TYPE, id, resp)
	return resp, err
}

func (c *HostIpAddressMapClient) Delete(container *HostIpAddressMap) error {
	return c.rancherClient.doResourceDelete(HOST_IP_ADDRESS_MAP_TYPE, &container.Resource)
}
