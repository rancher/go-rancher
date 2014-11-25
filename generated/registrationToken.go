package client

const (
	CONTAINER_TYPE = "registrationToken"
)

type RegistrationToken struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
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

func (self *RegistrationTokenClient) Create(container *RegistrationToken) (*RegistrationToken, error) {
	resp := &RegistrationToken{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *RegistrationTokenClient) Update(existing *RegistrationToken, updates interface{}) (*RegistrationToken, error) {
	resp := &RegistrationToken{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *RegistrationTokenClient) List(opts *ListOpts) (*RegistrationTokenCollection, error) {
	resp := &RegistrationTokenCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *RegistrationTokenClient) ById(id string) (*RegistrationToken, error) {
	resp := &RegistrationToken{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *RegistrationTokenClient) Delete(container *RegistrationToken) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
