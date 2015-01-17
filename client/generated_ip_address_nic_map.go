package client

const (
	IP_ADDRESS_NIC_MAP_TYPE = "ipAddressNicMap"
)

type IpAddressNicMap struct {
	Resource
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    IpAddressId string `json:"ipAddressId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NicId string `json:"nicId,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type IpAddressNicMapCollection struct {
	Collection
	Data []IpAddressNicMap `json:"data,omitempty"`
}

type IpAddressNicMapClient struct {
	rancherClient *RancherClient
}

type IpAddressNicMapOperations interface {
	List(opts *ListOpts) (*IpAddressNicMapCollection, error)
	Create(opts *IpAddressNicMap) (*IpAddressNicMap, error)
	Update(existing *IpAddressNicMap, updates interface{}) (*IpAddressNicMap, error)
	ById(id string) (*IpAddressNicMap, error)
	Delete(container *IpAddressNicMap) error
}

func newIpAddressNicMapClient(rancherClient *RancherClient) *IpAddressNicMapClient {
	return &IpAddressNicMapClient{
		rancherClient: rancherClient,
	}
}

func (c *IpAddressNicMapClient) Create(container *IpAddressNicMap) (*IpAddressNicMap, error) {
	resp := &IpAddressNicMap{}
	err := c.rancherClient.doCreate(IP_ADDRESS_NIC_MAP_TYPE, container, resp)
	return resp, err
}

func (c *IpAddressNicMapClient) Update(existing *IpAddressNicMap, updates interface{}) (*IpAddressNicMap, error) {
	resp := &IpAddressNicMap{}
	err := c.rancherClient.doUpdate(IP_ADDRESS_NIC_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *IpAddressNicMapClient) List(opts *ListOpts) (*IpAddressNicMapCollection, error) {
	resp := &IpAddressNicMapCollection{}
	err := c.rancherClient.doList(IP_ADDRESS_NIC_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *IpAddressNicMapClient) ById(id string) (*IpAddressNicMap, error) {
	resp := &IpAddressNicMap{}
	err := c.rancherClient.doById(IP_ADDRESS_NIC_MAP_TYPE, id, resp)
	return resp, err
}

func (c *IpAddressNicMapClient) Delete(container *IpAddressNicMap) error {
	return c.rancherClient.doResourceDelete(IP_ADDRESS_NIC_MAP_TYPE, &container.Resource)
}
