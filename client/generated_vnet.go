package client

const (
	VNET_TYPE = "vnet"
)

type Vnet struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkId string `json:"networkId,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uri string `json:"uri,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type VnetCollection struct {
	Collection
	Data []Vnet `json:"data,omitempty"`
}

type VnetClient struct {
	rancherClient *RancherClient
}

type VnetOperations interface {
	List(opts *ListOpts) (*VnetCollection, error)
	Create(opts *Vnet) (*Vnet, error)
	Update(existing *Vnet, updates interface{}) (*Vnet, error)
	ById(id string) (*Vnet, error)
	Delete(container *Vnet) error
}

func newVnetClient(rancherClient *RancherClient) *VnetClient {
	return &VnetClient{
		rancherClient: rancherClient,
	}
}

func (c *VnetClient) Create(container *Vnet) (*Vnet, error) {
	resp := &Vnet{}
	err := c.rancherClient.doCreate(VNET_TYPE, container, resp)
	return resp, err
}

func (c *VnetClient) Update(existing *Vnet, updates interface{}) (*Vnet, error) {
	resp := &Vnet{}
	err := c.rancherClient.doUpdate(VNET_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VnetClient) List(opts *ListOpts) (*VnetCollection, error) {
	resp := &VnetCollection{}
	err := c.rancherClient.doList(VNET_TYPE, opts, resp)
	return resp, err
}

func (c *VnetClient) ById(id string) (*Vnet, error) {
	resp := &Vnet{}
	err := c.rancherClient.doById(VNET_TYPE, id, resp)
	return resp, err
}

func (c *VnetClient) Delete(container *Vnet) error {
	return c.rancherClient.doResourceDelete(VNET_TYPE, &container.Resource)
}
