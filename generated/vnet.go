package client

const (
	CONTAINER_TYPE = "vnet"
)

type Vnet struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkId string `json:"NetworkId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uri string `json:"Uri,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type VnetCollection struct {
	Collection
	Data []Vnet `json:"data,omitempty"`
}

type VnetClient struct {
	rancherClient *RancherClient
}

type VnetOperations interface {
	List(opts *ListOpts) (*VnetCollection, error)
	Create(opts *Vnet) (*Vnet, error)
	Update(existing *Vnet, updates interface{}) (*Vnet, error)
	ById(id string) (*Vnet, error)
	Delete(container *Vnet) error
}

func newVnetClient(rancherClient *RancherClient) *VnetClient {
	return &VnetClient{
		rancherClient: rancherClient,
	}
}

func (self *VnetClient) Create(container *Vnet) (*Vnet, error) {
	resp := &Vnet{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *VnetClient) Update(existing *Vnet, updates interface{}) (*Vnet, error) {
	resp := &Vnet{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *VnetClient) List(opts *ListOpts) (*VnetCollection, error) {
	resp := &VnetCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *VnetClient) ById(id string) (*Vnet, error) {
	resp := &Vnet{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *VnetClient) Delete(container *Vnet) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
