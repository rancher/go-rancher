package client

const (
	VMWAREVCLOUDAIR_CONFIG_TYPE = "vmwarevcloudairConfig"
)

type VmwarevcloudairConfig struct {
	Resource
    
    Catalog string `json:"catalog,omitempty"`
    
    Catalogitem string `json:"catalogitem,omitempty"`
    
    Computeid string `json:"computeid,omitempty"`
    
    CpuCount string `json:"cpuCount,omitempty"`
    
    DockerPort string `json:"dockerPort,omitempty"`
    
    Edgegateway string `json:"edgegateway,omitempty"`
    
    MemorySize string `json:"memorySize,omitempty"`
    
    Orgvdcnetwork string `json:"orgvdcnetwork,omitempty"`
    
    Password string `json:"password,omitempty"`
    
    Provision bool `json:"provision,omitempty"`
    
    Publicip string `json:"publicip,omitempty"`
    
    SshPort string `json:"sshPort,omitempty"`
    
    Username string `json:"username,omitempty"`
    
    Vdcid string `json:"vdcid,omitempty"`
    
}

type VmwarevcloudairConfigCollection struct {
	Collection
	Data []VmwarevcloudairConfig `json:"data,omitempty"`
}

type VmwarevcloudairConfigClient struct {
	rancherClient *RancherClient
}

type VmwarevcloudairConfigOperations interface {
	List(opts *ListOpts) (*VmwarevcloudairConfigCollection, error)
	Create(opts *VmwarevcloudairConfig) (*VmwarevcloudairConfig, error)
	Update(existing *VmwarevcloudairConfig, updates interface{}) (*VmwarevcloudairConfig, error)
	ById(id string) (*VmwarevcloudairConfig, error)
	Delete(container *VmwarevcloudairConfig) error
}

func newVmwarevcloudairConfigClient(rancherClient *RancherClient) *VmwarevcloudairConfigClient {
	return &VmwarevcloudairConfigClient{
		rancherClient: rancherClient,
	}
}

func (c *VmwarevcloudairConfigClient) Create(container *VmwarevcloudairConfig) (*VmwarevcloudairConfig, error) {
	resp := &VmwarevcloudairConfig{}
	err := c.rancherClient.doCreate(VMWAREVCLOUDAIR_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *VmwarevcloudairConfigClient) Update(existing *VmwarevcloudairConfig, updates interface{}) (*VmwarevcloudairConfig, error) {
	resp := &VmwarevcloudairConfig{}
	err := c.rancherClient.doUpdate(VMWAREVCLOUDAIR_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VmwarevcloudairConfigClient) List(opts *ListOpts) (*VmwarevcloudairConfigCollection, error) {
	resp := &VmwarevcloudairConfigCollection{}
	err := c.rancherClient.doList(VMWAREVCLOUDAIR_CONFIG_TYPE, opts, resp)
	return resp, err
}

func (c *VmwarevcloudairConfigClient) ById(id string) (*VmwarevcloudairConfig, error) {
	resp := &VmwarevcloudairConfig{}
	err := c.rancherClient.doById(VMWAREVCLOUDAIR_CONFIG_TYPE, id, resp)
	return resp, err
}

func (c *VmwarevcloudairConfigClient) Delete(container *VmwarevcloudairConfig) error {
	return c.rancherClient.doResourceDelete(VMWAREVCLOUDAIR_CONFIG_TYPE, &container.Resource)
}
