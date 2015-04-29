package client

const (
	LOAD_BALANCER_CONFIG_LISTENER_MAP_TYPE = "loadBalancerConfigListenerMap"
)

type LoadBalancerConfigListenerMap struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    LoadBalancerConfigId string `json:"loadBalancerConfigId,omitempty"`
    
    LoadBalancerListenerId string `json:"loadBalancerListenerId,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type LoadBalancerConfigListenerMapCollection struct {
	Collection
	Data []LoadBalancerConfigListenerMap `json:"data,omitempty"`
}

type LoadBalancerConfigListenerMapClient struct {
	rancherClient *RancherClient
}

type LoadBalancerConfigListenerMapOperations interface {
	List(opts *ListOpts) (*LoadBalancerConfigListenerMapCollection, error)
	Create(opts *LoadBalancerConfigListenerMap) (*LoadBalancerConfigListenerMap, error)
	Update(existing *LoadBalancerConfigListenerMap, updates interface{}) (*LoadBalancerConfigListenerMap, error)
	ById(id string) (*LoadBalancerConfigListenerMap, error)
	Delete(container *LoadBalancerConfigListenerMap) error
    ActionCreate (*LoadBalancerConfigListenerMap) (*LoadBalancerConfigListenerMap, error)
    ActionRemove (*LoadBalancerConfigListenerMap) (*LoadBalancerConfigListenerMap, error)
}

func newLoadBalancerConfigListenerMapClient(rancherClient *RancherClient) *LoadBalancerConfigListenerMapClient {
	return &LoadBalancerConfigListenerMapClient{
		rancherClient: rancherClient,
	}
}

func (c *LoadBalancerConfigListenerMapClient) Create(container *LoadBalancerConfigListenerMap) (*LoadBalancerConfigListenerMap, error) {
	resp := &LoadBalancerConfigListenerMap{}
	err := c.rancherClient.doCreate(LOAD_BALANCER_CONFIG_LISTENER_MAP_TYPE, container, resp)
	return resp, err
}

func (c *LoadBalancerConfigListenerMapClient) Update(existing *LoadBalancerConfigListenerMap, updates interface{}) (*LoadBalancerConfigListenerMap, error) {
	resp := &LoadBalancerConfigListenerMap{}
	err := c.rancherClient.doUpdate(LOAD_BALANCER_CONFIG_LISTENER_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LoadBalancerConfigListenerMapClient) List(opts *ListOpts) (*LoadBalancerConfigListenerMapCollection, error) {
	resp := &LoadBalancerConfigListenerMapCollection{}
	err := c.rancherClient.doList(LOAD_BALANCER_CONFIG_LISTENER_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *LoadBalancerConfigListenerMapClient) ById(id string) (*LoadBalancerConfigListenerMap, error) {
	resp := &LoadBalancerConfigListenerMap{}
	err := c.rancherClient.doById(LOAD_BALANCER_CONFIG_LISTENER_MAP_TYPE, id, resp)
	return resp, err
}

func (c *LoadBalancerConfigListenerMapClient) Delete(container *LoadBalancerConfigListenerMap) error {
	return c.rancherClient.doResourceDelete(LOAD_BALANCER_CONFIG_LISTENER_MAP_TYPE, &container.Resource)
}

func (c *LoadBalancerConfigListenerMapClient) ActionCreate(resource *LoadBalancerConfigListenerMap) (*LoadBalancerConfigListenerMap, error) {
	resp := &LoadBalancerConfigListenerMap{}
	err := c.rancherClient.doEmptyAction(LOAD_BALANCER_CONFIG_LISTENER_MAP_TYPE, "create", &resource.Resource, resp)
	return resp, err
}

func (c *LoadBalancerConfigListenerMapClient) ActionRemove(resource *LoadBalancerConfigListenerMap) (*LoadBalancerConfigListenerMap, error) {
	resp := &LoadBalancerConfigListenerMap{}
	err := c.rancherClient.doEmptyAction(LOAD_BALANCER_CONFIG_LISTENER_MAP_TYPE, "remove", &resource.Resource, resp)
	return resp, err
}
