package client

const (
	CONTAINER_TYPE = "ipAddressNicMap"
)

type IpAddressNicMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    IpAddressId string `json:"IpAddressId,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NicId string `json:"NicId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type IpAddressNicMapCollection struct {
	Collection
	Data []IpAddressNicMap `json:"data,omitempty"`
}

type IpAddressNicMapClient struct {
	rancherClient *RancherClient
}

type IpAddressNicMapOperations interface {
	List(opts *ListOpts) (*IpAddressNicMapCollection, error)
	Create(opts *IpAddressNicMap) (*IpAddressNicMap, error)
	Update(existing *IpAddressNicMap, updates interface{}) (*IpAddressNicMap, error)
	ById(id string) (*IpAddressNicMap, error)
	Delete(container *IpAddressNicMap) error
}

func newIpAddressNicMapClient(rancherClient *RancherClient) *IpAddressNicMapClient {
	return &IpAddressNicMapClient{
		rancherClient: rancherClient,
	}
}

func (self *IpAddressNicMapClient) Create(container *IpAddressNicMap) (*IpAddressNicMap, error) {
	resp := &IpAddressNicMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *IpAddressNicMapClient) Update(existing *IpAddressNicMap, updates interface{}) (*IpAddressNicMap, error) {
	resp := &IpAddressNicMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *IpAddressNicMapClient) List(opts *ListOpts) (*IpAddressNicMapCollection, error) {
	resp := &IpAddressNicMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *IpAddressNicMapClient) ById(id string) (*IpAddressNicMap, error) {
	resp := &IpAddressNicMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *IpAddressNicMapClient) Delete(container *IpAddressNicMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
