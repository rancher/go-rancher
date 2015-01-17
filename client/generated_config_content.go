package client

const (
	CONFIG_CONTENT_TYPE = "configContent"
)

type ConfigContent struct {
	Resource
    
    AgentId string `json:"agentId,omitempty"`
    
    Current string `json:"current,omitempty"`
    
    Version string `json:"version,omitempty"`
    
}

type ConfigContentCollection struct {
	Collection
	Data []ConfigContent `json:"data,omitempty"`
}

type ConfigContentClient struct {
	rancherClient *RancherClient
}

type ConfigContentOperations interface {
	List(opts *ListOpts) (*ConfigContentCollection, error)
	Create(opts *ConfigContent) (*ConfigContent, error)
	Update(existing *ConfigContent, updates interface{}) (*ConfigContent, error)
	ById(id string) (*ConfigContent, error)
	Delete(container *ConfigContent) error
}

func newConfigContentClient(rancherClient *RancherClient) *ConfigContentClient {
	return &ConfigContentClient{
		rancherClient: rancherClient,
	}
}

func (c *ConfigContentClient) Create(container *ConfigContent) (*ConfigContent, error) {
	resp := &ConfigContent{}
	err := c.rancherClient.doCreate(CONFIG_CONTENT_TYPE, container, resp)
	return resp, err
}

func (c *ConfigContentClient) Update(existing *ConfigContent, updates interface{}) (*ConfigContent, error) {
	resp := &ConfigContent{}
	err := c.rancherClient.doUpdate(CONFIG_CONTENT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ConfigContentClient) List(opts *ListOpts) (*ConfigContentCollection, error) {
	resp := &ConfigContentCollection{}
	err := c.rancherClient.doList(CONFIG_CONTENT_TYPE, opts, resp)
	return resp, err
}

func (c *ConfigContentClient) ById(id string) (*ConfigContent, error) {
	resp := &ConfigContent{}
	err := c.rancherClient.doById(CONFIG_CONTENT_TYPE, id, resp)
	return resp, err
}

func (c *ConfigContentClient) Delete(container *ConfigContent) error {
	return c.rancherClient.doResourceDelete(CONFIG_CONTENT_TYPE, &container.Resource)
}
