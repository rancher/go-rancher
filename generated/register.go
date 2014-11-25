package client

const (
	CONTAINER_TYPE = "register"
)

type Register struct {
	Resource
    
    AccessKey string `json:"AccessKey,omitempty"`
    
    AccountId string `json:"AccountId,omitempty"`
    
    AuthToken string `json:"AuthToken,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Key string `json:"Key,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    SecretKey string `json:"SecretKey,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type RegisterCollection struct {
	Collection
	Data []Register `json:"data,omitempty"`
}

type RegisterClient struct {
	rancherClient *RancherClient
}

type RegisterOperations interface {
	List(opts *ListOpts) (*RegisterCollection, error)
	Create(opts *Register) (*Register, error)
	Update(existing *Register, updates interface{}) (*Register, error)
	ById(id string) (*Register, error)
	Delete(container *Register) error
}

func newRegisterClient(rancherClient *RancherClient) *RegisterClient {
	return &RegisterClient{
		rancherClient: rancherClient,
	}
}

func (self *RegisterClient) Create(container *Register) (*Register, error) {
	resp := &Register{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *RegisterClient) Update(existing *Register, updates interface{}) (*Register, error) {
	resp := &Register{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *RegisterClient) List(opts *ListOpts) (*RegisterCollection, error) {
	resp := &RegisterCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *RegisterClient) ById(id string) (*Register, error) {
	resp := &Register{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *RegisterClient) Delete(container *Register) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
