package client

const (
	NETWORK_SERVICE_TYPE = "networkService"
)

type NetworkService struct {
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

type NetworkServiceCollection struct {
	Collection
	Data []NetworkService `json:"data,omitempty"`
}

type NetworkServiceClient struct {
	rancherClient *RancherClient
}

type NetworkServiceOperations interface {
	List(opts *ListOpts) (*NetworkServiceCollection, error)
	Create(opts *NetworkService) (*NetworkService, error)
	Update(existing *NetworkService, updates interface{}) (*NetworkService, error)
	ById(id string) (*NetworkService, error)
	Delete(container *NetworkService) error
}

func newNetworkServiceClient(rancherClient *RancherClient) *NetworkServiceClient {
	return &NetworkServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *NetworkServiceClient) Create(container *NetworkService) (*NetworkService, error) {
	resp := &NetworkService{}
	err := c.rancherClient.doCreate(NETWORK_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *NetworkServiceClient) Update(existing *NetworkService, updates interface{}) (*NetworkService, error) {
	resp := &NetworkService{}
	err := c.rancherClient.doUpdate(NETWORK_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NetworkServiceClient) List(opts *ListOpts) (*NetworkServiceCollection, error) {
	resp := &NetworkServiceCollection{}
	err := c.rancherClient.doList(NETWORK_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *NetworkServiceClient) ById(id string) (*NetworkService, error) {
	resp := &NetworkService{}
	err := c.rancherClient.doById(NETWORK_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *NetworkServiceClient) Delete(container *NetworkService) error {
	return c.rancherClient.doResourceDelete(NETWORK_SERVICE_TYPE, &container.Resource)
}
