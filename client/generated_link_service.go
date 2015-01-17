package client

const (
	LINK_SERVICE_TYPE = "linkService"
)

type LinkService struct {
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

type LinkServiceCollection struct {
	Collection
	Data []LinkService `json:"data,omitempty"`
}

type LinkServiceClient struct {
	rancherClient *RancherClient
}

type LinkServiceOperations interface {
	List(opts *ListOpts) (*LinkServiceCollection, error)
	Create(opts *LinkService) (*LinkService, error)
	Update(existing *LinkService, updates interface{}) (*LinkService, error)
	ById(id string) (*LinkService, error)
	Delete(container *LinkService) error
}

func newLinkServiceClient(rancherClient *RancherClient) *LinkServiceClient {
	return &LinkServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *LinkServiceClient) Create(container *LinkService) (*LinkService, error) {
	resp := &LinkService{}
	err := c.rancherClient.doCreate(LINK_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *LinkServiceClient) Update(existing *LinkService, updates interface{}) (*LinkService, error) {
	resp := &LinkService{}
	err := c.rancherClient.doUpdate(LINK_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LinkServiceClient) List(opts *ListOpts) (*LinkServiceCollection, error) {
	resp := &LinkServiceCollection{}
	err := c.rancherClient.doList(LINK_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *LinkServiceClient) ById(id string) (*LinkService, error) {
	resp := &LinkService{}
	err := c.rancherClient.doById(LINK_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *LinkServiceClient) Delete(container *LinkService) error {
	return c.rancherClient.doResourceDelete(LINK_SERVICE_TYPE, &container.Resource)
}
