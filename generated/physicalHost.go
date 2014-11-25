package client

const (
	CONTAINER_TYPE = "physicalHost"
)

type PhysicalHost struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
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

type PhysicalHostCollection struct {
	Collection
	Data []PhysicalHost `json:"data,omitempty"`
}

type PhysicalHostClient struct {
	rancherClient *RancherClient
}

type PhysicalHostOperations interface {
	List(opts *ListOpts) (*PhysicalHostCollection, error)
	Create(opts *PhysicalHost) (*PhysicalHost, error)
	Update(existing *PhysicalHost, updates interface{}) (*PhysicalHost, error)
	ById(id string) (*PhysicalHost, error)
	Delete(container *PhysicalHost) error
}

func newPhysicalHostClient(rancherClient *RancherClient) *PhysicalHostClient {
	return &PhysicalHostClient{
		rancherClient: rancherClient,
	}
}

func (self *PhysicalHostClient) Create(container *PhysicalHost) (*PhysicalHost, error) {
	resp := &PhysicalHost{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *PhysicalHostClient) Update(existing *PhysicalHost, updates interface{}) (*PhysicalHost, error) {
	resp := &PhysicalHost{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *PhysicalHostClient) List(opts *ListOpts) (*PhysicalHostCollection, error) {
	resp := &PhysicalHostCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *PhysicalHostClient) ById(id string) (*PhysicalHost, error) {
	resp := &PhysicalHost{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *PhysicalHostClient) Delete(container *PhysicalHost) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
