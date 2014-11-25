package client

const (
	CONTAINER_TYPE = "agent"
)

type Agent struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    AgentGroupId string `json:"AgentGroupId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    ManagedConfig boolean `json:"ManagedConfig,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uri string `json:"Uri,omitempty"`
    
    User string `json:"User,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
    ZoneId string `json:"ZoneId,omitempty"`
    
}

type AgentCollection struct {
	Collection
	Data []Agent `json:"data,omitempty"`
}

type AgentClient struct {
	rancherClient *RancherClient
}

type AgentOperations interface {
	List(opts *ListOpts) (*AgentCollection, error)
	Create(opts *Agent) (*Agent, error)
	Update(existing *Agent, updates interface{}) (*Agent, error)
	ById(id string) (*Agent, error)
	Delete(container *Agent) error
}

func newAgentClient(rancherClient *RancherClient) *AgentClient {
	return &AgentClient{
		rancherClient: rancherClient,
	}
}

func (self *AgentClient) Create(container *Agent) (*Agent, error) {
	resp := &Agent{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *AgentClient) Update(existing *Agent, updates interface{}) (*Agent, error) {
	resp := &Agent{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *AgentClient) List(opts *ListOpts) (*AgentCollection, error) {
	resp := &AgentCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *AgentClient) ById(id string) (*Agent, error) {
	resp := &Agent{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *AgentClient) Delete(container *Agent) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
