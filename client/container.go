package client

const (
	container_TYPE = "container"
)

type Container struct {
	Resource

	AccountId string `json:"accountId,omitempty"`

	AgentId string `json:"agentId,omitempty"`

	AllocationState string `json:"allocationState,omitempty"`

	Command string `json:"command,omitempty"`

	Compute int `json:"compute,omitempty"`

	Count int `json:"count,omitempty"`

	Created string `json:"created,omitempty"`

	CredentialIds []string `json:"credentialIds,omitempty"`

	Data map[string]interface{} `json:"data,omitempty"`

	Description string `json:"description,omitempty"`

	Directory string `json:"directory,omitempty"`

	Domain string `json:"domain,omitempty"`

	Environment map[string]interface{} `json:"environment,omitempty"`

	FirstRunning string `json:"firstRunning,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	Id string `json:"id,omitempty"`

	ImageId string `json:"imageId,omitempty"`

	ImageUuid string `json:"imageUuid,omitempty"`

	InstanceLinks map[string]interface{} `json:"instanceLinks,omitempty"`

	InstanceTriggeredStop string `json:"instanceTriggeredStop,omitempty"`

	Kind string `json:"kind,omitempty"`

	MemoryMb int `json:"memoryMb,omitempty"`

	Name string `json:"name,omitempty"`

	NetworkIds []string `json:"networkIds,omitempty"`

	PrimaryAssociatedIpAddress string `json:"primaryAssociatedIpAddress,omitempty"`

	PrimaryIpAddress string `json:"primaryIpAddress,omitempty"`

	Privileged bool `json:"privileged,omitempty"`

	PublishAllPorts bool `json:"publishAllPorts,omitempty"`

	RemoveTime string `json:"removeTime,omitempty"`

	Removed string `json:"removed,omitempty"`

	RequestedHostId string `json:"requestedHostId,omitempty"`

	StartOnCreate bool `json:"startOnCreate,omitempty"`

	State string `json:"state,omitempty"`

	SubnetIds []string `json:"subnetIds,omitempty"`

	Token string `json:"token,omitempty"`

	Transitioning string `json:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty"`

	TransitioningProgress int `json:"transitioningProgress,omitempty"`

	User string `json:"user,omitempty"`

	Userdata string `json:"userdata,omitempty"`

	Uuid string `json:"uuid,omitempty"`

	ValidHostIds []string `json:"validHostIds,omitempty"`

	VnetIds []string `json:"vnetIds,omitempty"`

	ZoneId string `json:"zoneId,omitempty"`
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
}

func newContainerClient(rancherClient *RancherClient) *ContainerClient {
	return &ContainerClient{
		rancherClient: rancherClient,
	}
}

func (self *ContainerClient) Create(container *Container) (*Container, error) {
	resp := &Container{}
	err := self.rancherClient.doCreate(container_TYPE, container, resp)
	return resp, err
}

func (self *ContainerClient) Update(existing *Container, updates interface{}) (*Container, error) {
	resp := &Container{}
	err := self.rancherClient.doUpdate(container_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ContainerClient) List(opts *ListOpts) (*ContainerCollection, error) {
	resp := &ContainerCollection{}
	err := self.rancherClient.doList(container_TYPE, opts, resp)
	return resp, err
}

func (self *ContainerClient) ById(id string) (*Container, error) {
	resp := &Container{}
	err := self.rancherClient.doById(container_TYPE, id, resp)
	return resp, err
}

func (self *ContainerClient) Delete(container *Container) error {
	return self.rancherClient.doResourceDelete(container_TYPE, &container.Resource)
}

func (c *Container) Allocate() *Container { return nil }

func (c *Container) Console(input *InstanceConsoleInput) *Container { return nil }

func (c *Container) Create() *Container { return nil }

func (c *Container) Deallocate() *Container { return nil }

func (c *Container) Migrate() *Container { return nil }

func (c *Container) Purge() *Container { return nil }

func (c *Container) Remove() *Container { return nil }

func (c *Container) Restart() *Container { return nil }

func (c *Container) Restore() *Container { return nil }

func (c *Container) Start() *Container { return nil }

func (c *Container) Stop(input *InstanceStop) *Container { return nil }

func (c *Container) Update() *Container { return nil }
