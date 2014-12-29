package client

const (
	CONTAINER_TYPE = "hostIpAddressMap"
)

type HostIpAddressMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    HostId string `json:"HostId,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    IpAddressId string `json:"IpAddressId,omitempty"`
    
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

type HostIpAddressMapCollection struct {
	Collection
	Data []HostIpAddressMap `json:"data,omitempty"`
}

type HostIpAddressMapClient struct {
	rancherClient *RancherClient
}

type HostIpAddressMapOperations interface {
	List(opts *ListOpts) (*HostIpAddressMapCollection, error)
	Create(opts *HostIpAddressMap) (*HostIpAddressMap, error)
	Update(existing *HostIpAddressMap, updates interface{}) (*HostIpAddressMap, error)
	ById(id string) (*HostIpAddressMap, error)
	Delete(container *HostIpAddressMap) error
}

func newHostIpAddressMapClient(rancherClient *RancherClient) *HostIpAddressMapClient {
	return &HostIpAddressMapClient{
		rancherClient: rancherClient,
	}
}

func (self *HostIpAddressMapClient) Create(container *HostIpAddressMap) (*HostIpAddressMap, error) {
	resp := &HostIpAddressMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *HostIpAddressMapClient) Update(existing *HostIpAddressMap, updates interface{}) (*HostIpAddressMap, error) {
	resp := &HostIpAddressMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *HostIpAddressMapClient) List(opts *ListOpts) (*HostIpAddressMapCollection, error) {
	resp := &HostIpAddressMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *HostIpAddressMapClient) ById(id string) (*HostIpAddressMap, error) {
	resp := &HostIpAddressMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *HostIpAddressMapClient) Delete(container *HostIpAddressMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
