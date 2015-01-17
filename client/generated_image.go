package client

const (
	IMAGE_TYPE = "image"
)

type Image struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Checksum string `json:"checksum,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Format string `json:"format,omitempty"`
    
    InstanceKind string `json:"instanceKind,omitempty"`
    
    IsPublic bool `json:"isPublic,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    PhysicalSizeMb int `json:"physicalSizeMb,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Url string `json:"url,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    VirtualSizeMb int `json:"virtualSizeMb,omitempty"`
    
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

func (c *ImageClient) Create(container *Image) (*Image, error) {
	resp := &Image{}
	err := c.rancherClient.doCreate(IMAGE_TYPE, container, resp)
	return resp, err
}

func (c *ImageClient) Update(existing *Image, updates interface{}) (*Image, error) {
	resp := &Image{}
	err := c.rancherClient.doUpdate(IMAGE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ImageClient) List(opts *ListOpts) (*ImageCollection, error) {
	resp := &ImageCollection{}
	err := c.rancherClient.doList(IMAGE_TYPE, opts, resp)
	return resp, err
}

func (c *ImageClient) ById(id string) (*Image, error) {
	resp := &Image{}
	err := c.rancherClient.doById(IMAGE_TYPE, id, resp)
	return resp, err
}

func (c *ImageClient) Delete(container *Image) error {
	return c.rancherClient.doResourceDelete(IMAGE_TYPE, &container.Resource)
}
