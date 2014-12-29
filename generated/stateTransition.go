package client

const (
	CONTAINER_TYPE = "stateTransition"
)

type StateTransition struct {
	Resource
    
    Field string `json:"Field,omitempty"`
    
    FromState string `json:"FromState,omitempty"`
    
    ToState string `json:"ToState,omitempty"`
    
    Type enum `json:"Type,omitempty"`
    
}

type StateTransitionCollection struct {
	Collection
	Data []StateTransition `json:"data,omitempty"`
}

type StateTransitionClient struct {
	rancherClient *RancherClient
}

type StateTransitionOperations interface {
	List(opts *ListOpts) (*StateTransitionCollection, error)
	Create(opts *StateTransition) (*StateTransition, error)
	Update(existing *StateTransition, updates interface{}) (*StateTransition, error)
	ById(id string) (*StateTransition, error)
	Delete(container *StateTransition) error
}

func newStateTransitionClient(rancherClient *RancherClient) *StateTransitionClient {
	return &StateTransitionClient{
		rancherClient: rancherClient,
	}
}

func (self *StateTransitionClient) Create(container *StateTransition) (*StateTransition, error) {
	resp := &StateTransition{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *StateTransitionClient) Update(existing *StateTransition, updates interface{}) (*StateTransition, error) {
	resp := &StateTransition{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *StateTransitionClient) List(opts *ListOpts) (*StateTransitionCollection, error) {
	resp := &StateTransitionCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *StateTransitionClient) ById(id string) (*StateTransition, error) {
	resp := &StateTransition{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *StateTransitionClient) Delete(container *StateTransition) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
