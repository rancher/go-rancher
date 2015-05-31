package client

const (
	CONTAINER_TYPE = "container"
)

type Container struct {
	Resource

	AccountId string `json:"accountId,omitempty"`

	AgentId string `json:"agentId,omitempty"`

	AllocationState string `json:"allocationState,omitempty"`

	CapAdd []string `json:"capAdd,omitempty"`

	CapDrop []string `json:"capDrop,omitempty"`

	Command []string `json:"command,omitempty"`

	Count int `json:"count,omitempty"`

	CpuSet string `json:"cpuSet,omitempty"`

	CpuShares int `json:"cpuShares,omitempty"`

	Created string `json:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty"`

	DataVolumes []string `json:"dataVolumes,omitempty"`

	DataVolumesFrom []string `json:"dataVolumesFrom,omitempty"`

	Description string `json:"description,omitempty"`

	Devices []string `json:"devices,omitempty"`

	Directory string `json:"directory,omitempty"`

	Dns []string `json:"dns,omitempty"`

	DnsSearch []string `json:"dnsSearch,omitempty"`

	DomainName string `json:"domainName,omitempty"`

	EntryPoint []string `json:"entryPoint,omitempty"`

	Environment map[string]interface{} `json:"environment,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	FirstRunning string `json:"firstRunning,omitempty"`

	HealthCheck *InstanceHealthCheck `json:"healthCheck,omitempty"`

	HealthState string `json:"healthState,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	ImageUuid string `json:"imageUuid,omitempty"`

	InstanceLinks map[string]interface{} `json:"instanceLinks,omitempty"`

	Kind string `json:"kind,omitempty"`

	Labels map[string]interface{} `json:"labels,omitempty"`

	LxcConf map[string]interface{} `json:"lxcConf,omitempty"`

	Memory int `json:"memory,omitempty"`

	MemorySwap int `json:"memorySwap,omitempty"`

	Name string `json:"name,omitempty"`

	NativeContainer bool `json:"nativeContainer,omitempty"`

	NetworkContainerId string `json:"networkContainerId,omitempty"`

	NetworkIds []string `json:"networkIds,omitempty"`

	NetworkMode string `json:"networkMode,omitempty"`

	Ports []string `json:"ports,omitempty"`

	PrimaryIpAddress string `json:"primaryIpAddress,omitempty"`

	Privileged bool `json:"privileged,omitempty"`

	PublishAllPorts bool `json:"publishAllPorts,omitempty"`

	RegistryCredentialId string `json:"registryCredentialId,omitempty"`

	RemoveTime string `json:"removeTime,omitempty"`

	Removed string `json:"removed,omitempty"`

	RequestedHostId string `json:"requestedHostId,omitempty"`

	RestartPolicy *RestartPolicy `json:"restartPolicy,omitempty"`

	StartOnCreate bool `json:"startOnCreate,omitempty"`

	State string `json:"state,omitempty"`

	StdinOpen bool `json:"stdinOpen,omitempty"`

	SystemContainer string `json:"systemContainer,omitempty"`

	Token string `json:"token,omitempty"`

	Transitioning string `json:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty"`

	TransitioningProgress int `json:"transitioningProgress,omitempty"`

	Tty bool `json:"tty,omitempty"`

	User string `json:"user,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}

type ContainerCollection struct {
	Collection
	Data []Container `json:"data,omitempty"`
}

type ContainerClient struct {
	rancherClient *RancherClient
}

type ContainerOperations interface {
	List(opts *ListOpts) (*ContainerCollection, error)
	Create(opts *Container) (*Container, error)
	Update(existing *Container, updates interface{}) (*Container, error)
	ById(id string) (*Container, error)
	Delete(container *Container) error

	ActionSetlabels(*Container, *SetLabelsInput) (*Container, error)
}

func newContainerClient(rancherClient *RancherClient) *ContainerClient {
	return &ContainerClient{
		rancherClient: rancherClient,
	}
}

func (c *ContainerClient) Create(container *Container) (*Container, error) {
	resp := &Container{}
	err := c.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (c *ContainerClient) Update(existing *Container, updates interface{}) (*Container, error) {
	resp := &Container{}
	err := c.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ContainerClient) List(opts *ListOpts) (*ContainerCollection, error) {
	resp := &ContainerCollection{}
	err := c.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (c *ContainerClient) ById(id string) (*Container, error) {
	resp := &Container{}
	err := c.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (c *ContainerClient) Delete(container *Container) error {
	return c.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}

func (c *ContainerClient) ActionSetlabels(resource *Container, input *SetLabelsInput) (*Container, error) {

	resp := &Container{}

	err := c.rancherClient.doAction(CONTAINER_TYPE, "setlabels", &resource.Resource, input, resp)

	return resp, err
}
