package client

const (
	CONTAINER_TYPE = "host"
)

type Host struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    AgentId string `json:"AgentId,omitempty"`
    
    ComputeFree int `json:"ComputeFree,omitempty"`
    
    ComputeTotal int `json:"ComputeTotal,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    IsPublic boolean `json:"IsPublic,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    PhysicalHostId string `json:"PhysicalHostId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uri string `json:"Uri,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
    ZoneId string `json:"ZoneId,omitempty"`
    
}

type HostCollection struct {
	Collection
	Data []Host `json:"data,omitempty"`
}

type HostClient struct {
	rancherClient *RancherClient
}

type HostOperations interface {
	List(opts *ListOpts) (*HostCollection, error)
	Create(opts *Host) (*Host, error)
	Update(existing *Host, updates interface{}) (*Host, error)
	ById(id string) (*Host, error)
	Delete(container *Host) error
}

func newHostClient(rancherClient *RancherClient) *HostClient {
	return &HostClient{
		rancherClient: rancherClient,
	}
}

func (self *HostClient) Create(container *Host) (*Host, error) {
	resp := &Host{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *HostClient) Update(existing *Host, updates interface{}) (*Host, error) {
	resp := &Host{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *HostClient) List(opts *ListOpts) (*HostCollection, error) {
	resp := &HostCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *HostClient) ById(id string) (*Host, error) {
	resp := &Host{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *HostClient) Delete(container *Host) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
