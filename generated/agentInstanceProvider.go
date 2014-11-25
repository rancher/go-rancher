package client

const (
	CONTAINER_TYPE = "agentInstanceProvider"
)

type AgentInstanceProvider struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    AgentInstanceImageUuid string `json:"AgentInstanceImageUuid,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkId string `json:"NetworkId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type AgentInstanceProviderCollection struct {
	Collection
	Data []AgentInstanceProvider `json:"data,omitempty"`
}

type AgentInstanceProviderClient struct {
	rancherClient *RancherClient
}

type AgentInstanceProviderOperations interface {
	List(opts *ListOpts) (*AgentInstanceProviderCollection, error)
	Create(opts *AgentInstanceProvider) (*AgentInstanceProvider, error)
	Update(existing *AgentInstanceProvider, updates interface{}) (*AgentInstanceProvider, error)
	ById(id string) (*AgentInstanceProvider, error)
	Delete(container *AgentInstanceProvider) error
}

func newAgentInstanceProviderClient(rancherClient *RancherClient) *AgentInstanceProviderClient {
	return &AgentInstanceProviderClient{
		rancherClient: rancherClient,
	}
}

func (self *AgentInstanceProviderClient) Create(container *AgentInstanceProvider) (*AgentInstanceProvider, error) {
	resp := &AgentInstanceProvider{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *AgentInstanceProviderClient) Update(existing *AgentInstanceProvider, updates interface{}) (*AgentInstanceProvider, error) {
	resp := &AgentInstanceProvider{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *AgentInstanceProviderClient) List(opts *ListOpts) (*AgentInstanceProviderCollection, error) {
	resp := &AgentInstanceProviderCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *AgentInstanceProviderClient) ById(id string) (*AgentInstanceProvider, error) {
	resp := &AgentInstanceProvider{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *AgentInstanceProviderClient) Delete(container *AgentInstanceProvider) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
