package client

const (
	CONTAINER_TYPE = "collection"
)

type Collection struct {
	Resource
    
    Actions map[json] `json:"Actions,omitempty"`
    
    CreateDefaults map[json] `json:"CreateDefaults,omitempty"`
    
    CreateTypes map[json] `json:"CreateTypes,omitempty"`
    
    Data array[resource] `json:"Data,omitempty"`
    
    Filters map[array[json]] `json:"Filters,omitempty"`
    
    Links map[json] `json:"Links,omitempty"`
    
    Pagination json `json:"Pagination,omitempty"`
    
    ResourceType string `json:"ResourceType,omitempty"`
    
    Sort json `json:"Sort,omitempty"`
    
    SortLinks map[json] `json:"SortLinks,omitempty"`
    
    Type string `json:"Type,omitempty"`
    
}

type CollectionCollection struct {
	Collection
	Data []Collection `json:"data,omitempty"`
}

type CollectionClient struct {
	rancherClient *RancherClient
}

type CollectionOperations interface {
	List(opts *ListOpts) (*CollectionCollection, error)
	Create(opts *Collection) (*Collection, error)
	Update(existing *Collection, updates interface{}) (*Collection, error)
	ById(id string) (*Collection, error)
	Delete(container *Collection) error
}

func newCollectionClient(rancherClient *RancherClient) *CollectionClient {
	return &CollectionClient{
		rancherClient: rancherClient,
	}
}

func (self *CollectionClient) Create(container *Collection) (*Collection, error) {
	resp := &Collection{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *CollectionClient) Update(existing *Collection, updates interface{}) (*Collection, error) {
	resp := &Collection{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *CollectionClient) List(opts *ListOpts) (*CollectionCollection, error) {
	resp := &CollectionCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *CollectionClient) ById(id string) (*Collection, error) {
	resp := &Collection{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *CollectionClient) Delete(container *Collection) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
