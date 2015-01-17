package client

const (
	IPSEC_TUNNEL_SERVICE_TYPE = "ipsecTunnelService"
)

type IpsecTunnelService struct {
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

type IpsecTunnelServiceCollection struct {
	Collection
	Data []IpsecTunnelService `json:"data,omitempty"`
}

type IpsecTunnelServiceClient struct {
	rancherClient *RancherClient
}

type IpsecTunnelServiceOperations interface {
	List(opts *ListOpts) (*IpsecTunnelServiceCollection, error)
	Create(opts *IpsecTunnelService) (*IpsecTunnelService, error)
	Update(existing *IpsecTunnelService, updates interface{}) (*IpsecTunnelService, error)
	ById(id string) (*IpsecTunnelService, error)
	Delete(container *IpsecTunnelService) error
}

func newIpsecTunnelServiceClient(rancherClient *RancherClient) *IpsecTunnelServiceClient {
	return &IpsecTunnelServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *IpsecTunnelServiceClient) Create(container *IpsecTunnelService) (*IpsecTunnelService, error) {
	resp := &IpsecTunnelService{}
	err := c.rancherClient.doCreate(IPSEC_TUNNEL_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *IpsecTunnelServiceClient) Update(existing *IpsecTunnelService, updates interface{}) (*IpsecTunnelService, error) {
	resp := &IpsecTunnelService{}
	err := c.rancherClient.doUpdate(IPSEC_TUNNEL_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *IpsecTunnelServiceClient) List(opts *ListOpts) (*IpsecTunnelServiceCollection, error) {
	resp := &IpsecTunnelServiceCollection{}
	err := c.rancherClient.doList(IPSEC_TUNNEL_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *IpsecTunnelServiceClient) ById(id string) (*IpsecTunnelService, error) {
	resp := &IpsecTunnelService{}
	err := c.rancherClient.doById(IPSEC_TUNNEL_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *IpsecTunnelServiceClient) Delete(container *IpsecTunnelService) error {
	return c.rancherClient.doResourceDelete(IPSEC_TUNNEL_SERVICE_TYPE, &container.Resource)
}
