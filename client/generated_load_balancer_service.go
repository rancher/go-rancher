package client

const (
	LOAD_BALANCER_SERVICE_TYPE = "loadBalancerService"
)

type LoadBalancerService struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    DataVolumesFromService []string `json:"dataVolumesFromService,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    EnvironmentId string `json:"environmentId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    LaunchConfig Container `json:"launchConfig,omitempty"`
    
    LoadBalancerConfig LoadBalancerConfig `json:"loadBalancerConfig,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    Scale int `json:"scale,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type LoadBalancerServiceCollection struct {
	Collection
	Data []LoadBalancerService `json:"data,omitempty"`
}

type LoadBalancerServiceClient struct {
	rancherClient *RancherClient
}

type LoadBalancerServiceOperations interface {
	List(opts *ListOpts) (*LoadBalancerServiceCollection, error)
	Create(opts *LoadBalancerService) (*LoadBalancerService, error)
	Update(existing *LoadBalancerService, updates interface{}) (*LoadBalancerService, error)
	ById(id string) (*LoadBalancerService, error)
	Delete(container *LoadBalancerService) error
}

func newLoadBalancerServiceClient(rancherClient *RancherClient) *LoadBalancerServiceClient {
	return &LoadBalancerServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *LoadBalancerServiceClient) Create(container *LoadBalancerService) (*LoadBalancerService, error) {
	resp := &LoadBalancerService{}
	err := c.rancherClient.doCreate(LOAD_BALANCER_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *LoadBalancerServiceClient) Update(existing *LoadBalancerService, updates interface{}) (*LoadBalancerService, error) {
	resp := &LoadBalancerService{}
	err := c.rancherClient.doUpdate(LOAD_BALANCER_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LoadBalancerServiceClient) List(opts *ListOpts) (*LoadBalancerServiceCollection, error) {
	resp := &LoadBalancerServiceCollection{}
	err := c.rancherClient.doList(LOAD_BALANCER_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *LoadBalancerServiceClient) ById(id string) (*LoadBalancerService, error) {
	resp := &LoadBalancerService{}
	err := c.rancherClient.doById(LOAD_BALANCER_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *LoadBalancerServiceClient) Delete(container *LoadBalancerService) error {
	return c.rancherClient.doResourceDelete(LOAD_BALANCER_SERVICE_TYPE, &container.Resource)
}
