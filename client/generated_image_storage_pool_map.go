package client

const (
	IMAGE_STORAGE_POOL_MAP_TYPE = "imageStoragePoolMap"
)

type ImageStoragePoolMap struct {
	Resource
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    ImageId string `json:"imageId,omitempty"`
    
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

type ImageStoragePoolMapCollection struct {
	Collection
	Data []ImageStoragePoolMap `json:"data,omitempty"`
}

type ImageStoragePoolMapClient struct {
	rancherClient *RancherClient
}

type ImageStoragePoolMapOperations interface {
	List(opts *ListOpts) (*ImageStoragePoolMapCollection, error)
	Create(opts *ImageStoragePoolMap) (*ImageStoragePoolMap, error)
	Update(existing *ImageStoragePoolMap, updates interface{}) (*ImageStoragePoolMap, error)
	ById(id string) (*ImageStoragePoolMap, error)
	Delete(container *ImageStoragePoolMap) error
}

func newImageStoragePoolMapClient(rancherClient *RancherClient) *ImageStoragePoolMapClient {
	return &ImageStoragePoolMapClient{
		rancherClient: rancherClient,
	}
}

func (c *ImageStoragePoolMapClient) Create(container *ImageStoragePoolMap) (*ImageStoragePoolMap, error) {
	resp := &ImageStoragePoolMap{}
	err := c.rancherClient.doCreate(IMAGE_STORAGE_POOL_MAP_TYPE, container, resp)
	return resp, err
}

func (c *ImageStoragePoolMapClient) Update(existing *ImageStoragePoolMap, updates interface{}) (*ImageStoragePoolMap, error) {
	resp := &ImageStoragePoolMap{}
	err := c.rancherClient.doUpdate(IMAGE_STORAGE_POOL_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ImageStoragePoolMapClient) List(opts *ListOpts) (*ImageStoragePoolMapCollection, error) {
	resp := &ImageStoragePoolMapCollection{}
	err := c.rancherClient.doList(IMAGE_STORAGE_POOL_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *ImageStoragePoolMapClient) ById(id string) (*ImageStoragePoolMap, error) {
	resp := &ImageStoragePoolMap{}
	err := c.rancherClient.doById(IMAGE_STORAGE_POOL_MAP_TYPE, id, resp)
	return resp, err
}

func (c *ImageStoragePoolMapClient) Delete(container *ImageStoragePoolMap) error {
	return c.rancherClient.doResourceDelete(IMAGE_STORAGE_POOL_MAP_TYPE, &container.Resource)
}
