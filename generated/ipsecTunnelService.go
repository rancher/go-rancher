package client

const (
	CONTAINER_TYPE = "ipsecTunnelService"
)

type IpsecTunnelService struct {
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

type IpsecTunnelServiceCollection struct {
	Collection
	Data []IpsecTunnelService `json:"data,omitempty"`
}

type IpsecTunnelServiceClient struct {
	rancherClient *RancherClient
}

type IpsecTunnelServiceOperations interface {
	List(opts *ListOpts) (*IpsecTunnelServiceCollection, error)
	Create(opts *IpsecTunnelService) (*IpsecTunnelService, error)
	Update(existing *IpsecTunnelService, updates interface{}) (*IpsecTunnelService, error)
	ById(id string) (*IpsecTunnelService, error)
	Delete(container *IpsecTunnelService) error
}

func newIpsecTunnelServiceClient(rancherClient *RancherClient) *IpsecTunnelServiceClient {
	return &IpsecTunnelServiceClient{
		rancherClient: rancherClient,
	}
}

func (self *IpsecTunnelServiceClient) Create(container *IpsecTunnelService) (*IpsecTunnelService, error) {
	resp := &IpsecTunnelService{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *IpsecTunnelServiceClient) Update(existing *IpsecTunnelService, updates interface{}) (*IpsecTunnelService, error) {
	resp := &IpsecTunnelService{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *IpsecTunnelServiceClient) List(opts *ListOpts) (*IpsecTunnelServiceCollection, error) {
	resp := &IpsecTunnelServiceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *IpsecTunnelServiceClient) ById(id string) (*IpsecTunnelService, error) {
	resp := &IpsecTunnelService{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *IpsecTunnelServiceClient) Delete(container *IpsecTunnelService) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
