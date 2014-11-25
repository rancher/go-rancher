package client

const (
	CONTAINER_TYPE = "storagePoolHostMap"
)

type StoragePoolHostMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    HostId string `json:"HostId,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    StoragePoolId string `json:"StoragePoolId,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type StoragePoolHostMapCollection struct {
	Collection
	Data []StoragePoolHostMap `json:"data,omitempty"`
}

type StoragePoolHostMapClient struct {
	rancherClient *RancherClient
}

type StoragePoolHostMapOperations interface {
	List(opts *ListOpts) (*StoragePoolHostMapCollection, error)
	Create(opts *StoragePoolHostMap) (*StoragePoolHostMap, error)
	Update(existing *StoragePoolHostMap, updates interface{}) (*StoragePoolHostMap, error)
	ById(id string) (*StoragePoolHostMap, error)
	Delete(container *StoragePoolHostMap) error
}

func newStoragePoolHostMapClient(rancherClient *RancherClient) *StoragePoolHostMapClient {
	return &StoragePoolHostMapClient{
		rancherClient: rancherClient,
	}
}

func (self *StoragePoolHostMapClient) Create(container *StoragePoolHostMap) (*StoragePoolHostMap, error) {
	resp := &StoragePoolHostMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *StoragePoolHostMapClient) Update(existing *StoragePoolHostMap, updates interface{}) (*StoragePoolHostMap, error) {
	resp := &StoragePoolHostMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *StoragePoolHostMapClient) List(opts *ListOpts) (*StoragePoolHostMapCollection, error) {
	resp := &StoragePoolHostMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *StoragePoolHostMapClient) ById(id string) (*StoragePoolHostMap, error) {
	resp := &StoragePoolHostMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *StoragePoolHostMapClient) Delete(container *StoragePoolHostMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
