package client

const (
	SOFTLAYER_CONFIG_TYPE = "softlayerConfig"
)

type SoftlayerConfig struct {
	Resource

	ApiEndpoint string `json:"apiEndpoint,omitempty"`

	ApiKey string `json:"apiKey,omitempty"`

	Cpu string `json:"cpu,omitempty"`

	DiskSize string `json:"diskSize,omitempty"`

	Domain string `json:"domain,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	HourlyBilling bool `json:"hourlyBilling,omitempty"`

	Image string `json:"image,omitempty"`

	LocalDisk bool `json:"localDisk,omitempty"`

	Memory string `json:"memory,omitempty"`

	PrivateNetOnly bool `json:"privateNetOnly,omitempty"`

	PrivateVlanId string `json:"privateVlanId,omitempty"`

	PublicVlanId string `json:"publicVlanId,omitempty"`

	Region string `json:"region,omitempty"`

	User string `json:"user,omitempty"`
}

type SoftlayerConfigCollection struct {
	Collection
	Data []SoftlayerConfig `json:"data,omitempty"`
}

type SoftlayerConfigClient struct {
	rancherClient *RancherClient
}

type SoftlayerConfigOperations interface {
	List(opts *ListOpts) (*SoftlayerConfigCollection, error)
	Create(opts *SoftlayerConfig) (*SoftlayerConfig, error)
	Update(existing *SoftlayerConfig, updates interface{}) (*SoftlayerConfig, error)
	ById(id string) (*SoftlayerConfig, error)
	Delete(container *SoftlayerConfig) error
}

func newSoftlayerConfigClient(rancherClient *RancherClient) *SoftlayerConfigClient {
	return &SoftlayerConfigClient{
		rancherClient: rancherClient,
	}
}

func (c *SoftlayerConfigClient) Create(container *SoftlayerConfig) (*SoftlayerConfig, error) {
	resp := &SoftlayerConfig{}
	err := c.rancherClient.doCreate(SOFTLAYER_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *SoftlayerConfigClient) Update(existing *SoftlayerConfig, updates interface{}) (*SoftlayerConfig, error) {
	resp := &SoftlayerConfig{}
	err := c.rancherClient.doUpdate(SOFTLAYER_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SoftlayerConfigClient) List(opts *ListOpts) (*SoftlayerConfigCollection, error) {
	resp := &SoftlayerConfigCollection{}
	err := c.rancherClient.doList(SOFTLAYER_CONFIG_TYPE, opts, resp)
	return resp, err
}

func (c *SoftlayerConfigClient) ById(id string) (*SoftlayerConfig, error) {
	resp := &SoftlayerConfig{}
	err := c.rancherClient.doById(SOFTLAYER_CONFIG_TYPE, id, resp)
	return resp, err
}

func (c *SoftlayerConfigClient) Delete(container *SoftlayerConfig) error {
	return c.rancherClient.doResourceDelete(SOFTLAYER_CONFIG_TYPE, &container.Resource)
}
