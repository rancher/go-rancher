package client

const (
	CONTAINER_TYPE = "dhcpService"
)

type DhcpService struct {
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

type DhcpServiceCollection struct {
	Collection
	Data []DhcpService `json:"data,omitempty"`
}

type DhcpServiceClient struct {
	rancherClient *RancherClient
}

type DhcpServiceOperations interface {
	List(opts *ListOpts) (*DhcpServiceCollection, error)
	Create(opts *DhcpService) (*DhcpService, error)
	Update(existing *DhcpService, updates interface{}) (*DhcpService, error)
	ById(id string) (*DhcpService, error)
	Delete(container *DhcpService) error
}

func newDhcpServiceClient(rancherClient *RancherClient) *DhcpServiceClient {
	return &DhcpServiceClient{
		rancherClient: rancherClient,
	}
}

func (self *DhcpServiceClient) Create(container *DhcpService) (*DhcpService, error) {
	resp := &DhcpService{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *DhcpServiceClient) Update(existing *DhcpService, updates interface{}) (*DhcpService, error) {
	resp := &DhcpService{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *DhcpServiceClient) List(opts *ListOpts) (*DhcpServiceCollection, error) {
	resp := &DhcpServiceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *DhcpServiceClient) ById(id string) (*DhcpService, error) {
	resp := &DhcpService{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *DhcpServiceClient) Delete(container *DhcpService) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
