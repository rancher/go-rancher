package client

const (
	PORT_SERVICE_TYPE = "portService"
)

type PortService struct {
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

type PortServiceCollection struct {
	Collection
	Data []PortService `json:"data,omitempty"`
}

type PortServiceClient struct {
	rancherClient *RancherClient
}

type PortServiceOperations interface {
	List(opts *ListOpts) (*PortServiceCollection, error)
	Create(opts *PortService) (*PortService, error)
	Update(existing *PortService, updates interface{}) (*PortService, error)
	ById(id string) (*PortService, error)
	Delete(container *PortService) error
}

func newPortServiceClient(rancherClient *RancherClient) *PortServiceClient {
	return &PortServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *PortServiceClient) Create(container *PortService) (*PortService, error) {
	resp := &PortService{}
	err := c.rancherClient.doCreate(PORT_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *PortServiceClient) Update(existing *PortService, updates interface{}) (*PortService, error) {
	resp := &PortService{}
	err := c.rancherClient.doUpdate(PORT_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PortServiceClient) List(opts *ListOpts) (*PortServiceCollection, error) {
	resp := &PortServiceCollection{}
	err := c.rancherClient.doList(PORT_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *PortServiceClient) ById(id string) (*PortService, error) {
	resp := &PortService{}
	err := c.rancherClient.doById(PORT_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *PortServiceClient) Delete(container *PortService) error {
	return c.rancherClient.doResourceDelete(PORT_SERVICE_TYPE, &container.Resource)
}
