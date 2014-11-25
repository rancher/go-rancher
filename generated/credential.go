package client

const (
	CONTAINER_TYPE = "credential"
)

type Credential struct {
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

type CredentialCollection struct {
	Collection
	Data []Credential `json:"data,omitempty"`
}

type CredentialClient struct {
	rancherClient *RancherClient
}

type CredentialOperations interface {
	List(opts *ListOpts) (*CredentialCollection, error)
	Create(opts *Credential) (*Credential, error)
	Update(existing *Credential, updates interface{}) (*Credential, error)
	ById(id string) (*Credential, error)
	Delete(container *Credential) error
}

func newCredentialClient(rancherClient *RancherClient) *CredentialClient {
	return &CredentialClient{
		rancherClient: rancherClient,
	}
}

func (self *CredentialClient) Create(container *Credential) (*Credential, error) {
	resp := &Credential{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *CredentialClient) Update(existing *Credential, updates interface{}) (*Credential, error) {
	resp := &Credential{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *CredentialClient) List(opts *ListOpts) (*CredentialCollection, error) {
	resp := &CredentialCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *CredentialClient) ById(id string) (*Credential, error) {
	resp := &Credential{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *CredentialClient) Delete(container *Credential) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
