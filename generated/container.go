package client

const (
	CONTAINER_TYPE = "container"
)

type Container struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    AgentId string `json:"AgentId,omitempty"`
    
    AllocationState string `json:"AllocationState,omitempty"`
    
    Command string `json:"Command,omitempty"`
    
    CommandArgs array[string] `json:"CommandArgs,omitempty"`
    
    Compute int `json:"Compute,omitempty"`
    
    Count int `json:"Count,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    CredentialIds array[reference[credential]] `json:"CredentialIds,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Directory string `json:"Directory,omitempty"`
    
    DockerVolumes array[string] `json:"DockerVolumes,omitempty"`
    
    Domain string `json:"Domain,omitempty"`
    
    Environment map[string] `json:"Environment,omitempty"`
    
    FirstRunning date `json:"FirstRunning,omitempty"`
    
    Hostname string `json:"Hostname,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    ImageId string `json:"ImageId,omitempty"`
    
    ImageUuid string `json:"ImageUuid,omitempty"`
    
    InstanceLinks map[reference[instance]] `json:"InstanceLinks,omitempty"`
    
    InstanceTriggeredStop enum `json:"InstanceTriggeredStop,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    MemoryMb int `json:"MemoryMb,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkIds array[reference[network]] `json:"NetworkIds,omitempty"`
    
    Ports array[string] `json:"Ports,omitempty"`
    
    PrimaryAssociatedIpAddress string `json:"PrimaryAssociatedIpAddress,omitempty"`
    
    PrimaryIpAddress string `json:"PrimaryIpAddress,omitempty"`
    
    Privileged boolean `json:"Privileged,omitempty"`
    
    PublishAllPorts boolean `json:"PublishAllPorts,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    RequestedHostId string `json:"RequestedHostId,omitempty"`
    
    StartOnCreate boolean `json:"StartOnCreate,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    SubnetIds array[reference[subnet]] `json:"SubnetIds,omitempty"`
    
    Token string `json:"Token,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    User string `json:"User,omitempty"`
    
    Userdata string `json:"Userdata,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
    ValidHostIds array[reference[host]] `json:"ValidHostIds,omitempty"`
    
    VnetIds array[reference[vnet]] `json:"VnetIds,omitempty"`
    
    ZoneId string `json:"ZoneId,omitempty"`
    
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
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ContainerClient) Update(existing *Container, updates interface{}) (*Container, error) {
	resp := &Container{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ContainerClient) List(opts *ListOpts) (*ContainerCollection, error) {
	resp := &ContainerCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ContainerClient) ById(id string) (*Container, error) {
	resp := &Container{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ContainerClient) Delete(container *Container) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
