package client

const (
	STORAGE_POOL_HOST_MAP_TYPE = "storagePoolHostMap"
)

type StoragePoolHostMap struct {
	Resource
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    HostId string `json:"hostId,omitempty"`
    
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

func (c *StoragePoolHostMapClient) Create(container *StoragePoolHostMap) (*StoragePoolHostMap, error) {
	resp := &StoragePoolHostMap{}
	err := c.rancherClient.doCreate(STORAGE_POOL_HOST_MAP_TYPE, container, resp)
	return resp, err
}

func (c *StoragePoolHostMapClient) Update(existing *StoragePoolHostMap, updates interface{}) (*StoragePoolHostMap, error) {
	resp := &StoragePoolHostMap{}
	err := c.rancherClient.doUpdate(STORAGE_POOL_HOST_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *StoragePoolHostMapClient) List(opts *ListOpts) (*StoragePoolHostMapCollection, error) {
	resp := &StoragePoolHostMapCollection{}
	err := c.rancherClient.doList(STORAGE_POOL_HOST_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *StoragePoolHostMapClient) ById(id string) (*StoragePoolHostMap, error) {
	resp := &StoragePoolHostMap{}
	err := c.rancherClient.doById(STORAGE_POOL_HOST_MAP_TYPE, id, resp)
	return resp, err
}

func (c *StoragePoolHostMapClient) Delete(container *StoragePoolHostMap) error {
	return c.rancherClient.doResourceDelete(STORAGE_POOL_HOST_MAP_TYPE, &container.Resource)
}
