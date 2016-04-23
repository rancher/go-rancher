package client

const (
	VMWAREFUSION_CONFIG_TYPE = "vmwarefusionConfig"
)

type VmwarefusionConfig struct {
	Resource
}

type VmwarefusionConfigCollection struct {
	Collection
	Data []VmwarefusionConfig `json:"data,omitempty"`
}

type VmwarefusionConfigClient struct {
	rancherClient *RancherClient
}

type VmwarefusionConfigOperations interface {
	List(opts *ListOpts) (*VmwarefusionConfigCollection, error)
	Create(opts *VmwarefusionConfig) (*VmwarefusionConfig, error)
	Update(existing *VmwarefusionConfig, updates interface{}) (*VmwarefusionConfig, error)
	ById(id string) (*VmwarefusionConfig, error)
	Delete(container *VmwarefusionConfig) error
}

func newVmwarefusionConfigClient(rancherClient *RancherClient) *VmwarefusionConfigClient {
	return &VmwarefusionConfigClient{
		rancherClient: rancherClient,
	}
}

func (c *VmwarefusionConfigClient) Create(container *VmwarefusionConfig) (*VmwarefusionConfig, error) {
	resp := &VmwarefusionConfig{}
	err := c.rancherClient.doCreate(VMWAREFUSION_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *VmwarefusionConfigClient) Update(existing *VmwarefusionConfig, updates interface{}) (*VmwarefusionConfig, error) {
	resp := &VmwarefusionConfig{}
	err := c.rancherClient.doUpdate(VMWAREFUSION_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VmwarefusionConfigClient) List(opts *ListOpts) (*VmwarefusionConfigCollection, error) {
	resp := &VmwarefusionConfigCollection{}
	err := c.rancherClient.doList(VMWAREFUSION_CONFIG_TYPE, opts, resp)
	return resp, err
}

func (c *VmwarefusionConfigClient) ById(id string) (*VmwarefusionConfig, error) {
	resp := &VmwarefusionConfig{}
	err := c.rancherClient.doById(VMWAREFUSION_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *VmwarefusionConfigClient) Delete(container *VmwarefusionConfig) error {
	return c.rancherClient.doResourceDelete(VMWAREFUSION_CONFIG_TYPE, &container.Resource)
}
