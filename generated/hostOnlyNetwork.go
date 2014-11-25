package client

const (
	CONTAINER_TYPE = "hostOnlyNetwork"
)

type HostOnlyNetwork struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Domain string `json:"Domain,omitempty"`
    
    DynamicCreateVnet boolean `json:"DynamicCreateVnet,omitempty"`
    
    HostVnetUri string `json:"HostVnetUri,omitempty"`
    
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

type HostOnlyNetworkCollection struct {
	Collection
	Data []HostOnlyNetwork `json:"data,omitempty"`
}

type HostOnlyNetworkClient struct {
	rancherClient *RancherClient
}

type HostOnlyNetworkOperations interface {
	List(opts *ListOpts) (*HostOnlyNetworkCollection, error)
	Create(opts *HostOnlyNetwork) (*HostOnlyNetwork, error)
	Update(existing *HostOnlyNetwork, updates interface{}) (*HostOnlyNetwork, error)
	ById(id string) (*HostOnlyNetwork, error)
	Delete(container *HostOnlyNetwork) error
}

func newHostOnlyNetworkClient(rancherClient *RancherClient) *HostOnlyNetworkClient {
	return &HostOnlyNetworkClient{
		rancherClient: rancherClient,
	}
}

func (self *HostOnlyNetworkClient) Create(container *HostOnlyNetwork) (*HostOnlyNetwork, error) {
	resp := &HostOnlyNetwork{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *HostOnlyNetworkClient) Update(existing *HostOnlyNetwork, updates interface{}) (*HostOnlyNetwork, error) {
	resp := &HostOnlyNetwork{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *HostOnlyNetworkClient) List(opts *ListOpts) (*HostOnlyNetworkCollection, error) {
	resp := &HostOnlyNetworkCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *HostOnlyNetworkClient) ById(id string) (*HostOnlyNetwork, error) {
	resp := &HostOnlyNetwork{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *HostOnlyNetworkClient) Delete(container *HostOnlyNetwork) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
