package client

const (
	HOST_NAT_GATEWAY_SERVICE_TYPE = "hostNatGatewayService"
)

type HostNatGatewayService struct {
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

type HostNatGatewayServiceCollection struct {
	Collection
	Data []HostNatGatewayService `json:"data,omitempty"`
}

type HostNatGatewayServiceClient struct {
	rancherClient *RancherClient
}

type HostNatGatewayServiceOperations interface {
	List(opts *ListOpts) (*HostNatGatewayServiceCollection, error)
	Create(opts *HostNatGatewayService) (*HostNatGatewayService, error)
	Update(existing *HostNatGatewayService, updates interface{}) (*HostNatGatewayService, error)
	ById(id string) (*HostNatGatewayService, error)
	Delete(container *HostNatGatewayService) error
}

func newHostNatGatewayServiceClient(rancherClient *RancherClient) *HostNatGatewayServiceClient {
	return &HostNatGatewayServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *HostNatGatewayServiceClient) Create(container *HostNatGatewayService) (*HostNatGatewayService, error) {
	resp := &HostNatGatewayService{}
	err := c.rancherClient.doCreate(HOST_NAT_GATEWAY_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *HostNatGatewayServiceClient) Update(existing *HostNatGatewayService, updates interface{}) (*HostNatGatewayService, error) {
	resp := &HostNatGatewayService{}
	err := c.rancherClient.doUpdate(HOST_NAT_GATEWAY_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HostNatGatewayServiceClient) List(opts *ListOpts) (*HostNatGatewayServiceCollection, error) {
	resp := &HostNatGatewayServiceCollection{}
	err := c.rancherClient.doList(HOST_NAT_GATEWAY_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *HostNatGatewayServiceClient) ById(id string) (*HostNatGatewayService, error) {
	resp := &HostNatGatewayService{}
	err := c.rancherClient.doById(HOST_NAT_GATEWAY_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *HostNatGatewayServiceClient) Delete(container *HostNatGatewayService) error {
	return c.rancherClient.doResourceDelete(HOST_NAT_GATEWAY_SERVICE_TYPE, &container.Resource)
}
