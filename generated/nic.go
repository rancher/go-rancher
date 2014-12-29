package client

const (
	CONTAINER_TYPE = "nic"
)

type Nic struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    DeviceNumber int `json:"DeviceNumber,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    InstanceId string `json:"InstanceId,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    MacAddress string `json:"MacAddress,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkId string `json:"NetworkId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    SubnetId string `json:"SubnetId,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
    VnetId string `json:"VnetId,omitempty"`
    
}

type NicCollection struct {
	Collection
	Data []Nic `json:"data,omitempty"`
}

type NicClient struct {
	rancherClient *RancherClient
}

type NicOperations interface {
	List(opts *ListOpts) (*NicCollection, error)
	Create(opts *Nic) (*Nic, error)
	Update(existing *Nic, updates interface{}) (*Nic, error)
	ById(id string) (*Nic, error)
	Delete(container *Nic) error
}

func newNicClient(rancherClient *RancherClient) *NicClient {
	return &NicClient{
		rancherClient: rancherClient,
	}
}

func (self *NicClient) Create(container *Nic) (*Nic, error) {
	resp := &Nic{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *NicClient) Update(existing *Nic, updates interface{}) (*Nic, error) {
	resp := &Nic{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *NicClient) List(opts *ListOpts) (*NicCollection, error) {
	resp := &NicCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *NicClient) ById(id string) (*Nic, error) {
	resp := &Nic{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *NicClient) Delete(container *Nic) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
