package client

const (
	CONTAINER_TYPE = "fieldDocumentation"
)

type FieldDocumentation struct {
	Resource
    
    Description string `json:"Description,omitempty"`
    
    Placeholder string `json:"Placeholder,omitempty"`
    
}

type FieldDocumentationCollection struct {
	Collection
	Data []FieldDocumentation `json:"data,omitempty"`
}

type FieldDocumentationClient struct {
	rancherClient *RancherClient
}

type FieldDocumentationOperations interface {
	List(opts *ListOpts) (*FieldDocumentationCollection, error)
	Create(opts *FieldDocumentation) (*FieldDocumentation, error)
	Update(existing *FieldDocumentation, updates interface{}) (*FieldDocumentation, error)
	ById(id string) (*FieldDocumentation, error)
	Delete(container *FieldDocumentation) error
}

func newFieldDocumentationClient(rancherClient *RancherClient) *FieldDocumentationClient {
	return &FieldDocumentationClient{
		rancherClient: rancherClient,
	}
}

func (self *FieldDocumentationClient) Create(container *FieldDocumentation) (*FieldDocumentation, error) {
	resp := &FieldDocumentation{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *FieldDocumentationClient) Update(existing *FieldDocumentation, updates interface{}) (*FieldDocumentation, error) {
	resp := &FieldDocumentation{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *FieldDocumentationClient) List(opts *ListOpts) (*FieldDocumentationCollection, error) {
	resp := &FieldDocumentationCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *FieldDocumentationClient) ById(id string) (*FieldDocumentation, error) {
	resp := &FieldDocumentation{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *FieldDocumentationClient) Delete(container *FieldDocumentation) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
