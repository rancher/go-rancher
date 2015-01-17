package client

const (
	DHCP_SERVICE_TYPE = "dhcpService"
)

type DhcpService struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkId string `json:"networkId,omitempty"`
    
    NetworkServiceProviderId string `json:"networkServiceProviderId,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type DhcpServiceCollection struct {
	Collection
	Data []DhcpService `json:"data,omitempty"`
}

type DhcpServiceClient struct {
	rancherClient *RancherClient
}

type DhcpServiceOperations interface {
	List(opts *ListOpts) (*DhcpServiceCollection, error)
	Create(opts *DhcpService) (*DhcpService, error)
	Update(existing *DhcpService, updates interface{}) (*DhcpService, error)
	ById(id string) (*DhcpService, error)
	Delete(container *DhcpService) error
}

func newDhcpServiceClient(rancherClient *RancherClient) *DhcpServiceClient {
	return &DhcpServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *DhcpServiceClient) Create(container *DhcpService) (*DhcpService, error) {
	resp := &DhcpService{}
	err := c.rancherClient.doCreate(DHCP_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *DhcpServiceClient) Update(existing *DhcpService, updates interface{}) (*DhcpService, error) {
	resp := &DhcpService{}
	err := c.rancherClient.doUpdate(DHCP_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DhcpServiceClient) List(opts *ListOpts) (*DhcpServiceCollection, error) {
	resp := &DhcpServiceCollection{}
	err := c.rancherClient.doList(DHCP_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *DhcpServiceClient) ById(id string) (*DhcpService, error) {
	resp := &DhcpService{}
	err := c.rancherClient.doById(DHCP_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *DhcpServiceClient) Delete(container *DhcpService) error {
	return c.rancherClient.doResourceDelete(DHCP_SERVICE_TYPE, &container.Resource)
}
