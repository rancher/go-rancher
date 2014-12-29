package client

const (
	CONTAINER_TYPE = "account"
)

type Account struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    DefaultCredentialIds array[reference[credential]] `json:"DefaultCredentialIds,omitempty"`
    
    DefaultNetworkIds array[reference[network]] `json:"DefaultNetworkIds,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    ExternalId string `json:"ExternalId,omitempty"`
    
    ExternalIdType string `json:"ExternalIdType,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type AccountCollection struct {
	Collection
	Data []Account `json:"data,omitempty"`
}

type AccountClient struct {
	rancherClient *RancherClient
}

type AccountOperations interface {
	List(opts *ListOpts) (*AccountCollection, error)
	Create(opts *Account) (*Account, error)
	Update(existing *Account, updates interface{}) (*Account, error)
	ById(id string) (*Account, error)
	Delete(container *Account) error
}

func newAccountClient(rancherClient *RancherClient) *AccountClient {
	return &AccountClient{
		rancherClient: rancherClient,
	}
}

func (self *AccountClient) Create(container *Account) (*Account, error) {
	resp := &Account{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *AccountClient) Update(existing *Account, updates interface{}) (*Account, error) {
	resp := &Account{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *AccountClient) List(opts *ListOpts) (*AccountCollection, error) {
	resp := &AccountCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *AccountClient) ById(id string) (*Account, error) {
	resp := &Account{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *AccountClient) Delete(container *Account) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
