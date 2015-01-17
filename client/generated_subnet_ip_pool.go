package client

const (
	SUBNET_IP_POOL_TYPE = "subnetIpPool"
)

type SubnetIpPool struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    CidrSize int `json:"cidrSize,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    EndAddress string `json:"endAddress,omitempty"`
    
    Gateway string `json:"gateway,omitempty"`
    
    IsPublic bool `json:"isPublic,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkAddress string `json:"networkAddress,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    StartAddress string `json:"startAddress,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type SubnetIpPoolCollection struct {
	Collection
	Data []SubnetIpPool `json:"data,omitempty"`
}

type SubnetIpPoolClient struct {
	rancherClient *RancherClient
}

type SubnetIpPoolOperations interface {
	List(opts *ListOpts) (*SubnetIpPoolCollection, error)
	Create(opts *SubnetIpPool) (*SubnetIpPool, error)
	Update(existing *SubnetIpPool, updates interface{}) (*SubnetIpPool, error)
	ById(id string) (*SubnetIpPool, error)
	Delete(container *SubnetIpPool) error
}

func newSubnetIpPoolClient(rancherClient *RancherClient) *SubnetIpPoolClient {
	return &SubnetIpPoolClient{
		rancherClient: rancherClient,
	}
}

func (c *SubnetIpPoolClient) Create(container *SubnetIpPool) (*SubnetIpPool, error) {
	resp := &SubnetIpPool{}
	err := c.rancherClient.doCreate(SUBNET_IP_POOL_TYPE, container, resp)
	return resp, err
}

func (c *SubnetIpPoolClient) Update(existing *SubnetIpPool, updates interface{}) (*SubnetIpPool, error) {
	resp := &SubnetIpPool{}
	err := c.rancherClient.doUpdate(SUBNET_IP_POOL_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SubnetIpPoolClient) List(opts *ListOpts) (*SubnetIpPoolCollection, error) {
	resp := &SubnetIpPoolCollection{}
	err := c.rancherClient.doList(SUBNET_IP_POOL_TYPE, opts, resp)
	return resp, err
}

func (c *SubnetIpPoolClient) ById(id string) (*SubnetIpPool, error) {
	resp := &SubnetIpPool{}
	err := c.rancherClient.doById(SUBNET_IP_POOL_TYPE, id, resp)
	return resp, err
}

func (c *SubnetIpPoolClient) Delete(container *SubnetIpPool) error {
	return c.rancherClient.doResourceDelete(SUBNET_IP_POOL_TYPE, &container.Resource)
}
