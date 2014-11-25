package client

const (
	CONTAINER_TYPE = "subnetIpPool"
)

type SubnetIpPool struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    CidrSize int `json:"CidrSize,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    EndAddress string `json:"EndAddress,omitempty"`
    
    Gateway string `json:"Gateway,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    IsPublic boolean `json:"IsPublic,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkAddress string `json:"NetworkAddress,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    StartAddress string `json:"StartAddress,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type SubnetIpPoolCollection struct {
	Collection
	Data []SubnetIpPool `json:"data,omitempty"`
}

type SubnetIpPoolClient struct {
	rancherClient *RancherClient
}

type SubnetIpPoolOperations interface {
	List(opts *ListOpts) (*SubnetIpPoolCollection, error)
	Create(opts *SubnetIpPool) (*SubnetIpPool, error)
	Update(existing *SubnetIpPool, updates interface{}) (*SubnetIpPool, error)
	ById(id string) (*SubnetIpPool, error)
	Delete(container *SubnetIpPool) error
}

func newSubnetIpPoolClient(rancherClient *RancherClient) *SubnetIpPoolClient {
	return &SubnetIpPoolClient{
		rancherClient: rancherClient,
	}
}

func (self *SubnetIpPoolClient) Create(container *SubnetIpPool) (*SubnetIpPool, error) {
	resp := &SubnetIpPool{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *SubnetIpPoolClient) Update(existing *SubnetIpPool, updates interface{}) (*SubnetIpPool, error) {
	resp := &SubnetIpPool{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *SubnetIpPoolClient) List(opts *ListOpts) (*SubnetIpPoolCollection, error) {
	resp := &SubnetIpPoolCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *SubnetIpPoolClient) ById(id string) (*SubnetIpPool, error) {
	resp := &SubnetIpPool{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *SubnetIpPoolClient) Delete(container *SubnetIpPool) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
