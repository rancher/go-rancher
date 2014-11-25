package client

const (
	CONTAINER_TYPE = "extensionPoint"
)

type ExtensionPoint struct {
	Resource
    
    ExcludeSetting string `json:"ExcludeSetting,omitempty"`
    
    Implementations array[extensionImplementation] `json:"Implementations,omitempty"`
    
    IncludeSetting string `json:"IncludeSetting,omitempty"`
    
    ListSetting string `json:"ListSetting,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
}

type ExtensionPointCollection struct {
	Collection
	Data []ExtensionPoint `json:"data,omitempty"`
}

type ExtensionPointClient struct {
	rancherClient *RancherClient
}

type ExtensionPointOperations interface {
	List(opts *ListOpts) (*ExtensionPointCollection, error)
	Create(opts *ExtensionPoint) (*ExtensionPoint, error)
	Update(existing *ExtensionPoint, updates interface{}) (*ExtensionPoint, error)
	ById(id string) (*ExtensionPoint, error)
	Delete(container *ExtensionPoint) error
}

func newExtensionPointClient(rancherClient *RancherClient) *ExtensionPointClient {
	return &ExtensionPointClient{
		rancherClient: rancherClient,
	}
}

func (self *ExtensionPointClient) Create(container *ExtensionPoint) (*ExtensionPoint, error) {
	resp := &ExtensionPoint{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ExtensionPointClient) Update(existing *ExtensionPoint, updates interface{}) (*ExtensionPoint, error) {
	resp := &ExtensionPoint{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ExtensionPointClient) List(opts *ListOpts) (*ExtensionPointCollection, error) {
	resp := &ExtensionPointCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ExtensionPointClient) ById(id string) (*ExtensionPoint, error) {
	resp := &ExtensionPoint{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ExtensionPointClient) Delete(container *ExtensionPoint) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
