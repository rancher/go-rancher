package client

const (
	IP_ADDRESS_TYPE = "ipAddress"
)

type IpAddress struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Address string `json:"address,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Hostname string `json:"hostname,omitempty"`
    
    IpPoolId string `json:"ipPoolId,omitempty"`
    
    IsPublic bool `json:"isPublic,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkId string `json:"networkId,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    Role string `json:"role,omitempty"`
    
    State string `json:"state,omitempty"`
    
    SubnetId string `json:"subnetId,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type IpAddressCollection struct {
	Collection
	Data []IpAddress `json:"data,omitempty"`
}

type IpAddressClient struct {
	rancherClient *RancherClient
}

type IpAddressOperations interface {
	List(opts *ListOpts) (*IpAddressCollection, error)
	Create(opts *IpAddress) (*IpAddress, error)
	Update(existing *IpAddress, updates interface{}) (*IpAddress, error)
	ById(id string) (*IpAddress, error)
	Delete(container *IpAddress) error
}

func newIpAddressClient(rancherClient *RancherClient) *IpAddressClient {
	return &IpAddressClient{
		rancherClient: rancherClient,
	}
}

func (c *IpAddressClient) Create(container *IpAddress) (*IpAddress, error) {
	resp := &IpAddress{}
	err := c.rancherClient.doCreate(IP_ADDRESS_TYPE, container, resp)
	return resp, err
}

func (c *IpAddressClient) Update(existing *IpAddress, updates interface{}) (*IpAddress, error) {
	resp := &IpAddress{}
	err := c.rancherClient.doUpdate(IP_ADDRESS_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *IpAddressClient) List(opts *ListOpts) (*IpAddressCollection, error) {
	resp := &IpAddressCollection{}
	err := c.rancherClient.doList(IP_ADDRESS_TYPE, opts, resp)
	return resp, err
}

func (c *IpAddressClient) ById(id string) (*IpAddress, error) {
	resp := &IpAddress{}
	err := c.rancherClient.doById(IP_ADDRESS_TYPE, id, resp)
	return resp, err
}

func (c *IpAddressClient) Delete(container *IpAddress) error {
	return c.rancherClient.doResourceDelete(IP_ADDRESS_TYPE, &container.Resource)
}
