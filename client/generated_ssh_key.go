package client

const (
	SSH_KEY_TYPE = "sshKey"
)

type SshKey struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    PublicValue string `json:"publicValue,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    SecretValue string `json:"secretValue,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type SshKeyCollection struct {
	Collection
	Data []SshKey `json:"data,omitempty"`
}

type SshKeyClient struct {
	rancherClient *RancherClient
}

type SshKeyOperations interface {
	List(opts *ListOpts) (*SshKeyCollection, error)
	Create(opts *SshKey) (*SshKey, error)
	Update(existing *SshKey, updates interface{}) (*SshKey, error)
	ById(id string) (*SshKey, error)
	Delete(container *SshKey) error
}

func newSshKeyClient(rancherClient *RancherClient) *SshKeyClient {
	return &SshKeyClient{
		rancherClient: rancherClient,
	}
}

func (c *SshKeyClient) Create(container *SshKey) (*SshKey, error) {
	resp := &SshKey{}
	err := c.rancherClient.doCreate(SSH_KEY_TYPE, container, resp)
	return resp, err
}

func (c *SshKeyClient) Update(existing *SshKey, updates interface{}) (*SshKey, error) {
	resp := &SshKey{}
	err := c.rancherClient.doUpdate(SSH_KEY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SshKeyClient) List(opts *ListOpts) (*SshKeyCollection, error) {
	resp := &SshKeyCollection{}
	err := c.rancherClient.doList(SSH_KEY_TYPE, opts, resp)
	return resp, err
}

func (c *SshKeyClient) ById(id string) (*SshKey, error) {
	resp := &SshKey{}
	err := c.rancherClient.doById(SSH_KEY_TYPE, id, resp)
	return resp, err
}

func (c *SshKeyClient) Delete(container *SshKey) error {
	return c.rancherClient.doResourceDelete(SSH_KEY_TYPE, &container.Resource)
}
