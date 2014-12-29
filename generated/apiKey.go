package client

const (
	CONTAINER_TYPE = "apiKey"
)

type ApiKey struct {
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

type ApiKeyCollection struct {
	Collection
	Data []ApiKey `json:"data,omitempty"`
}

type ApiKeyClient struct {
	rancherClient *RancherClient
}

type ApiKeyOperations interface {
	List(opts *ListOpts) (*ApiKeyCollection, error)
	Create(opts *ApiKey) (*ApiKey, error)
	Update(existing *ApiKey, updates interface{}) (*ApiKey, error)
	ById(id string) (*ApiKey, error)
	Delete(container *ApiKey) error
}

func newApiKeyClient(rancherClient *RancherClient) *ApiKeyClient {
	return &ApiKeyClient{
		rancherClient: rancherClient,
	}
}

func (self *ApiKeyClient) Create(container *ApiKey) (*ApiKey, error) {
	resp := &ApiKey{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ApiKeyClient) Update(existing *ApiKey, updates interface{}) (*ApiKey, error) {
	resp := &ApiKey{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ApiKeyClient) List(opts *ListOpts) (*ApiKeyCollection, error) {
	resp := &ApiKeyCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ApiKeyClient) ById(id string) (*ApiKey, error) {
	resp := &ApiKey{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ApiKeyClient) Delete(container *ApiKey) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
