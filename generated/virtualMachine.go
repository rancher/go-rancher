package client

const (
	CONTAINER_TYPE = "virtualMachine"
)

type VirtualMachine struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    AgentId string `json:"AgentId,omitempty"`
    
    AllocationState string `json:"AllocationState,omitempty"`
    
    Compute int `json:"Compute,omitempty"`
    
    Count int `json:"Count,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    CredentialIds array[reference[credential]] `json:"CredentialIds,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Domain string `json:"Domain,omitempty"`
    
    FirstRunning date `json:"FirstRunning,omitempty"`
    
    Hostname string `json:"Hostname,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    ImageId string `json:"ImageId,omitempty"`
    
    ImageUuid string `json:"ImageUuid,omitempty"`
    
    InstanceTriggeredStop enum `json:"InstanceTriggeredStop,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    LibvirtVncAddress string `json:"LibvirtVncAddress,omitempty"`
    
    LibvirtVncPassword string `json:"LibvirtVncPassword,omitempty"`
    
    MemoryMb int `json:"MemoryMb,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkIds array[reference[network]] `json:"NetworkIds,omitempty"`
    
    Ports array[string] `json:"Ports,omitempty"`
    
    PrimaryAssociatedIpAddress string `json:"PrimaryAssociatedIpAddress,omitempty"`
    
    PrimaryIpAddress string `json:"PrimaryIpAddress,omitempty"`
    
    PublicIpAddressPoolId string `json:"PublicIpAddressPoolId,omitempty"`
    
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
    
    Userdata string `json:"Userdata,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
    ValidHostIds array[reference[host]] `json:"ValidHostIds,omitempty"`
    
    Vcpu int `json:"Vcpu,omitempty"`
    
    VnetIds array[reference[vnet]] `json:"VnetIds,omitempty"`
    
    ZoneId string `json:"ZoneId,omitempty"`
    
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

func (self *VirtualMachineClient) Create(container *VirtualMachine) (*VirtualMachine, error) {
	resp := &VirtualMachine{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *VirtualMachineClient) Update(existing *VirtualMachine, updates interface{}) (*VirtualMachine, error) {
	resp := &VirtualMachine{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *VirtualMachineClient) List(opts *ListOpts) (*VirtualMachineCollection, error) {
	resp := &VirtualMachineCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *VirtualMachineClient) ById(id string) (*VirtualMachine, error) {
	resp := &VirtualMachine{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *VirtualMachineClient) Delete(container *VirtualMachine) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
