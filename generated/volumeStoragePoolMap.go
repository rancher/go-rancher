package client

const (
	CONTAINER_TYPE = "volumeStoragePoolMap"
)

type VolumeStoragePoolMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
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
    
    VolumeId string `json:"VolumeId,omitempty"`
    
}

type VolumeStoragePoolMapCollection struct {
	Collection
	Data []VolumeStoragePoolMap `json:"data,omitempty"`
}

type VolumeStoragePoolMapClient struct {
	rancherClient *RancherClient
}

type VolumeStoragePoolMapOperations interface {
	List(opts *ListOpts) (*VolumeStoragePoolMapCollection, error)
	Create(opts *VolumeStoragePoolMap) (*VolumeStoragePoolMap, error)
	Update(existing *VolumeStoragePoolMap, updates interface{}) (*VolumeStoragePoolMap, error)
	ById(id string) (*VolumeStoragePoolMap, error)
	Delete(container *VolumeStoragePoolMap) error
}

func newVolumeStoragePoolMapClient(rancherClient *RancherClient) *VolumeStoragePoolMapClient {
	return &VolumeStoragePoolMapClient{
		rancherClient: rancherClient,
	}
}

func (self *VolumeStoragePoolMapClient) Create(container *VolumeStoragePoolMap) (*VolumeStoragePoolMap, error) {
	resp := &VolumeStoragePoolMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *VolumeStoragePoolMapClient) Update(existing *VolumeStoragePoolMap, updates interface{}) (*VolumeStoragePoolMap, error) {
	resp := &VolumeStoragePoolMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *VolumeStoragePoolMapClient) List(opts *ListOpts) (*VolumeStoragePoolMapCollection, error) {
	resp := &VolumeStoragePoolMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *VolumeStoragePoolMapClient) ById(id string) (*VolumeStoragePoolMap, error) {
	resp := &VolumeStoragePoolMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *VolumeStoragePoolMapClient) Delete(container *VolumeStoragePoolMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
