package client

const (
	CONTAINER_TYPE = "agentGroup"
)

type AgentGroup struct {
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
    
    State string `json:"State,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type AgentGroupCollection struct {
	Collection
	Data []AgentGroup `json:"data,omitempty"`
}

type AgentGroupClient struct {
	rancherClient *RancherClient
}

type AgentGroupOperations interface {
	List(opts *ListOpts) (*AgentGroupCollection, error)
	Create(opts *AgentGroup) (*AgentGroup, error)
	Update(existing *AgentGroup, updates interface{}) (*AgentGroup, error)
	ById(id string) (*AgentGroup, error)
	Delete(container *AgentGroup) error
}

func newAgentGroupClient(rancherClient *RancherClient) *AgentGroupClient {
	return &AgentGroupClient{
		rancherClient: rancherClient,
	}
}

func (self *AgentGroupClient) Create(container *AgentGroup) (*AgentGroup, error) {
	resp := &AgentGroup{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *AgentGroupClient) Update(existing *AgentGroup, updates interface{}) (*AgentGroup, error) {
	resp := &AgentGroup{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *AgentGroupClient) List(opts *ListOpts) (*AgentGroupCollection, error) {
	resp := &AgentGroupCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *AgentGroupClient) ById(id string) (*AgentGroup, error) {
	resp := &AgentGroup{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *AgentGroupClient) Delete(container *AgentGroup) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
