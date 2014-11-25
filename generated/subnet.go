package client

const (
	CONTAINER_TYPE = "subnet"
)

type Subnet struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    CidrSize int `json:"CidrSize,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    EndAddress string `json:"EndAddress,omitempty"`
    
    Gateway string `json:"Gateway,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    IpPoolId string `json:"IpPoolId,omitempty"`
    
    IsPublic boolean `json:"IsPublic,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkAddress string `json:"NetworkAddress,omitempty"`
    
    NetworkId string `json:"NetworkId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    StartAddress string `json:"StartAddress,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type SubnetCollection struct {
	Collection
	Data []Subnet `json:"data,omitempty"`
}

type SubnetClient struct {
	rancherClient *RancherClient
}

type SubnetOperations interface {
	List(opts *ListOpts) (*SubnetCollection, error)
	Create(opts *Subnet) (*Subnet, error)
	Update(existing *Subnet, updates interface{}) (*Subnet, error)
	ById(id string) (*Subnet, error)
	Delete(container *Subnet) error
}

func newSubnetClient(rancherClient *RancherClient) *SubnetClient {
	return &SubnetClient{
		rancherClient: rancherClient,
	}
}

func (self *SubnetClient) Create(container *Subnet) (*Subnet, error) {
	resp := &Subnet{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *SubnetClient) Update(existing *Subnet, updates interface{}) (*Subnet, error) {
	resp := &Subnet{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *SubnetClient) List(opts *ListOpts) (*SubnetCollection, error) {
	resp := &SubnetCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *SubnetClient) ById(id string) (*Subnet, error) {
	resp := &Subnet{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *SubnetClient) Delete(container *Subnet) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
