package client

const (
	AGENT_GROUP_TYPE = "agentGroup"
)

type AgentGroup struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
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

func (c *AgentGroupClient) Create(container *AgentGroup) (*AgentGroup, error) {
	resp := &AgentGroup{}
	err := c.rancherClient.doCreate(AGENT_GROUP_TYPE, container, resp)
	return resp, err
}

func (c *AgentGroupClient) Update(existing *AgentGroup, updates interface{}) (*AgentGroup, error) {
	resp := &AgentGroup{}
	err := c.rancherClient.doUpdate(AGENT_GROUP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AgentGroupClient) List(opts *ListOpts) (*AgentGroupCollection, error) {
	resp := &AgentGroupCollection{}
	err := c.rancherClient.doList(AGENT_GROUP_TYPE, opts, resp)
	return resp, err
}

func (c *AgentGroupClient) ById(id string) (*AgentGroup, error) {
	resp := &AgentGroup{}
	err := c.rancherClient.doById(AGENT_GROUP_TYPE, id, resp)
	return resp, err
}

func (c *AgentGroupClient) Delete(container *AgentGroup) error {
	return c.rancherClient.doResourceDelete(AGENT_GROUP_TYPE, &container.Resource)
}
