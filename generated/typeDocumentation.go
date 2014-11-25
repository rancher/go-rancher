package client

const (
	CONTAINER_TYPE = "typeDocumentation"
)

type TypeDocumentation struct {
	Resource
    
    Description string `json:"Description,omitempty"`
    
    Id string `json:"Id,omitempty"`
    
    ResourceFields map[fieldDocumentation] `json:"ResourceFields,omitempty"`
    
}

type TypeDocumentationCollection struct {
	Collection
	Data []TypeDocumentation `json:"data,omitempty"`
}

type TypeDocumentationClient struct {
	rancherClient *RancherClient
}

type TypeDocumentationOperations interface {
	List(opts *ListOpts) (*TypeDocumentationCollection, error)
	Create(opts *TypeDocumentation) (*TypeDocumentation, error)
	Update(existing *TypeDocumentation, updates interface{}) (*TypeDocumentation, error)
	ById(id string) (*TypeDocumentation, error)
	Delete(container *TypeDocumentation) error
}

func newTypeDocumentationClient(rancherClient *RancherClient) *TypeDocumentationClient {
	return &TypeDocumentationClient{
		rancherClient: rancherClient,
	}
}

func (self *TypeDocumentationClient) Create(container *TypeDocumentation) (*TypeDocumentation, error) {
	resp := &TypeDocumentation{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *TypeDocumentationClient) Update(existing *TypeDocumentation, updates interface{}) (*TypeDocumentation, error) {
	resp := &TypeDocumentation{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *TypeDocumentationClient) List(opts *ListOpts) (*TypeDocumentationCollection, error) {
	resp := &TypeDocumentationCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *TypeDocumentationClient) ById(id string) (*TypeDocumentation, error) {
	resp := &TypeDocumentation{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *TypeDocumentationClient) Delete(container *TypeDocumentation) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
