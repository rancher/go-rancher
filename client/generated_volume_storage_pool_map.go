package client

const (
	VOLUME_STORAGE_POOL_MAP_TYPE = "volumeStoragePoolMap"
)

type VolumeStoragePoolMap struct {
	Resource
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    StoragePoolId string `json:"storagePoolId,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    VolumeId string `json:"volumeId,omitempty"`
    
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

func (c *VolumeStoragePoolMapClient) Create(container *VolumeStoragePoolMap) (*VolumeStoragePoolMap, error) {
	resp := &VolumeStoragePoolMap{}
	err := c.rancherClient.doCreate(VOLUME_STORAGE_POOL_MAP_TYPE, container, resp)
	return resp, err
}

func (c *VolumeStoragePoolMapClient) Update(existing *VolumeStoragePoolMap, updates interface{}) (*VolumeStoragePoolMap, error) {
	resp := &VolumeStoragePoolMap{}
	err := c.rancherClient.doUpdate(VOLUME_STORAGE_POOL_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VolumeStoragePoolMapClient) List(opts *ListOpts) (*VolumeStoragePoolMapCollection, error) {
	resp := &VolumeStoragePoolMapCollection{}
	err := c.rancherClient.doList(VOLUME_STORAGE_POOL_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *VolumeStoragePoolMapClient) ById(id string) (*VolumeStoragePoolMap, error) {
	resp := &VolumeStoragePoolMap{}
	err := c.rancherClient.doById(VOLUME_STORAGE_POOL_MAP_TYPE, id, resp)
	return resp, err
}

func (c *VolumeStoragePoolMapClient) Delete(container *VolumeStoragePoolMap) error {
	return c.rancherClient.doResourceDelete(VOLUME_STORAGE_POOL_MAP_TYPE, &container.Resource)
}
