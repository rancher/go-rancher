package client

const (
	CONTAINER_TYPE = "imageStoragePoolMap"
)

type ImageStoragePoolMap struct {
	Resource
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    ImageId string `json:"ImageId,omitempty"`
    
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

func (self *ImageStoragePoolMapClient) Create(container *ImageStoragePoolMap) (*ImageStoragePoolMap, error) {
	resp := &ImageStoragePoolMap{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ImageStoragePoolMapClient) Update(existing *ImageStoragePoolMap, updates interface{}) (*ImageStoragePoolMap, error) {
	resp := &ImageStoragePoolMap{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ImageStoragePoolMapClient) List(opts *ListOpts) (*ImageStoragePoolMapCollection, error) {
	resp := &ImageStoragePoolMapCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ImageStoragePoolMapClient) ById(id string) (*ImageStoragePoolMap, error) {
	resp := &ImageStoragePoolMap{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ImageStoragePoolMapClient) Delete(container *ImageStoragePoolMap) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
