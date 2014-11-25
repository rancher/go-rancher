package client

const (
	CONTAINER_TYPE = "network"
)

type Network struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Domain string `json:"Domain,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    IsPublic boolean `json:"IsPublic,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    MacPrefix string `json:"MacPrefix,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type NetworkCollection struct {
	Collection
	Data []Network `json:"data,omitempty"`
}

type NetworkClient struct {
	rancherClient *RancherClient
}

type NetworkOperations interface {
	List(opts *ListOpts) (*NetworkCollection, error)
	Create(opts *Network) (*Network, error)
	Update(existing *Network, updates interface{}) (*Network, error)
	ById(id string) (*Network, error)
	Delete(container *Network) error
}

func newNetworkClient(rancherClient *RancherClient) *NetworkClient {
	return &NetworkClient{
		rancherClient: rancherClient,
	}
}

func (self *NetworkClient) Create(container *Network) (*Network, error) {
	resp := &Network{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *NetworkClient) Update(existing *Network, updates interface{}) (*Network, error) {
	resp := &Network{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *NetworkClient) List(opts *ListOpts) (*NetworkCollection, error) {
	resp := &NetworkCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *NetworkClient) ById(id string) (*Network, error) {
	resp := &Network{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *NetworkClient) Delete(container *Network) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
