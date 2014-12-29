package client

const (
	CONTAINER_TYPE = "hostVnetMap"
)

type HostVnetMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    HostId string `json:"HostId,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
    VnetId string `json:"VnetId,omitempty"`
    
}

type HostVnetMapCollection struct {
	Collection
	Data []HostVnetMap `json:"data,omitempty"`
}

type HostVnetMapClient struct {
	rancherClient *RancherClient
}

type HostVnetMapOperations interface {
	List(opts *ListOpts) (*HostVnetMapCollection, error)
	Create(opts *HostVnetMap) (*HostVnetMap, error)
	Update(existing *HostVnetMap, updates interface{}) (*HostVnetMap, error)
	ById(id string) (*HostVnetMap, error)
	Delete(container *HostVnetMap) error
}

func newHostVnetMapClient(rancherClient *RancherClient) *HostVnetMapClient {
	return &HostVnetMapClient{
		rancherClient: rancherClient,
	}
}

func (self *HostVnetMapClient) Create(container *HostVnetMap) (*HostVnetMap, error) {
	resp := &HostVnetMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *HostVnetMapClient) Update(existing *HostVnetMap, updates interface{}) (*HostVnetMap, error) {
	resp := &HostVnetMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *HostVnetMapClient) List(opts *ListOpts) (*HostVnetMapCollection, error) {
	resp := &HostVnetMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *HostVnetMapClient) ById(id string) (*HostVnetMap, error) {
	resp := &HostVnetMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *HostVnetMapClient) Delete(container *HostVnetMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
