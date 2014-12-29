package client

const (
	CONTAINER_TYPE = "subnetVnetMap"
)

type SubnetVnetMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    SubnetId string `json:"SubnetId,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
    VnetId string `json:"VnetId,omitempty"`
    
}

type SubnetVnetMapCollection struct {
	Collection
	Data []SubnetVnetMap `json:"data,omitempty"`
}

type SubnetVnetMapClient struct {
	rancherClient *RancherClient
}

type SubnetVnetMapOperations interface {
	List(opts *ListOpts) (*SubnetVnetMapCollection, error)
	Create(opts *SubnetVnetMap) (*SubnetVnetMap, error)
	Update(existing *SubnetVnetMap, updates interface{}) (*SubnetVnetMap, error)
	ById(id string) (*SubnetVnetMap, error)
	Delete(container *SubnetVnetMap) error
}

func newSubnetVnetMapClient(rancherClient *RancherClient) *SubnetVnetMapClient {
	return &SubnetVnetMapClient{
		rancherClient: rancherClient,
	}
}

func (self *SubnetVnetMapClient) Create(container *SubnetVnetMap) (*SubnetVnetMap, error) {
	resp := &SubnetVnetMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *SubnetVnetMapClient) Update(existing *SubnetVnetMap, updates interface{}) (*SubnetVnetMap, error) {
	resp := &SubnetVnetMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *SubnetVnetMapClient) List(opts *ListOpts) (*SubnetVnetMapCollection, error) {
	resp := &SubnetVnetMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *SubnetVnetMapClient) ById(id string) (*SubnetVnetMap, error) {
	resp := &SubnetVnetMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *SubnetVnetMapClient) Delete(container *SubnetVnetMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
