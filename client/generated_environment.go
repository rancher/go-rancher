package client

const (
	ENVIRONMENT_TYPE = "environment"
)

type Environment struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
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

type EnvironmentCollection struct {
	Collection
	Data []Environment `json:"data,omitempty"`
}

type EnvironmentClient struct {
	rancherClient *RancherClient
}

type EnvironmentOperations interface {
	List(opts *ListOpts) (*EnvironmentCollection, error)
	Create(opts *Environment) (*Environment, error)
	Update(existing *Environment, updates interface{}) (*Environment, error)
	ById(id string) (*Environment, error)
	Delete(container *Environment) error
    
    ActionCreate (*Environment) (*Environment, error)
    
    
    ActionExportconfig (*Environment, *ComposeConfigInput) (*ComposeConfig, error)
    
    
    ActionRemove (*Environment) (*Environment, error)
    
}

func newEnvironmentClient(rancherClient *RancherClient) *EnvironmentClient {
	return &EnvironmentClient{
		rancherClient: rancherClient,
	}
}

func (c *EnvironmentClient) Create(container *Environment) (*Environment, error) {
	resp := &Environment{}
	err := c.rancherClient.doCreate(ENVIRONMENT_TYPE, container, resp)
	return resp, err
}

func (c *EnvironmentClient) Update(existing *Environment, updates interface{}) (*Environment, error) {
	resp := &Environment{}
	err := c.rancherClient.doUpdate(ENVIRONMENT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *EnvironmentClient) List(opts *ListOpts) (*EnvironmentCollection, error) {
	resp := &EnvironmentCollection{}
	err := c.rancherClient.doList(ENVIRONMENT_TYPE, opts, resp)
	return resp, err
}

func (c *EnvironmentClient) ById(id string) (*Environment, error) {
	resp := &Environment{}
	err := c.rancherClient.doById(ENVIRONMENT_TYPE, id, resp)
	return resp, err
}

func (c *EnvironmentClient) Delete(container *Environment) error {
	return c.rancherClient.doResourceDelete(ENVIRONMENT_TYPE, &container.Resource)
}
    
func (c *EnvironmentClient) ActionCreate (resource *Environment) (*Environment, error) {
    
	resp := &Environment{}
    
	err := c.rancherClient.doAction(ENVIRONMENT_TYPE, "create", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *EnvironmentClient) ActionExportconfig (resource *Environment, input *ComposeConfigInput) (*ComposeConfig, error) {
    
	resp := &ComposeConfig{}
    
	err := c.rancherClient.doAction(ENVIRONMENT_TYPE, "exportconfig", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *EnvironmentClient) ActionRemove (resource *Environment) (*Environment, error) {
    
	resp := &Environment{}
    
	err := c.rancherClient.doAction(ENVIRONMENT_TYPE, "remove", &resource.Resource, nil, resp)
    
	return resp, err
}
