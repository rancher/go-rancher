package client

const (
	VMWAREVSPHERE_CONFIG_TYPE = "vmwarevsphereConfig"
)

type VmwarevsphereConfig struct {
	Resource

	Boot2dockerUrl string `json:"boot2dockerUrl,omitempty"`

	ComputeIp string `json:"computeIp,omitempty"`

	CpuCount string `json:"cpuCount,omitempty"`

	Datacenter string `json:"datacenter,omitempty"`

	Datastore string `json:"datastore,omitempty"`

	DiskSize string `json:"diskSize,omitempty"`

	MemorySize string `json:"memorySize,omitempty"`

	Network string `json:"network,omitempty"`

	Password string `json:"password,omitempty"`

	Pool string `json:"pool,omitempty"`

	Username string `json:"username,omitempty"`

	Vcenter string `json:"vcenter,omitempty"`
}

type VmwarevsphereConfigCollection struct {
	Collection
	Data []VmwarevsphereConfig `json:"data,omitempty"`
}

type VmwarevsphereConfigClient struct {
	rancherClient *RancherClient
}

type VmwarevsphereConfigOperations interface {
	List(opts *ListOpts) (*VmwarevsphereConfigCollection, error)
	Create(opts *VmwarevsphereConfig) (*VmwarevsphereConfig, error)
	Update(existing *VmwarevsphereConfig, updates interface{}) (*VmwarevsphereConfig, error)
	ById(id string) (*VmwarevsphereConfig, error)
	Delete(container *VmwarevsphereConfig) error
}

func newVmwarevsphereConfigClient(rancherClient *RancherClient) *VmwarevsphereConfigClient {
	return &VmwarevsphereConfigClient{
		rancherClient: rancherClient,
	}
}

func (c *VmwarevsphereConfigClient) Create(container *VmwarevsphereConfig) (*VmwarevsphereConfig, error) {
	resp := &VmwarevsphereConfig{}
	err := c.rancherClient.doCreate(VMWAREVSPHERE_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *VmwarevsphereConfigClient) Update(existing *VmwarevsphereConfig, updates interface{}) (*VmwarevsphereConfig, error) {
	resp := &VmwarevsphereConfig{}
	err := c.rancherClient.doUpdate(VMWAREVSPHERE_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VmwarevsphereConfigClient) List(opts *ListOpts) (*VmwarevsphereConfigCollection, error) {
	resp := &VmwarevsphereConfigCollection{}
	err := c.rancherClient.doList(VMWAREVSPHERE_CONFIG_TYPE, opts, resp)
	return resp, err
}

func (c *VmwarevsphereConfigClient) ById(id string) (*VmwarevsphereConfig, error) {
	resp := &VmwarevsphereConfig{}
	err := c.rancherClient.doById(VMWAREVSPHERE_CONFIG_TYPE, id, resp)
	return resp, err
}

func (c *VmwarevsphereConfigClient) Delete(container *VmwarevsphereConfig) error {
	return c.rancherClient.doResourceDelete(VMWAREVSPHERE_CONFIG_TYPE, &container.Resource)
}
