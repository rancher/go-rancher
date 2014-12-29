package client

const (
	CONTAINER_TYPE = "processDefinition"
)

type ProcessDefinition struct {
	Resource
    
    ExtensionBased boolean `json:"ExtensionBased,omitempty"`
    
    Id string `json:"Id,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    PostProcessListeners extensionPoint `json:"PostProcessListeners,omitempty"`
    
    PreProcessListeners extensionPoint `json:"PreProcessListeners,omitempty"`
    
    ProcessHandlers extensionPoint `json:"ProcessHandlers,omitempty"`
    
    ResourceType string `json:"ResourceType,omitempty"`
    
    StateTransitions array[stateTransition] `json:"StateTransitions,omitempty"`
    
}

type ProcessDefinitionCollection struct {
	Collection
	Data []ProcessDefinition `json:"data,omitempty"`
}

type ProcessDefinitionClient struct {
	rancherClient *RancherClient
}

type ProcessDefinitionOperations interface {
	List(opts *ListOpts) (*ProcessDefinitionCollection, error)
	Create(opts *ProcessDefinition) (*ProcessDefinition, error)
	Update(existing *ProcessDefinition, updates interface{}) (*ProcessDefinition, error)
	ById(id string) (*ProcessDefinition, error)
	Delete(container *ProcessDefinition) error
}

func newProcessDefinitionClient(rancherClient *RancherClient) *ProcessDefinitionClient {
	return &ProcessDefinitionClient{
		rancherClient: rancherClient,
	}
}

func (self *ProcessDefinitionClient) Create(container *ProcessDefinition) (*ProcessDefinition, error) {
	resp := &ProcessDefinition{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ProcessDefinitionClient) Update(existing *ProcessDefinition, updates interface{}) (*ProcessDefinition, error) {
	resp := &ProcessDefinition{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ProcessDefinitionClient) List(opts *ListOpts) (*ProcessDefinitionCollection, error) {
	resp := &ProcessDefinitionCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ProcessDefinitionClient) ById(id string) (*ProcessDefinition, error) {
	resp := &ProcessDefinition{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ProcessDefinitionClient) Delete(container *ProcessDefinition) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
