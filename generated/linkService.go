package client

const (
	CONTAINER_TYPE = "linkService"
)

type LinkService struct {
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

type LinkServiceCollection struct {
	Collection
	Data []LinkService `json:"data,omitempty"`
}

type LinkServiceClient struct {
	rancherClient *RancherClient
}

type LinkServiceOperations interface {
	List(opts *ListOpts) (*LinkServiceCollection, error)
	Create(opts *LinkService) (*LinkService, error)
	Update(existing *LinkService, updates interface{}) (*LinkService, error)
	ById(id string) (*LinkService, error)
	Delete(container *LinkService) error
}

func newLinkServiceClient(rancherClient *RancherClient) *LinkServiceClient {
	return &LinkServiceClient{
		rancherClient: rancherClient,
	}
}

func (self *LinkServiceClient) Create(container *LinkService) (*LinkService, error) {
	resp := &LinkService{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *LinkServiceClient) Update(existing *LinkService, updates interface{}) (*LinkService, error) {
	resp := &LinkService{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *LinkServiceClient) List(opts *ListOpts) (*LinkServiceCollection, error) {
	resp := &LinkServiceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *LinkServiceClient) ById(id string) (*LinkService, error) {
	resp := &LinkService{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *LinkServiceClient) Delete(container *LinkService) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
