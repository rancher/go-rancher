package client

const (
	CONTAINER_TYPE = "credentialInstanceMap"
)

type CredentialInstanceMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    CredentialId string `json:"CredentialId,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    InstanceId string `json:"InstanceId,omitempty"`
    
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

func (self *CredentialInstanceMapClient) Create(container *CredentialInstanceMap) (*CredentialInstanceMap, error) {
	resp := &CredentialInstanceMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *CredentialInstanceMapClient) Update(existing *CredentialInstanceMap, updates interface{}) (*CredentialInstanceMap, error) {
	resp := &CredentialInstanceMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *CredentialInstanceMapClient) List(opts *ListOpts) (*CredentialInstanceMapCollection, error) {
	resp := &CredentialInstanceMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *CredentialInstanceMapClient) ById(id string) (*CredentialInstanceMap, error) {
	resp := &CredentialInstanceMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *CredentialInstanceMapClient) Delete(container *CredentialInstanceMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
