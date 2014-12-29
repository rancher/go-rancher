package client

const (
	CONTAINER_TYPE = "hostNatGatewayService"
)

type HostNatGatewayService struct {
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

type HostNatGatewayServiceCollection struct {
	Collection
	Data []HostNatGatewayService `json:"data,omitempty"`
}

type HostNatGatewayServiceClient struct {
	rancherClient *RancherClient
}

type HostNatGatewayServiceOperations interface {
	List(opts *ListOpts) (*HostNatGatewayServiceCollection, error)
	Create(opts *HostNatGatewayService) (*HostNatGatewayService, error)
	Update(existing *HostNatGatewayService, updates interface{}) (*HostNatGatewayService, error)
	ById(id string) (*HostNatGatewayService, error)
	Delete(container *HostNatGatewayService) error
}

func newHostNatGatewayServiceClient(rancherClient *RancherClient) *HostNatGatewayServiceClient {
	return &HostNatGatewayServiceClient{
		rancherClient: rancherClient,
	}
}

func (self *HostNatGatewayServiceClient) Create(container *HostNatGatewayService) (*HostNatGatewayService, error) {
	resp := &HostNatGatewayService{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *HostNatGatewayServiceClient) Update(existing *HostNatGatewayService, updates interface{}) (*HostNatGatewayService, error) {
	resp := &HostNatGatewayService{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *HostNatGatewayServiceClient) List(opts *ListOpts) (*HostNatGatewayServiceCollection, error) {
	resp := &HostNatGatewayServiceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *HostNatGatewayServiceClient) ById(id string) (*HostNatGatewayService, error) {
	resp := &HostNatGatewayService{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *HostNatGatewayServiceClient) Delete(container *HostNatGatewayService) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
