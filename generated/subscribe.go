package client

const (
	CONTAINER_TYPE = "subscribe"
)

type Subscribe struct {
	Resource
    
    AgentId string `json:"AgentId,omitempty"`
    
    EventNames array[string] `json:"EventNames,omitempty"`
    
}

type SubscribeCollection struct {
	Collection
	Data []Subscribe `json:"data,omitempty"`
}

type SubscribeClient struct {
	rancherClient *RancherClient
}

type SubscribeOperations interface {
	List(opts *ListOpts) (*SubscribeCollection, error)
	Create(opts *Subscribe) (*Subscribe, error)
	Update(existing *Subscribe, updates interface{}) (*Subscribe, error)
	ById(id string) (*Subscribe, error)
	Delete(container *Subscribe) error
}

func newSubscribeClient(rancherClient *RancherClient) *SubscribeClient {
	return &SubscribeClient{
		rancherClient: rancherClient,
	}
}

func (self *SubscribeClient) Create(container *Subscribe) (*Subscribe, error) {
	resp := &Subscribe{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *SubscribeClient) Update(existing *Subscribe, updates interface{}) (*Subscribe, error) {
	resp := &Subscribe{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *SubscribeClient) List(opts *ListOpts) (*SubscribeCollection, error) {
	resp := &SubscribeCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *SubscribeClient) ById(id string) (*Subscribe, error) {
	resp := &Subscribe{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *SubscribeClient) Delete(container *Subscribe) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
