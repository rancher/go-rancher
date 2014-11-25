package client

const (
	CONTAINER_TYPE = "externalHandlerProcess"
)

type ExternalHandlerProcess struct {
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
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type ExternalHandlerProcessCollection struct {
	Collection
	Data []ExternalHandlerProcess `json:"data,omitempty"`
}

type ExternalHandlerProcessClient struct {
	rancherClient *RancherClient
}

type ExternalHandlerProcessOperations interface {
	List(opts *ListOpts) (*ExternalHandlerProcessCollection, error)
	Create(opts *ExternalHandlerProcess) (*ExternalHandlerProcess, error)
	Update(existing *ExternalHandlerProcess, updates interface{}) (*ExternalHandlerProcess, error)
	ById(id string) (*ExternalHandlerProcess, error)
	Delete(container *ExternalHandlerProcess) error
}

func newExternalHandlerProcessClient(rancherClient *RancherClient) *ExternalHandlerProcessClient {
	return &ExternalHandlerProcessClient{
		rancherClient: rancherClient,
	}
}

func (self *ExternalHandlerProcessClient) Create(container *ExternalHandlerProcess) (*ExternalHandlerProcess, error) {
	resp := &ExternalHandlerProcess{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ExternalHandlerProcessClient) Update(existing *ExternalHandlerProcess, updates interface{}) (*ExternalHandlerProcess, error) {
	resp := &ExternalHandlerProcess{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ExternalHandlerProcessClient) List(opts *ListOpts) (*ExternalHandlerProcessCollection, error) {
	resp := &ExternalHandlerProcessCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ExternalHandlerProcessClient) ById(id string) (*ExternalHandlerProcess, error) {
	resp := &ExternalHandlerProcess{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ExternalHandlerProcessClient) Delete(container *ExternalHandlerProcess) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
