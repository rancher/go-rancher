package client

const (
	REGISTRATION_TOKEN_TYPE = "registrationToken"
)

type RegistrationToken struct {
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

type RegistrationTokenCollection struct {
	Collection
	Data []RegistrationToken `json:"data,omitempty"`
}

type RegistrationTokenClient struct {
	rancherClient *RancherClient
}

type RegistrationTokenOperations interface {
	List(opts *ListOpts) (*RegistrationTokenCollection, error)
	Create(opts *RegistrationToken) (*RegistrationToken, error)
	Update(existing *RegistrationToken, updates interface{}) (*RegistrationToken, error)
	ById(id string) (*RegistrationToken, error)
	Delete(container *RegistrationToken) error
}

func newRegistrationTokenClient(rancherClient *RancherClient) *RegistrationTokenClient {
	return &RegistrationTokenClient{
		rancherClient: rancherClient,
	}
}

func (c *RegistrationTokenClient) Create(container *RegistrationToken) (*RegistrationToken, error) {
	resp := &RegistrationToken{}
	err := c.rancherClient.doCreate(REGISTRATION_TOKEN_TYPE, container, resp)
	return resp, err
}

func (c *RegistrationTokenClient) Update(existing *RegistrationToken, updates interface{}) (*RegistrationToken, error) {
	resp := &RegistrationToken{}
	err := c.rancherClient.doUpdate(REGISTRATION_TOKEN_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RegistrationTokenClient) List(opts *ListOpts) (*RegistrationTokenCollection, error) {
	resp := &RegistrationTokenCollection{}
	err := c.rancherClient.doList(REGISTRATION_TOKEN_TYPE, opts, resp)
	return resp, err
}

func (c *RegistrationTokenClient) ById(id string) (*RegistrationToken, error) {
	resp := &RegistrationToken{}
	err := c.rancherClient.doById(REGISTRATION_TOKEN_TYPE, id, resp)
	return resp, err
}

func (c *RegistrationTokenClient) Delete(container *RegistrationToken) error {
	return c.rancherClient.doResourceDelete(REGISTRATION_TOKEN_TYPE, &container.Resource)
}
