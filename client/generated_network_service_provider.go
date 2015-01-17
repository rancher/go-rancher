package client

const (
	NETWORK_SERVICE_PROVIDER_TYPE = "networkServiceProvider"
)

type NetworkServiceProvider struct {
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
    
    Uuid string `json:"uuid,omitempty"`
    
}

type NetworkServiceProviderCollection struct {
	Collection
	Data []NetworkServiceProvider `json:"data,omitempty"`
}

type NetworkServiceProviderClient struct {
	rancherClient *RancherClient
}

type NetworkServiceProviderOperations interface {
	List(opts *ListOpts) (*NetworkServiceProviderCollection, error)
	Create(opts *NetworkServiceProvider) (*NetworkServiceProvider, error)
	Update(existing *NetworkServiceProvider, updates interface{}) (*NetworkServiceProvider, error)
	ById(id string) (*NetworkServiceProvider, error)
	Delete(container *NetworkServiceProvider) error
}

func newNetworkServiceProviderClient(rancherClient *RancherClient) *NetworkServiceProviderClient {
	return &NetworkServiceProviderClient{
		rancherClient: rancherClient,
	}
}

func (c *NetworkServiceProviderClient) Create(container *NetworkServiceProvider) (*NetworkServiceProvider, error) {
	resp := &NetworkServiceProvider{}
	err := c.rancherClient.doCreate(NETWORK_SERVICE_PROVIDER_TYPE, container, resp)
	return resp, err
}

func (c *NetworkServiceProviderClient) Update(existing *NetworkServiceProvider, updates interface{}) (*NetworkServiceProvider, error) {
	resp := &NetworkServiceProvider{}
	err := c.rancherClient.doUpdate(NETWORK_SERVICE_PROVIDER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NetworkServiceProviderClient) List(opts *ListOpts) (*NetworkServiceProviderCollection, error) {
	resp := &NetworkServiceProviderCollection{}
	err := c.rancherClient.doList(NETWORK_SERVICE_PROVIDER_TYPE, opts, resp)
	return resp, err
}

func (c *NetworkServiceProviderClient) ById(id string) (*NetworkServiceProvider, error) {
	resp := &NetworkServiceProvider{}
	err := c.rancherClient.doById(NETWORK_SERVICE_PROVIDER_TYPE, id, resp)
	return resp, err
}

func (c *NetworkServiceProviderClient) Delete(container *NetworkServiceProvider) error {
	return c.rancherClient.doResourceDelete(NETWORK_SERVICE_PROVIDER_TYPE, &container.Resource)
}
