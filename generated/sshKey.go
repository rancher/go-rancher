package client

const (
	CONTAINER_TYPE = "sshKey"
)

type SshKey struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    PublicValue string `json:"PublicValue,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    SecretValue string `json:"SecretValue,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
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

func (self *SshKeyClient) Create(container *SshKey) (*SshKey, error) {
	resp := &SshKey{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *SshKeyClient) Update(existing *SshKey, updates interface{}) (*SshKey, error) {
	resp := &SshKey{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *SshKeyClient) List(opts *ListOpts) (*SshKeyCollection, error) {
	resp := &SshKeyCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *SshKeyClient) ById(id string) (*SshKey, error) {
	resp := &SshKey{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *SshKeyClient) Delete(container *SshKey) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
