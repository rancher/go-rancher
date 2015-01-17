package client

const (
	CREDENTIAL_INSTANCE_MAP_TYPE = "credentialInstanceMap"
)

type CredentialInstanceMap struct {
	Resource
    
    Created string `json:"created,omitempty"`
    
    CredentialId string `json:"credentialId,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    InstanceId string `json:"instanceId,omitempty"`
    
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

type CredentialInstanceMapCollection struct {
	Collection
	Data []CredentialInstanceMap `json:"data,omitempty"`
}

type CredentialInstanceMapClient struct {
	rancherClient *RancherClient
}

type CredentialInstanceMapOperations interface {
	List(opts *ListOpts) (*CredentialInstanceMapCollection, error)
	Create(opts *CredentialInstanceMap) (*CredentialInstanceMap, error)
	Update(existing *CredentialInstanceMap, updates interface{}) (*CredentialInstanceMap, error)
	ById(id string) (*CredentialInstanceMap, error)
	Delete(container *CredentialInstanceMap) error
}

func newCredentialInstanceMapClient(rancherClient *RancherClient) *CredentialInstanceMapClient {
	return &CredentialInstanceMapClient{
		rancherClient: rancherClient,
	}
}

func (c *CredentialInstanceMapClient) Create(container *CredentialInstanceMap) (*CredentialInstanceMap, error) {
	resp := &CredentialInstanceMap{}
	err := c.rancherClient.doCreate(CREDENTIAL_INSTANCE_MAP_TYPE, container, resp)
	return resp, err
}

func (c *CredentialInstanceMapClient) Update(existing *CredentialInstanceMap, updates interface{}) (*CredentialInstanceMap, error) {
	resp := &CredentialInstanceMap{}
	err := c.rancherClient.doUpdate(CREDENTIAL_INSTANCE_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *CredentialInstanceMapClient) List(opts *ListOpts) (*CredentialInstanceMapCollection, error) {
	resp := &CredentialInstanceMapCollection{}
	err := c.rancherClient.doList(CREDENTIAL_INSTANCE_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *CredentialInstanceMapClient) ById(id string) (*CredentialInstanceMap, error) {
	resp := &CredentialInstanceMap{}
	err := c.rancherClient.doById(CREDENTIAL_INSTANCE_MAP_TYPE, id, resp)
	return resp, err
}

func (c *CredentialInstanceMapClient) Delete(container *CredentialInstanceMap) error {
	return c.rancherClient.doResourceDelete(CREDENTIAL_INSTANCE_MAP_TYPE, &container.Resource)
}
