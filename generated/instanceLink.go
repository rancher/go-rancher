package client

const (
	CONTAINER_TYPE = "instanceLink"
)

type InstanceLink struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    InstanceId string `json:"InstanceId,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    LinkName string `json:"LinkName,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    Ports array[json] `json:"Ports,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    TargetInstanceId string `json:"TargetInstanceId,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type InstanceLinkCollection struct {
	Collection
	Data []InstanceLink `json:"data,omitempty"`
}

type InstanceLinkClient struct {
	rancherClient *RancherClient
}

type InstanceLinkOperations interface {
	List(opts *ListOpts) (*InstanceLinkCollection, error)
	Create(opts *InstanceLink) (*InstanceLink, error)
	Update(existing *InstanceLink, updates interface{}) (*InstanceLink, error)
	ById(id string) (*InstanceLink, error)
	Delete(container *InstanceLink) error
}

func newInstanceLinkClient(rancherClient *RancherClient) *InstanceLinkClient {
	return &InstanceLinkClient{
		rancherClient: rancherClient,
	}
}

func (self *InstanceLinkClient) Create(container *InstanceLink) (*InstanceLink, error) {
	resp := &InstanceLink{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *InstanceLinkClient) Update(existing *InstanceLink, updates interface{}) (*InstanceLink, error) {
	resp := &InstanceLink{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *InstanceLinkClient) List(opts *ListOpts) (*InstanceLinkCollection, error) {
	resp := &InstanceLinkCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *InstanceLinkClient) ById(id string) (*InstanceLink, error) {
	resp := &InstanceLink{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *InstanceLinkClient) Delete(container *InstanceLink) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
