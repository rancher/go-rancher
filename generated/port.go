package client

const (
	CONTAINER_TYPE = "port"
)

type Port struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    InstanceId string `json:"InstanceId,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    PrivateIpAddressId string `json:"PrivateIpAddressId,omitempty"`
    
    PrivatePort int `json:"PrivatePort,omitempty"`
    
    Protocol string `json:"Protocol,omitempty"`
    
    PublicIpAddressId string `json:"PublicIpAddressId,omitempty"`
    
    PublicPort int `json:"PublicPort,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type PortCollection struct {
	Collection
	Data []Port `json:"data,omitempty"`
}

type PortClient struct {
	rancherClient *RancherClient
}

type PortOperations interface {
	List(opts *ListOpts) (*PortCollection, error)
	Create(opts *Port) (*Port, error)
	Update(existing *Port, updates interface{}) (*Port, error)
	ById(id string) (*Port, error)
	Delete(container *Port) error
}

func newPortClient(rancherClient *RancherClient) *PortClient {
	return &PortClient{
		rancherClient: rancherClient,
	}
}

func (self *PortClient) Create(container *Port) (*Port, error) {
	resp := &Port{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *PortClient) Update(existing *Port, updates interface{}) (*Port, error) {
	resp := &Port{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *PortClient) List(opts *ListOpts) (*PortCollection, error) {
	resp := &PortCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *PortClient) ById(id string) (*Port, error) {
	resp := &Port{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *PortClient) Delete(container *Port) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
