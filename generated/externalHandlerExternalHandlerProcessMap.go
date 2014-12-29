package client

const (
	CONTAINER_TYPE = "externalHandlerExternalHandlerProcessMap"
)

type ExternalHandlerExternalHandlerProcessMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    ExternalHandlerId string `json:"ExternalHandlerId,omitempty"`
    
    ExternalHandlerProcessId string `json:"ExternalHandlerProcessId,omitempty"`
    
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
    
}

type ExternalHandlerExternalHandlerProcessMapCollection struct {
	Collection
	Data []ExternalHandlerExternalHandlerProcessMap `json:"data,omitempty"`
}

type ExternalHandlerExternalHandlerProcessMapClient struct {
	rancherClient *RancherClient
}

type ExternalHandlerExternalHandlerProcessMapOperations interface {
	List(opts *ListOpts) (*ExternalHandlerExternalHandlerProcessMapCollection, error)
	Create(opts *ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error)
	Update(existing *ExternalHandlerExternalHandlerProcessMap, updates interface{}) (*ExternalHandlerExternalHandlerProcessMap, error)
	ById(id string) (*ExternalHandlerExternalHandlerProcessMap, error)
	Delete(container *ExternalHandlerExternalHandlerProcessMap) error
}

func newExternalHandlerExternalHandlerProcessMapClient(rancherClient *RancherClient) *ExternalHandlerExternalHandlerProcessMapClient {
	return &ExternalHandlerExternalHandlerProcessMapClient{
		rancherClient: rancherClient,
	}
}

func (self *ExternalHandlerExternalHandlerProcessMapClient) Create(container *ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error) {
	resp := &ExternalHandlerExternalHandlerProcessMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ExternalHandlerExternalHandlerProcessMapClient) Update(existing *ExternalHandlerExternalHandlerProcessMap, updates interface{}) (*ExternalHandlerExternalHandlerProcessMap, error) {
	resp := &ExternalHandlerExternalHandlerProcessMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ExternalHandlerExternalHandlerProcessMapClient) List(opts *ListOpts) (*ExternalHandlerExternalHandlerProcessMapCollection, error) {
	resp := &ExternalHandlerExternalHandlerProcessMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ExternalHandlerExternalHandlerProcessMapClient) ById(id string) (*ExternalHandlerExternalHandlerProcessMap, error) {
	resp := &ExternalHandlerExternalHandlerProcessMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ExternalHandlerExternalHandlerProcessMapClient) Delete(container *ExternalHandlerExternalHandlerProcessMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
