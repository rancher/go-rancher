package client

const (
	GENERIC_CONFIG_TYPE = "genericConfig"
)

type GenericConfig struct {
	Resource

	IpAddress string `json:"ipAddress,omitempty" yaml:"ip_address,omitempty"`

	SshKey string `json:"sshKey,omitempty" yaml:"ssh_key,omitempty"`

	SshPort string `json:"sshPort,omitempty" yaml:"ssh_port,omitempty"`

	SshUser string `json:"sshUser,omitempty" yaml:"ssh_user,omitempty"`
}

type GenericConfigCollection struct {
	Collection
	Data []GenericConfig `json:"data,omitempty"`
}

type GenericConfigClient struct {
	rancherClient *RancherClient
}

type GenericConfigOperations interface {
	List(opts *ListOpts) (*GenericConfigCollection, error)
	Create(opts *GenericConfig) (*GenericConfig, error)
	Update(existing *GenericConfig, updates interface{}) (*GenericConfig, error)
	ById(id string) (*GenericConfig, error)
	Delete(container *GenericConfig) error
}

func newGenericConfigClient(rancherClient *RancherClient) *GenericConfigClient {
	return &GenericConfigClient{
		rancherClient: rancherClient,
	}
}

func (c *GenericConfigClient) Create(container *GenericConfig) (*GenericConfig, error) {
	resp := &GenericConfig{}
	err := c.rancherClient.doCreate(GENERIC_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *GenericConfigClient) Update(existing *GenericConfig, updates interface{}) (*GenericConfig, error) {
	resp := &GenericConfig{}
	err := c.rancherClient.doUpdate(GENERIC_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GenericConfigClient) List(opts *ListOpts) (*GenericConfigCollection, error) {
	resp := &GenericConfigCollection{}
	err := c.rancherClient.doList(GENERIC_CONFIG_TYPE, opts, resp)
	return resp, err
}

func (c *GenericConfigClient) ById(id string) (*GenericConfig, error) {
	resp := &GenericConfig{}
	err := c.rancherClient.doById(GENERIC_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *GenericConfigClient) Delete(container *GenericConfig) error {
	return c.rancherClient.doResourceDelete(GENERIC_CONFIG_TYPE, &container.Resource)
}
