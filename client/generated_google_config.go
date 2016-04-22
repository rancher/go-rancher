package client

const (
	GOOGLE_CONFIG_TYPE = "googleConfig"
)

type GoogleConfig struct {
	Resource

	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	DiskSize string `json:"diskSize,omitempty" yaml:"disk_size,omitempty"`

	DiskType string `json:"diskType,omitempty" yaml:"disk_type,omitempty"`

	MachineImage string `json:"machineImage,omitempty" yaml:"machine_image,omitempty"`

	MachineType string `json:"machineType,omitempty" yaml:"machine_type,omitempty"`

	Preemptible bool `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`

	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	Scopes string `json:"scopes,omitempty" yaml:"scopes,omitempty"`

	Tags string `json:"tags,omitempty" yaml:"tags,omitempty"`

	UseExisting bool `json:"useExisting,omitempty" yaml:"use_existing,omitempty"`

	UseInternalIp bool `json:"useInternalIp,omitempty" yaml:"use_internal_ip,omitempty"`

	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

type GoogleConfigCollection struct {
	Collection
	Data []GoogleConfig `json:"data,omitempty"`
}

type GoogleConfigClient struct {
	rancherClient *RancherClient
}

type GoogleConfigOperations interface {
	List(opts *ListOpts) (*GoogleConfigCollection, error)
	Create(opts *GoogleConfig) (*GoogleConfig, error)
	Update(existing *GoogleConfig, updates interface{}) (*GoogleConfig, error)
	ById(id string) (*GoogleConfig, error)
	Delete(container *GoogleConfig) error
}

func newGoogleConfigClient(rancherClient *RancherClient) *GoogleConfigClient {
	return &GoogleConfigClient{
		rancherClient: rancherClient,
	}
}

func (c *GoogleConfigClient) Create(container *GoogleConfig) (*GoogleConfig, error) {
	resp := &GoogleConfig{}
	err := c.rancherClient.doCreate(GOOGLE_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *GoogleConfigClient) Update(existing *GoogleConfig, updates interface{}) (*GoogleConfig, error) {
	resp := &GoogleConfig{}
	err := c.rancherClient.doUpdate(GOOGLE_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GoogleConfigClient) List(opts *ListOpts) (*GoogleConfigCollection, error) {
	resp := &GoogleConfigCollection{}
	err := c.rancherClient.doList(GOOGLE_CONFIG_TYPE, opts, resp)
	return resp, err
}

func (c *GoogleConfigClient) ById(id string) (*GoogleConfig, error) {
	resp := &GoogleConfig{}
	err := c.rancherClient.doById(GOOGLE_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *GoogleConfigClient) Delete(container *GoogleConfig) error {
	return c.rancherClient.doResourceDelete(GOOGLE_CONFIG_TYPE, &container.Resource)
}
