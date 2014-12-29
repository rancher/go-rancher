package client

const (
	CONTAINER_TYPE = "portService"
)

type PortService struct {
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

type PortServiceCollection struct {
	Collection
	Data []PortService `json:"data,omitempty"`
}

type PortServiceClient struct {
	rancherClient *RancherClient
}

type PortServiceOperations interface {
	List(opts *ListOpts) (*PortServiceCollection, error)
	Create(opts *PortService) (*PortService, error)
	Update(existing *PortService, updates interface{}) (*PortService, error)
	ById(id string) (*PortService, error)
	Delete(container *PortService) error
}

func newPortServiceClient(rancherClient *RancherClient) *PortServiceClient {
	return &PortServiceClient{
		rancherClient: rancherClient,
	}
}

func (self *PortServiceClient) Create(container *PortService) (*PortService, error) {
	resp := &PortService{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *PortServiceClient) Update(existing *PortService, updates interface{}) (*PortService, error) {
	resp := &PortService{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *PortServiceClient) List(opts *ListOpts) (*PortServiceCollection, error) {
	resp := &PortServiceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *PortServiceClient) ById(id string) (*PortService, error) {
	resp := &PortService{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *PortServiceClient) Delete(container *PortService) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
