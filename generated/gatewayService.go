package client

const (
	CONTAINER_TYPE = "gatewayService"
)

type GatewayService struct {
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

type GatewayServiceCollection struct {
	Collection
	Data []GatewayService `json:"data,omitempty"`
}

type GatewayServiceClient struct {
	rancherClient *RancherClient
}

type GatewayServiceOperations interface {
	List(opts *ListOpts) (*GatewayServiceCollection, error)
	Create(opts *GatewayService) (*GatewayService, error)
	Update(existing *GatewayService, updates interface{}) (*GatewayService, error)
	ById(id string) (*GatewayService, error)
	Delete(container *GatewayService) error
}

func newGatewayServiceClient(rancherClient *RancherClient) *GatewayServiceClient {
	return &GatewayServiceClient{
		rancherClient: rancherClient,
	}
}

func (self *GatewayServiceClient) Create(container *GatewayService) (*GatewayService, error) {
	resp := &GatewayService{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *GatewayServiceClient) Update(existing *GatewayService, updates interface{}) (*GatewayService, error) {
	resp := &GatewayService{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *GatewayServiceClient) List(opts *ListOpts) (*GatewayServiceCollection, error) {
	resp := &GatewayServiceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *GatewayServiceClient) ById(id string) (*GatewayService, error) {
	resp := &GatewayService{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *GatewayServiceClient) Delete(container *GatewayService) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
