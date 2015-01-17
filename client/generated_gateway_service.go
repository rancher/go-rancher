package client

const (
	GATEWAY_SERVICE_TYPE = "gatewayService"
)

type GatewayService struct {
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

type GatewayServiceCollection struct {
	Collection
	Data []GatewayService `json:"data,omitempty"`
}

type GatewayServiceClient struct {
	rancherClient *RancherClient
}

type GatewayServiceOperations interface {
	List(opts *ListOpts) (*GatewayServiceCollection, error)
	Create(opts *GatewayService) (*GatewayService, error)
	Update(existing *GatewayService, updates interface{}) (*GatewayService, error)
	ById(id string) (*GatewayService, error)
	Delete(container *GatewayService) error
}

func newGatewayServiceClient(rancherClient *RancherClient) *GatewayServiceClient {
	return &GatewayServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *GatewayServiceClient) Create(container *GatewayService) (*GatewayService, error) {
	resp := &GatewayService{}
	err := c.rancherClient.doCreate(GATEWAY_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *GatewayServiceClient) Update(existing *GatewayService, updates interface{}) (*GatewayService, error) {
	resp := &GatewayService{}
	err := c.rancherClient.doUpdate(GATEWAY_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GatewayServiceClient) List(opts *ListOpts) (*GatewayServiceCollection, error) {
	resp := &GatewayServiceCollection{}
	err := c.rancherClient.doList(GATEWAY_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *GatewayServiceClient) ById(id string) (*GatewayService, error) {
	resp := &GatewayService{}
	err := c.rancherClient.doById(GATEWAY_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *GatewayServiceClient) Delete(container *GatewayService) error {
	return c.rancherClient.doResourceDelete(GATEWAY_SERVICE_TYPE, &container.Resource)
}
