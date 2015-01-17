package client

const (
	INSTANCE_TYPE = "instance"
)

type Instance struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    AgentId string `json:"agentId,omitempty"`
    
    AllocationState string `json:"allocationState,omitempty"`
    
    Compute int `json:"compute,omitempty"`
    
    Count int `json:"count,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    CredentialIds []string `json:"credentialIds,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Domain string `json:"domain,omitempty"`
    
    FirstRunning string `json:"firstRunning,omitempty"`
    
    Hostname string `json:"hostname,omitempty"`
    
    ImageId string `json:"imageId,omitempty"`
    
    ImageUuid string `json:"imageUuid,omitempty"`
    
    InstanceTriggeredStop string `json:"instanceTriggeredStop,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    MemoryMb int `json:"memoryMb,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkIds []string `json:"networkIds,omitempty"`
    
    PrimaryAssociatedIpAddress string `json:"primaryAssociatedIpAddress,omitempty"`
    
    PrimaryIpAddress string `json:"primaryIpAddress,omitempty"`
    
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
    
    Userdata string `json:"userdata,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    ValidHostIds []string `json:"validHostIds,omitempty"`
    
    VnetIds []string `json:"vnetIds,omitempty"`
    
    ZoneId string `json:"zoneId,omitempty"`
    
}

type InstanceCollection struct {
	Collection
	Data []Instance `json:"data,omitempty"`
}

type InstanceClient struct {
	rancherClient *RancherClient
}

type InstanceOperations interface {
	List(opts *ListOpts) (*InstanceCollection, error)
	Create(opts *Instance) (*Instance, error)
	Update(existing *Instance, updates interface{}) (*Instance, error)
	ById(id string) (*Instance, error)
	Delete(container *Instance) error
}

func newInstanceClient(rancherClient *RancherClient) *InstanceClient {
	return &InstanceClient{
		rancherClient: rancherClient,
	}
}

func (c *InstanceClient) Create(container *Instance) (*Instance, error) {
	resp := &Instance{}
	err := c.rancherClient.doCreate(INSTANCE_TYPE, container, resp)
	return resp, err
}

func (c *InstanceClient) Update(existing *Instance, updates interface{}) (*Instance, error) {
	resp := &Instance{}
	err := c.rancherClient.doUpdate(INSTANCE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *InstanceClient) List(opts *ListOpts) (*InstanceCollection, error) {
	resp := &InstanceCollection{}
	err := c.rancherClient.doList(INSTANCE_TYPE, opts, resp)
	return resp, err
}

func (c *InstanceClient) ById(id string) (*Instance, error) {
	resp := &Instance{}
	err := c.rancherClient.doById(INSTANCE_TYPE, id, resp)
	return resp, err
}

func (c *InstanceClient) Delete(container *Instance) error {
	return c.rancherClient.doResourceDelete(INSTANCE_TYPE, &container.Resource)
}
