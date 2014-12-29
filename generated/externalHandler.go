package client

const (
	CONTAINER_TYPE = "externalHandler"
)

type ExternalHandler struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    Priority int `json:"Priority,omitempty"`
    
    PriorityName enum `json:"PriorityName,omitempty"`
    
    ProcessNames array[string] `json:"ProcessNames,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    Retries int `json:"Retries,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    TimeoutMillis int `json:"TimeoutMillis,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type ExternalHandlerCollection struct {
	Collection
	Data []ExternalHandler `json:"data,omitempty"`
}

type ExternalHandlerClient struct {
	rancherClient *RancherClient
}

type ExternalHandlerOperations interface {
	List(opts *ListOpts) (*ExternalHandlerCollection, error)
	Create(opts *ExternalHandler) (*ExternalHandler, error)
	Update(existing *ExternalHandler, updates interface{}) (*ExternalHandler, error)
	ById(id string) (*ExternalHandler, error)
	Delete(container *ExternalHandler) error
}

func newExternalHandlerClient(rancherClient *RancherClient) *ExternalHandlerClient {
	return &ExternalHandlerClient{
		rancherClient: rancherClient,
	}
}

func (self *ExternalHandlerClient) Create(container *ExternalHandler) (*ExternalHandler, error) {
	resp := &ExternalHandler{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ExternalHandlerClient) Update(existing *ExternalHandler, updates interface{}) (*ExternalHandler, error) {
	resp := &ExternalHandler{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ExternalHandlerClient) List(opts *ListOpts) (*ExternalHandlerCollection, error) {
	resp := &ExternalHandlerCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ExternalHandlerClient) ById(id string) (*ExternalHandler, error) {
	resp := &ExternalHandler{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ExternalHandlerClient) Delete(container *ExternalHandler) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
