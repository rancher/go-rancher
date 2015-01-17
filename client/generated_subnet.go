package client

const (
	SUBNET_TYPE = "subnet"
)

type Subnet struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    CidrSize int `json:"cidrSize,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    EndAddress string `json:"endAddress,omitempty"`
    
    Gateway string `json:"gateway,omitempty"`
    
    IpPoolId string `json:"ipPoolId,omitempty"`
    
    IsPublic bool `json:"isPublic,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkAddress string `json:"networkAddress,omitempty"`
    
    NetworkId string `json:"networkId,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    StartAddress string `json:"startAddress,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type SubnetCollection struct {
	Collection
	Data []Subnet `json:"data,omitempty"`
}

type SubnetClient struct {
	rancherClient *RancherClient
}

type SubnetOperations interface {
	List(opts *ListOpts) (*SubnetCollection, error)
	Create(opts *Subnet) (*Subnet, error)
	Update(existing *Subnet, updates interface{}) (*Subnet, error)
	ById(id string) (*Subnet, error)
	Delete(container *Subnet) error
}

func newSubnetClient(rancherClient *RancherClient) *SubnetClient {
	return &SubnetClient{
		rancherClient: rancherClient,
	}
}

func (c *SubnetClient) Create(container *Subnet) (*Subnet, error) {
	resp := &Subnet{}
	err := c.rancherClient.doCreate(SUBNET_TYPE, container, resp)
	return resp, err
}

func (c *SubnetClient) Update(existing *Subnet, updates interface{}) (*Subnet, error) {
	resp := &Subnet{}
	err := c.rancherClient.doUpdate(SUBNET_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SubnetClient) List(opts *ListOpts) (*SubnetCollection, error) {
	resp := &SubnetCollection{}
	err := c.rancherClient.doList(SUBNET_TYPE, opts, resp)
	return resp, err
}

func (c *SubnetClient) ById(id string) (*Subnet, error) {
	resp := &Subnet{}
	err := c.rancherClient.doById(SUBNET_TYPE, id, resp)
	return resp, err
}

func (c *SubnetClient) Delete(container *Subnet) error {
	return c.rancherClient.doResourceDelete(SUBNET_TYPE, &container.Resource)
}
