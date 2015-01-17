package client

const (
	AGENT_TYPE = "agent"
)

type Agent struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    AgentGroupId string `json:"agentGroupId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    ManagedConfig bool `json:"managedConfig,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uri string `json:"uri,omitempty"`
    
    User string `json:"user,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    ZoneId string `json:"zoneId,omitempty"`
    
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

func (c *AgentClient) Create(container *Agent) (*Agent, error) {
	resp := &Agent{}
	err := c.rancherClient.doCreate(AGENT_TYPE, container, resp)
	return resp, err
}

func (c *AgentClient) Update(existing *Agent, updates interface{}) (*Agent, error) {
	resp := &Agent{}
	err := c.rancherClient.doUpdate(AGENT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AgentClient) List(opts *ListOpts) (*AgentCollection, error) {
	resp := &AgentCollection{}
	err := c.rancherClient.doList(AGENT_TYPE, opts, resp)
	return resp, err
}

func (c *AgentClient) ById(id string) (*Agent, error) {
	resp := &Agent{}
	err := c.rancherClient.doById(AGENT_TYPE, id, resp)
	return resp, err
}

func (c *AgentClient) Delete(container *Agent) error {
	return c.rancherClient.doResourceDelete(AGENT_TYPE, &container.Resource)
}
