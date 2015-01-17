package client

const (
	IP_POOL_TYPE = "ipPool"
)

type IpPool struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    IsPublic bool `json:"isPublic,omitempty"`
    
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

type IpPoolCollection struct {
	Collection
	Data []IpPool `json:"data,omitempty"`
}

type IpPoolClient struct {
	rancherClient *RancherClient
}

type IpPoolOperations interface {
	List(opts *ListOpts) (*IpPoolCollection, error)
	Create(opts *IpPool) (*IpPool, error)
	Update(existing *IpPool, updates interface{}) (*IpPool, error)
	ById(id string) (*IpPool, error)
	Delete(container *IpPool) error
}

func newIpPoolClient(rancherClient *RancherClient) *IpPoolClient {
	return &IpPoolClient{
		rancherClient: rancherClient,
	}
}

func (c *IpPoolClient) Create(container *IpPool) (*IpPool, error) {
	resp := &IpPool{}
	err := c.rancherClient.doCreate(IP_POOL_TYPE, container, resp)
	return resp, err
}

func (c *IpPoolClient) Update(existing *IpPool, updates interface{}) (*IpPool, error) {
	resp := &IpPool{}
	err := c.rancherClient.doUpdate(IP_POOL_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *IpPoolClient) List(opts *ListOpts) (*IpPoolCollection, error) {
	resp := &IpPoolCollection{}
	err := c.rancherClient.doList(IP_POOL_TYPE, opts, resp)
	return resp, err
}

func (c *IpPoolClient) ById(id string) (*IpPool, error) {
	resp := &IpPool{}
	err := c.rancherClient.doById(IP_POOL_TYPE, id, resp)
	return resp, err
}

func (c *IpPoolClient) Delete(container *IpPool) error {
	return c.rancherClient.doResourceDelete(IP_POOL_TYPE, &container.Resource)
}
