package client

const (
	HYPERV_CONFIG_TYPE = "hypervConfig"
)

type HypervConfig struct {
	Resource

	Boot2dockerUrl string `json:"boot2dockerUrl,omitempty" yaml:"boot2docker_url,omitempty"`

	CpuCount string `json:"cpuCount,omitempty" yaml:"cpu_count,omitempty"`

	DiskSize string `json:"diskSize,omitempty" yaml:"disk_size,omitempty"`

	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`

	VirtualSwitch string `json:"virtualSwitch,omitempty" yaml:"virtual_switch,omitempty"`
}

type HypervConfigCollection struct {
	Collection
	Data []HypervConfig `json:"data,omitempty"`
}

type HypervConfigClient struct {
	rancherClient *RancherClient
}

type HypervConfigOperations interface {
	List(opts *ListOpts) (*HypervConfigCollection, error)
	Create(opts *HypervConfig) (*HypervConfig, error)
	Update(existing *HypervConfig, updates interface{}) (*HypervConfig, error)
	ById(id string) (*HypervConfig, error)
	Delete(container *HypervConfig) error
}

func newHypervConfigClient(rancherClient *RancherClient) *HypervConfigClient {
	return &HypervConfigClient{
		rancherClient: rancherClient,
	}
}

func (c *HypervConfigClient) Create(container *HypervConfig) (*HypervConfig, error) {
	resp := &HypervConfig{}
	err := c.rancherClient.doCreate(HYPERV_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *HypervConfigClient) Update(existing *HypervConfig, updates interface{}) (*HypervConfig, error) {
	resp := &HypervConfig{}
	err := c.rancherClient.doUpdate(HYPERV_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HypervConfigClient) List(opts *ListOpts) (*HypervConfigCollection, error) {
	resp := &HypervConfigCollection{}
	err := c.rancherClient.doList(HYPERV_CONFIG_TYPE, opts, resp)
	return resp, err
}

func (c *HypervConfigClient) ById(id string) (*HypervConfig, error) {
	resp := &HypervConfig{}
	err := c.rancherClient.doById(HYPERV_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *HypervConfigClient) Delete(container *HypervConfig) error {
	return c.rancherClient.doResourceDelete(HYPERV_CONFIG_TYPE, &container.Resource)
}
