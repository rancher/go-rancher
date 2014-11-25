package client

const (
	CONTAINER_TYPE = "ipAddress"
)

type IpAddress struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Address string `json:"Address,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Hostname string `json:"Hostname,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    IpPoolId string `json:"IpPoolId,omitempty"`
    
    IsPublic boolean `json:"IsPublic,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkId string `json:"NetworkId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    Role string `json:"Role,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    SubnetId string `json:"SubnetId,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type IpAddressCollection struct {
	Collection
	Data []IpAddress `json:"data,omitempty"`
}

type IpAddressClient struct {
	rancherClient *RancherClient
}

type IpAddressOperations interface {
	List(opts *ListOpts) (*IpAddressCollection, error)
	Create(opts *IpAddress) (*IpAddress, error)
	Update(existing *IpAddress, updates interface{}) (*IpAddress, error)
	ById(id string) (*IpAddress, error)
	Delete(container *IpAddress) error
}

func newIpAddressClient(rancherClient *RancherClient) *IpAddressClient {
	return &IpAddressClient{
		rancherClient: rancherClient,
	}
}

func (self *IpAddressClient) Create(container *IpAddress) (*IpAddress, error) {
	resp := &IpAddress{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *IpAddressClient) Update(existing *IpAddress, updates interface{}) (*IpAddress, error) {
	resp := &IpAddress{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *IpAddressClient) List(opts *ListOpts) (*IpAddressCollection, error) {
	resp := &IpAddressCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *IpAddressClient) ById(id string) (*IpAddress, error) {
	resp := &IpAddress{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *IpAddressClient) Delete(container *IpAddress) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
