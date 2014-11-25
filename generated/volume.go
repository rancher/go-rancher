package client

const (
	CONTAINER_TYPE = "volume"
)

type Volume struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    AllocationState string `json:"AllocationState,omitempty"`
    
    AttachedState string `json:"AttachedState,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    DeviceNumber int `json:"DeviceNumber,omitempty"`
    
    Format string `json:"Format,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    ImageId string `json:"ImageId,omitempty"`
    
    InstanceId string `json:"InstanceId,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
    VirtualSizeMb int `json:"VirtualSizeMb,omitempty"`
    
    ZoneId string `json:"ZoneId,omitempty"`
    
}

type VolumeCollection struct {
	Collection
	Data []Volume `json:"data,omitempty"`
}

type VolumeClient struct {
	rancherClient *RancherClient
}

type VolumeOperations interface {
	List(opts *ListOpts) (*VolumeCollection, error)
	Create(opts *Volume) (*Volume, error)
	Update(existing *Volume, updates interface{}) (*Volume, error)
	ById(id string) (*Volume, error)
	Delete(container *Volume) error
}

func newVolumeClient(rancherClient *RancherClient) *VolumeClient {
	return &VolumeClient{
		rancherClient: rancherClient,
	}
}

func (self *VolumeClient) Create(container *Volume) (*Volume, error) {
	resp := &Volume{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *VolumeClient) Update(existing *Volume, updates interface{}) (*Volume, error) {
	resp := &Volume{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *VolumeClient) List(opts *ListOpts) (*VolumeCollection, error) {
	resp := &VolumeCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *VolumeClient) ById(id string) (*Volume, error) {
	resp := &Volume{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *VolumeClient) Delete(container *Volume) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
