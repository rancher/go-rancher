package client

const (
	CONTAINER_TYPE = "instanceHostMap"
)

type InstanceHostMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    HostId string `json:"HostId,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    InstanceId string `json:"InstanceId,omitempty"`
    
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

type InstanceHostMapCollection struct {
	Collection
	Data []InstanceHostMap `json:"data,omitempty"`
}

type InstanceHostMapClient struct {
	rancherClient *RancherClient
}

type InstanceHostMapOperations interface {
	List(opts *ListOpts) (*InstanceHostMapCollection, error)
	Create(opts *InstanceHostMap) (*InstanceHostMap, error)
	Update(existing *InstanceHostMap, updates interface{}) (*InstanceHostMap, error)
	ById(id string) (*InstanceHostMap, error)
	Delete(container *InstanceHostMap) error
}

func newInstanceHostMapClient(rancherClient *RancherClient) *InstanceHostMapClient {
	return &InstanceHostMapClient{
		rancherClient: rancherClient,
	}
}

func (self *InstanceHostMapClient) Create(container *InstanceHostMap) (*InstanceHostMap, error) {
	resp := &InstanceHostMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *InstanceHostMapClient) Update(existing *InstanceHostMap, updates interface{}) (*InstanceHostMap, error) {
	resp := &InstanceHostMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *InstanceHostMapClient) List(opts *ListOpts) (*InstanceHostMapCollection, error) {
	resp := &InstanceHostMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *InstanceHostMapClient) ById(id string) (*InstanceHostMap, error) {
	resp := &InstanceHostMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *InstanceHostMapClient) Delete(container *InstanceHostMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
