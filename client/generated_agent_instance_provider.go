package client

const (
	AGENT_INSTANCE_PROVIDER_TYPE = "agentInstanceProvider"
)

type AgentInstanceProvider struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    AgentInstanceImageUuid string `json:"agentInstanceImageUuid,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkId string `json:"networkId,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
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

func (c *AgentInstanceProviderClient) Create(container *AgentInstanceProvider) (*AgentInstanceProvider, error) {
	resp := &AgentInstanceProvider{}
	err := c.rancherClient.doCreate(AGENT_INSTANCE_PROVIDER_TYPE, container, resp)
	return resp, err
}

func (c *AgentInstanceProviderClient) Update(existing *AgentInstanceProvider, updates interface{}) (*AgentInstanceProvider, error) {
	resp := &AgentInstanceProvider{}
	err := c.rancherClient.doUpdate(AGENT_INSTANCE_PROVIDER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AgentInstanceProviderClient) List(opts *ListOpts) (*AgentInstanceProviderCollection, error) {
	resp := &AgentInstanceProviderCollection{}
	err := c.rancherClient.doList(AGENT_INSTANCE_PROVIDER_TYPE, opts, resp)
	return resp, err
}

func (c *AgentInstanceProviderClient) ById(id string) (*AgentInstanceProvider, error) {
	resp := &AgentInstanceProvider{}
	err := c.rancherClient.doById(AGENT_INSTANCE_PROVIDER_TYPE, id, resp)
	return resp, err
}

func (c *AgentInstanceProviderClient) Delete(container *AgentInstanceProvider) error {
	return c.rancherClient.doResourceDelete(AGENT_INSTANCE_PROVIDER_TYPE, &container.Resource)
}
