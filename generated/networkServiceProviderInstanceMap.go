package client

const (
	CONTAINER_TYPE = "networkServiceProviderInstanceMap"
)

type NetworkServiceProviderInstanceMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    InstanceId string `json:"InstanceId,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkServiceProviderId string `json:"NetworkServiceProviderId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type NetworkServiceProviderInstanceMapCollection struct {
	Collection
	Data []NetworkServiceProviderInstanceMap `json:"data,omitempty"`
}

type NetworkServiceProviderInstanceMapClient struct {
	rancherClient *RancherClient
}

type NetworkServiceProviderInstanceMapOperations interface {
	List(opts *ListOpts) (*NetworkServiceProviderInstanceMapCollection, error)
	Create(opts *NetworkServiceProviderInstanceMap) (*NetworkServiceProviderInstanceMap, error)
	Update(existing *NetworkServiceProviderInstanceMap, updates interface{}) (*NetworkServiceProviderInstanceMap, error)
	ById(id string) (*NetworkServiceProviderInstanceMap, error)
	Delete(container *NetworkServiceProviderInstanceMap) error
}

func newNetworkServiceProviderInstanceMapClient(rancherClient *RancherClient) *NetworkServiceProviderInstanceMapClient {
	return &NetworkServiceProviderInstanceMapClient{
		rancherClient: rancherClient,
	}
}

func (self *NetworkServiceProviderInstanceMapClient) Create(container *NetworkServiceProviderInstanceMap) (*NetworkServiceProviderInstanceMap, error) {
	resp := &NetworkServiceProviderInstanceMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *NetworkServiceProviderInstanceMapClient) Update(existing *NetworkServiceProviderInstanceMap, updates interface{}) (*NetworkServiceProviderInstanceMap, error) {
	resp := &NetworkServiceProviderInstanceMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *NetworkServiceProviderInstanceMapClient) List(opts *ListOpts) (*NetworkServiceProviderInstanceMapCollection, error) {
	resp := &NetworkServiceProviderInstanceMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *NetworkServiceProviderInstanceMapClient) ById(id string) (*NetworkServiceProviderInstanceMap, error) {
	resp := &NetworkServiceProviderInstanceMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *NetworkServiceProviderInstanceMapClient) Delete(container *NetworkServiceProviderInstanceMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
