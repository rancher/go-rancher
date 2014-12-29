package client

const (
	CONTAINER_TYPE = "image"
)

type Image struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Checksum string `json:"Checksum,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Format string `json:"Format,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    InstanceKind enum `json:"InstanceKind,omitempty"`
    
    IsPublic boolean `json:"IsPublic,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    PhysicalSizeMb int `json:"PhysicalSizeMb,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Url string `json:"Url,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
    VirtualSizeMb int `json:"VirtualSizeMb,omitempty"`
    
}

type ImageCollection struct {
	Collection
	Data []Image `json:"data,omitempty"`
}

type ImageClient struct {
	rancherClient *RancherClient
}

type ImageOperations interface {
	List(opts *ListOpts) (*ImageCollection, error)
	Create(opts *Image) (*Image, error)
	Update(existing *Image, updates interface{}) (*Image, error)
	ById(id string) (*Image, error)
	Delete(container *Image) error
}

func newImageClient(rancherClient *RancherClient) *ImageClient {
	return &ImageClient{
		rancherClient: rancherClient,
	}
}

func (self *ImageClient) Create(container *Image) (*Image, error) {
	resp := &Image{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ImageClient) Update(existing *Image, updates interface{}) (*Image, error) {
	resp := &Image{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ImageClient) List(opts *ListOpts) (*ImageCollection, error) {
	resp := &ImageCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ImageClient) ById(id string) (*Image, error) {
	resp := &Image{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ImageClient) Delete(container *Image) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
