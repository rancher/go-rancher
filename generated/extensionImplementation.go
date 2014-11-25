package client

const (
	CONTAINER_TYPE = "extensionImplementation"
)

type ExtensionImplementation struct {
	Resource
    
    ClassName string `json:"ClassName,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    Properties map[string] `json:"Properties,omitempty"`
    
}

type ExtensionImplementationCollection struct {
	Collection
	Data []ExtensionImplementation `json:"data,omitempty"`
}

type ExtensionImplementationClient struct {
	rancherClient *RancherClient
}

type ExtensionImplementationOperations interface {
	List(opts *ListOpts) (*ExtensionImplementationCollection, error)
	Create(opts *ExtensionImplementation) (*ExtensionImplementation, error)
	Update(existing *ExtensionImplementation, updates interface{}) (*ExtensionImplementation, error)
	ById(id string) (*ExtensionImplementation, error)
	Delete(container *ExtensionImplementation) error
}

func newExtensionImplementationClient(rancherClient *RancherClient) *ExtensionImplementationClient {
	return &ExtensionImplementationClient{
		rancherClient: rancherClient,
	}
}

func (self *ExtensionImplementationClient) Create(container *ExtensionImplementation) (*ExtensionImplementation, error) {
	resp := &ExtensionImplementation{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ExtensionImplementationClient) Update(existing *ExtensionImplementation, updates interface{}) (*ExtensionImplementation, error) {
	resp := &ExtensionImplementation{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ExtensionImplementationClient) List(opts *ListOpts) (*ExtensionImplementationCollection, error) {
	resp := &ExtensionImplementationCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ExtensionImplementationClient) ById(id string) (*ExtensionImplementation, error) {
	resp := &ExtensionImplementation{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ExtensionImplementationClient) Delete(container *ExtensionImplementation) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
