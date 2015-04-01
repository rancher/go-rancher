package client

const (
	SERVICE_TYPE = "service"
)

type Service struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    DataVolumesFromService []string `json:"dataVolumesFromService,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    EnvironmentId string `json:"environmentId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    LaunchConfig Container `json:"launchConfig,omitempty"`
    
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

type ServiceCollection struct {
	Collection
	Data []Service `json:"data,omitempty"`
}

type ServiceClient struct {
	rancherClient *RancherClient
}

type ServiceOperations interface {
	List(opts *ListOpts) (*ServiceCollection, error)
	Create(opts *Service) (*Service, error)
	Update(existing *Service, updates interface{}) (*Service, error)
	ById(id string) (*Service, error)
	Delete(container *Service) error
    ActionActivate (*Service) (*Service, error)
    ActionCreate (*Service) (*Service, error)
    ActionDeactivate (*Service) (*Service, error)
    ActionRemove (*Service) (*Service, error)
}

func newServiceClient(rancherClient *RancherClient) *ServiceClient {
	return &ServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *ServiceClient) Create(container *Service) (*Service, error) {
	resp := &Service{}
	err := c.rancherClient.doCreate(SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *ServiceClient) Update(existing *Service, updates interface{}) (*Service, error) {
	resp := &Service{}
	err := c.rancherClient.doUpdate(SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceClient) List(opts *ListOpts) (*ServiceCollection, error) {
	resp := &ServiceCollection{}
	err := c.rancherClient.doList(SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *ServiceClient) ById(id string) (*Service, error) {
	resp := &Service{}
	err := c.rancherClient.doById(SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *ServiceClient) Delete(container *Service) error {
	return c.rancherClient.doResourceDelete(SERVICE_TYPE, &container.Resource)
}

func (c *ServiceClient) ActionActivate(resource *Service) (*Service, error) {
	resp := &Service{}
	err := c.rancherClient.doEmptyAction(SERVICE_TYPE, "activate", &resource.Resource, resp)
	return resp, err
}

func (c *ServiceClient) ActionCreate(resource *Service) (*Service, error) {
	resp := &Service{}
	err := c.rancherClient.doEmptyAction(SERVICE_TYPE, "create", &resource.Resource, resp)
	return resp, err
}

func (c *ServiceClient) ActionDeactivate(resource *Service) (*Service, error) {
	resp := &Service{}
	err := c.rancherClient.doEmptyAction(SERVICE_TYPE, "deactivate", &resource.Resource, resp)
	return resp, err
}

func (c *ServiceClient) ActionRemove(resource *Service) (*Service, error) {
	resp := &Service{}
	err := c.rancherClient.doEmptyAction(SERVICE_TYPE, "remove", &resource.Resource, resp)
	return resp, err
}
