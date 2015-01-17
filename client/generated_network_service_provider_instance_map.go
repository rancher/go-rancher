package client

const (
	NETWORK_SERVICE_PROVIDER_INSTANCE_MAP_TYPE = "networkServiceProviderInstanceMap"
)

type NetworkServiceProviderInstanceMap struct {
	Resource
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    InstanceId string `json:"instanceId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkServiceProviderId string `json:"networkServiceProviderId,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type NetworkServiceProviderInstanceMapCollection struct {
	Collection
	Data []NetworkServiceProviderInstanceMap `json:"data,omitempty"`
}

type NetworkServiceProviderInstanceMapClient struct {
	rancherClient *RancherClient
}

type NetworkServiceProviderInstanceMapOperations interface {
	List(opts *ListOpts) (*NetworkServiceProviderInstanceMapCollection, error)
	Create(opts *NetworkServiceProviderInstanceMap) (*NetworkServiceProviderInstanceMap, error)
	Update(existing *NetworkServiceProviderInstanceMap, updates interface{}) (*NetworkServiceProviderInstanceMap, error)
	ById(id string) (*NetworkServiceProviderInstanceMap, error)
	Delete(container *NetworkServiceProviderInstanceMap) error
}

func newNetworkServiceProviderInstanceMapClient(rancherClient *RancherClient) *NetworkServiceProviderInstanceMapClient {
	return &NetworkServiceProviderInstanceMapClient{
		rancherClient: rancherClient,
	}
}

func (c *NetworkServiceProviderInstanceMapClient) Create(container *NetworkServiceProviderInstanceMap) (*NetworkServiceProviderInstanceMap, error) {
	resp := &NetworkServiceProviderInstanceMap{}
	err := c.rancherClient.doCreate(NETWORK_SERVICE_PROVIDER_INSTANCE_MAP_TYPE, container, resp)
	return resp, err
}

func (c *NetworkServiceProviderInstanceMapClient) Update(existing *NetworkServiceProviderInstanceMap, updates interface{}) (*NetworkServiceProviderInstanceMap, error) {
	resp := &NetworkServiceProviderInstanceMap{}
	err := c.rancherClient.doUpdate(NETWORK_SERVICE_PROVIDER_INSTANCE_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NetworkServiceProviderInstanceMapClient) List(opts *ListOpts) (*NetworkServiceProviderInstanceMapCollection, error) {
	resp := &NetworkServiceProviderInstanceMapCollection{}
	err := c.rancherClient.doList(NETWORK_SERVICE_PROVIDER_INSTANCE_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *NetworkServiceProviderInstanceMapClient) ById(id string) (*NetworkServiceProviderInstanceMap, error) {
	resp := &NetworkServiceProviderInstanceMap{}
	err := c.rancherClient.doById(NETWORK_SERVICE_PROVIDER_INSTANCE_MAP_TYPE, id, resp)
	return resp, err
}

func (c *NetworkServiceProviderInstanceMapClient) Delete(container *NetworkServiceProviderInstanceMap) error {
	return c.rancherClient.doResourceDelete(NETWORK_SERVICE_PROVIDER_INSTANCE_MAP_TYPE, &container.Resource)
}
