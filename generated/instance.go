package client

const (
	CONTAINER_TYPE = "instance"
)

type Instance struct {
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
    
    MemoryMb int `json:"MemoryMb,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkIds array[reference[network]] `json:"NetworkIds,omitempty"`
    
    Ports array[string] `json:"Ports,omitempty"`
    
    PrimaryAssociatedIpAddress string `json:"PrimaryAssociatedIpAddress,omitempty"`
    
    PrimaryIpAddress string `json:"PrimaryIpAddress,omitempty"`
    
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
    
    VnetIds array[reference[vnet]] `json:"VnetIds,omitempty"`
    
    ZoneId string `json:"ZoneId,omitempty"`
    
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

func (self *InstanceClient) Create(container *Instance) (*Instance, error) {
	resp := &Instance{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *InstanceClient) Update(existing *Instance, updates interface{}) (*Instance, error) {
	resp := &Instance{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *InstanceClient) List(opts *ListOpts) (*InstanceCollection, error) {
	resp := &InstanceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *InstanceClient) ById(id string) (*Instance, error) {
	resp := &Instance{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *InstanceClient) Delete(container *Instance) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
