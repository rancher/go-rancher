package client

const (
	OPENSTACK_CONFIG_TYPE = "openstackConfig"
)

type OpenstackConfig struct {
	Resource
    
    AuthUrl string `json:"authUrl,omitempty"`
    
    DomainId string `json:"domainId,omitempty"`
    
    DomainName string `json:"domainName,omitempty"`
    
    EndpointType string `json:"endpointType,omitempty"`
    
    FlavorId string `json:"flavorId,omitempty"`
    
    FlavorName string `json:"flavorName,omitempty"`
    
    FloatingipPool string `json:"floatingipPool,omitempty"`
    
    ImageId string `json:"imageId,omitempty"`
    
    ImageName string `json:"imageName,omitempty"`
    
    Insecure bool `json:"insecure,omitempty"`
    
    NetId string `json:"netId,omitempty"`
    
    NetName string `json:"netName,omitempty"`
    
    Password string `json:"password,omitempty"`
    
    Region string `json:"region,omitempty"`
    
    SecGroups string `json:"secGroups,omitempty"`
    
    SshPort string `json:"sshPort,omitempty"`
    
    SshUser string `json:"sshUser,omitempty"`
    
    TenantId string `json:"tenantId,omitempty"`
    
    TenantName string `json:"tenantName,omitempty"`
    
    Username string `json:"username,omitempty"`
    
}

type OpenstackConfigCollection struct {
	Collection
	Data []OpenstackConfig `json:"data,omitempty"`
}

type OpenstackConfigClient struct {
	rancherClient *RancherClient
}

type OpenstackConfigOperations interface {
	List(opts *ListOpts) (*OpenstackConfigCollection, error)
	Create(opts *OpenstackConfig) (*OpenstackConfig, error)
	Update(existing *OpenstackConfig, updates interface{}) (*OpenstackConfig, error)
	ById(id string) (*OpenstackConfig, error)
	Delete(container *OpenstackConfig) error
}

func newOpenstackConfigClient(rancherClient *RancherClient) *OpenstackConfigClient {
	return &OpenstackConfigClient{
		rancherClient: rancherClient,
	}
}

func (c *OpenstackConfigClient) Create(container *OpenstackConfig) (*OpenstackConfig, error) {
	resp := &OpenstackConfig{}
	err := c.rancherClient.doCreate(OPENSTACK_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *OpenstackConfigClient) Update(existing *OpenstackConfig, updates interface{}) (*OpenstackConfig, error) {
	resp := &OpenstackConfig{}
	err := c.rancherClient.doUpdate(OPENSTACK_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *OpenstackConfigClient) List(opts *ListOpts) (*OpenstackConfigCollection, error) {
	resp := &OpenstackConfigCollection{}
	err := c.rancherClient.doList(OPENSTACK_CONFIG_TYPE, opts, resp)
	return resp, err
}

func (c *OpenstackConfigClient) ById(id string) (*OpenstackConfig, error) {
	resp := &OpenstackConfig{}
	err := c.rancherClient.doById(OPENSTACK_CONFIG_TYPE, id, resp)
	return resp, err
}

func (c *OpenstackConfigClient) Delete(container *OpenstackConfig) error {
	return c.rancherClient.doResourceDelete(OPENSTACK_CONFIG_TYPE, &container.Resource)
}
