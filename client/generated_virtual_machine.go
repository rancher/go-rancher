package client

const (
	VIRTUAL_MACHINE_TYPE = "virtualMachine"
)

type VirtualMachine struct {
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
    
    LibvirtVncAddress string `json:"libvirtVncAddress,omitempty"`
    
    LibvirtVncPassword string `json:"libvirtVncPassword,omitempty"`
    
    MemoryMb int `json:"memoryMb,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkIds []string `json:"networkIds,omitempty"`
    
    PrimaryAssociatedIpAddress string `json:"primaryAssociatedIpAddress,omitempty"`
    
    PrimaryIpAddress string `json:"primaryIpAddress,omitempty"`
    
    PublicIpAddressPoolId string `json:"publicIpAddressPoolId,omitempty"`
    
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
    
    Vcpu int `json:"vcpu,omitempty"`
    
    VnetIds []string `json:"vnetIds,omitempty"`
    
    ZoneId string `json:"zoneId,omitempty"`
    
}

type VirtualMachineCollection struct {
	Collection
	Data []VirtualMachine `json:"data,omitempty"`
}

type VirtualMachineClient struct {
	rancherClient *RancherClient
}

type VirtualMachineOperations interface {
	List(opts *ListOpts) (*VirtualMachineCollection, error)
	Create(opts *VirtualMachine) (*VirtualMachine, error)
	Update(existing *VirtualMachine, updates interface{}) (*VirtualMachine, error)
	ById(id string) (*VirtualMachine, error)
	Delete(container *VirtualMachine) error
}

func newVirtualMachineClient(rancherClient *RancherClient) *VirtualMachineClient {
	return &VirtualMachineClient{
		rancherClient: rancherClient,
	}
}

func (c *VirtualMachineClient) Create(container *VirtualMachine) (*VirtualMachine, error) {
	resp := &VirtualMachine{}
	err := c.rancherClient.doCreate(VIRTUAL_MACHINE_TYPE, container, resp)
	return resp, err
}

func (c *VirtualMachineClient) Update(existing *VirtualMachine, updates interface{}) (*VirtualMachine, error) {
	resp := &VirtualMachine{}
	err := c.rancherClient.doUpdate(VIRTUAL_MACHINE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VirtualMachineClient) List(opts *ListOpts) (*VirtualMachineCollection, error) {
	resp := &VirtualMachineCollection{}
	err := c.rancherClient.doList(VIRTUAL_MACHINE_TYPE, opts, resp)
	return resp, err
}

func (c *VirtualMachineClient) ById(id string) (*VirtualMachine, error) {
	resp := &VirtualMachine{}
	err := c.rancherClient.doById(VIRTUAL_MACHINE_TYPE, id, resp)
	return resp, err
}

func (c *VirtualMachineClient) Delete(container *VirtualMachine) error {
	return c.rancherClient.doResourceDelete(VIRTUAL_MACHINE_TYPE, &container.Resource)
}
