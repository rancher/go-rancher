package client

const (
	CONTAINER_TYPE = "networkService"
)

type NetworkService struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkId string `json:"NetworkId,omitempty"`
    
    NetworkServiceProviderId string `json:"NetworkServiceProviderId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type NetworkServiceCollection struct {
	Collection
	Data []NetworkService `json:"data,omitempty"`
}

type NetworkServiceClient struct {
	rancherClient *RancherClient
}

type NetworkServiceOperations interface {
	List(opts *ListOpts) (*NetworkServiceCollection, error)
	Create(opts *NetworkService) (*NetworkService, error)
	Update(existing *NetworkService, updates interface{}) (*NetworkService, error)
	ById(id string) (*NetworkService, error)
	Delete(container *NetworkService) error
}

func newNetworkServiceClient(rancherClient *RancherClient) *NetworkServiceClient {
	return &NetworkServiceClient{
		rancherClient: rancherClient,
	}
}

func (self *NetworkServiceClient) Create(container *NetworkService) (*NetworkService, error) {
	resp := &NetworkService{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *NetworkServiceClient) Update(existing *NetworkService, updates interface{}) (*NetworkService, error) {
	resp := &NetworkService{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *NetworkServiceClient) List(opts *ListOpts) (*NetworkServiceCollection, error) {
	resp := &NetworkServiceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *NetworkServiceClient) ById(id string) (*NetworkService, error) {
	resp := &NetworkService{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *NetworkServiceClient) Delete(container *NetworkService) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
