package client

const (
	CONTAINER_TYPE = "ipPool"
)

type IpPool struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    IsPublic boolean `json:"IsPublic,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type IpPoolCollection struct {
	Collection
	Data []IpPool `json:"data,omitempty"`
}

type IpPoolClient struct {
	rancherClient *RancherClient
}

type IpPoolOperations interface {
	List(opts *ListOpts) (*IpPoolCollection, error)
	Create(opts *IpPool) (*IpPool, error)
	Update(existing *IpPool, updates interface{}) (*IpPool, error)
	ById(id string) (*IpPool, error)
	Delete(container *IpPool) error
}

func newIpPoolClient(rancherClient *RancherClient) *IpPoolClient {
	return &IpPoolClient{
		rancherClient: rancherClient,
	}
}

func (self *IpPoolClient) Create(container *IpPool) (*IpPool, error) {
	resp := &IpPool{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *IpPoolClient) Update(existing *IpPool, updates interface{}) (*IpPool, error) {
	resp := &IpPool{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *IpPoolClient) List(opts *ListOpts) (*IpPoolCollection, error) {
	resp := &IpPoolCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *IpPoolClient) ById(id string) (*IpPool, error) {
	resp := &IpPool{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *IpPoolClient) Delete(container *IpPool) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
