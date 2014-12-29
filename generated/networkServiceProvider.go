package client

const (
	CONTAINER_TYPE = "networkServiceProvider"
)

type NetworkServiceProvider struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkId string `json:"NetworkId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type NetworkServiceProviderCollection struct {
	Collection
	Data []NetworkServiceProvider `json:"data,omitempty"`
}

type NetworkServiceProviderClient struct {
	rancherClient *RancherClient
}

type NetworkServiceProviderOperations interface {
	List(opts *ListOpts) (*NetworkServiceProviderCollection, error)
	Create(opts *NetworkServiceProvider) (*NetworkServiceProvider, error)
	Update(existing *NetworkServiceProvider, updates interface{}) (*NetworkServiceProvider, error)
	ById(id string) (*NetworkServiceProvider, error)
	Delete(container *NetworkServiceProvider) error
}

func newNetworkServiceProviderClient(rancherClient *RancherClient) *NetworkServiceProviderClient {
	return &NetworkServiceProviderClient{
		rancherClient: rancherClient,
	}
}

func (self *NetworkServiceProviderClient) Create(container *NetworkServiceProvider) (*NetworkServiceProvider, error) {
	resp := &NetworkServiceProvider{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *NetworkServiceProviderClient) Update(existing *NetworkServiceProvider, updates interface{}) (*NetworkServiceProvider, error) {
	resp := &NetworkServiceProvider{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *NetworkServiceProviderClient) List(opts *ListOpts) (*NetworkServiceProviderCollection, error) {
	resp := &NetworkServiceProviderCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *NetworkServiceProviderClient) ById(id string) (*NetworkServiceProvider, error) {
	resp := &NetworkServiceProvider{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *NetworkServiceProviderClient) Delete(container *NetworkServiceProvider) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
