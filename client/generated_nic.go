package client

const (
	NIC_TYPE = "nic"
)

type Nic struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    DeviceNumber int `json:"deviceNumber,omitempty"`
    
    InstanceId string `json:"instanceId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    MacAddress string `json:"macAddress,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkId string `json:"networkId,omitempty"`
    
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

type NicCollection struct {
	Collection
	Data []Nic `json:"data,omitempty"`
}

type NicClient struct {
	rancherClient *RancherClient
}

type NicOperations interface {
	List(opts *ListOpts) (*NicCollection, error)
	Create(opts *Nic) (*Nic, error)
	Update(existing *Nic, updates interface{}) (*Nic, error)
	ById(id string) (*Nic, error)
	Delete(container *Nic) error
}

func newNicClient(rancherClient *RancherClient) *NicClient {
	return &NicClient{
		rancherClient: rancherClient,
	}
}

func (c *NicClient) Create(container *Nic) (*Nic, error) {
	resp := &Nic{}
	err := c.rancherClient.doCreate(NIC_TYPE, container, resp)
	return resp, err
}

func (c *NicClient) Update(existing *Nic, updates interface{}) (*Nic, error) {
	resp := &Nic{}
	err := c.rancherClient.doUpdate(NIC_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NicClient) List(opts *ListOpts) (*NicCollection, error) {
	resp := &NicCollection{}
	err := c.rancherClient.doList(NIC_TYPE, opts, resp)
	return resp, err
}

func (c *NicClient) ById(id string) (*Nic, error) {
	resp := &Nic{}
	err := c.rancherClient.doById(NIC_TYPE, id, resp)
	return resp, err
}

func (c *NicClient) Delete(container *Nic) error {
	return c.rancherClient.doResourceDelete(NIC_TYPE, &container.Resource)
}
