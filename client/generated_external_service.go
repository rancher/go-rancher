package client

const (
	EXTERNAL_SERVICE_TYPE = "externalService"
)

type ExternalService struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    EnvironmentId string `json:"environmentId,omitempty"`
    
    ExternalIpAddresses []string `json:"externalIpAddresses,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type ExternalServiceCollection struct {
	Collection
	Data []ExternalService `json:"data,omitempty"`
}

type ExternalServiceClient struct {
	rancherClient *RancherClient
}

type ExternalServiceOperations interface {
	List(opts *ListOpts) (*ExternalServiceCollection, error)
	Create(opts *ExternalService) (*ExternalService, error)
	Update(existing *ExternalService, updates interface{}) (*ExternalService, error)
	ById(id string) (*ExternalService, error)
	Delete(container *ExternalService) error
}

func newExternalServiceClient(rancherClient *RancherClient) *ExternalServiceClient {
	return &ExternalServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *ExternalServiceClient) Create(container *ExternalService) (*ExternalService, error) {
	resp := &ExternalService{}
	err := c.rancherClient.doCreate(EXTERNAL_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *ExternalServiceClient) Update(existing *ExternalService, updates interface{}) (*ExternalService, error) {
	resp := &ExternalService{}
	err := c.rancherClient.doUpdate(EXTERNAL_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalServiceClient) List(opts *ListOpts) (*ExternalServiceCollection, error) {
	resp := &ExternalServiceCollection{}
	err := c.rancherClient.doList(EXTERNAL_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *ExternalServiceClient) ById(id string) (*ExternalService, error) {
	resp := &ExternalService{}
	err := c.rancherClient.doById(EXTERNAL_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *ExternalServiceClient) Delete(container *ExternalService) error {
	return c.rancherClient.doResourceDelete(EXTERNAL_SERVICE_TYPE, &container.Resource)
}
