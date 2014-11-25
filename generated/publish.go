package client

const (
	CONTAINER_TYPE = "publish"
)

type Publish struct {
	Resource
    
    Data map[json] `json:"Data,omitempty"`
    
    Id string `json:"Id,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    PreviousIds array[string] `json:"PreviousIds,omitempty"`
    
    Publisher string `json:"Publisher,omitempty"`
    
    ResourceId string `json:"ResourceId,omitempty"`
    
    ResourceType string `json:"ResourceType,omitempty"`
    
    Time int `json:"Time,omitempty"`
    
    Transitioning string `json:"Transitioning,omitempty"`
    
    TransitioningInternalMessage string `json:"TransitioningInternalMessage,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
}

type PublishCollection struct {
	Collection
	Data []Publish `json:"data,omitempty"`
}

type PublishClient struct {
	rancherClient *RancherClient
}

type PublishOperations interface {
	List(opts *ListOpts) (*PublishCollection, error)
	Create(opts *Publish) (*Publish, error)
	Update(existing *Publish, updates interface{}) (*Publish, error)
	ById(id string) (*Publish, error)
	Delete(container *Publish) error
}

func newPublishClient(rancherClient *RancherClient) *PublishClient {
	return &PublishClient{
		rancherClient: rancherClient,
	}
}

func (self *PublishClient) Create(container *Publish) (*Publish, error) {
	resp := &Publish{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *PublishClient) Update(existing *Publish, updates interface{}) (*Publish, error) {
	resp := &Publish{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *PublishClient) List(opts *ListOpts) (*PublishCollection, error) {
	resp := &PublishCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *PublishClient) ById(id string) (*Publish, error) {
	resp := &Publish{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *PublishClient) Delete(container *Publish) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
