package client

const (
	CONTAINER_TYPE = "schema"
)

type Schema struct {
	Resource
    
    CollectionActions map[json] `json:"CollectionActions,omitempty"`
    
    CollectionFields map[json] `json:"CollectionFields,omitempty"`
    
    CollectionFilters map[json] `json:"CollectionFilters,omitempty"`
    
    CollectionMethods array[string] `json:"CollectionMethods,omitempty"`
    
    IncludeableLinks array[string] `json:"IncludeableLinks,omitempty"`
    
    PluralName string `json:"PluralName,omitempty"`
    
    ResourceActions map[json] `json:"ResourceActions,omitempty"`
    
    ResourceFields map[json] `json:"ResourceFields,omitempty"`
    
    ResourceMethods array[string] `json:"ResourceMethods,omitempty"`
    
}

type SchemaCollection struct {
	Collection
	Data []Schema `json:"data,omitempty"`
}

type SchemaClient struct {
	rancherClient *RancherClient
}

type SchemaOperations interface {
	List(opts *ListOpts) (*SchemaCollection, error)
	Create(opts *Schema) (*Schema, error)
	Update(existing *Schema, updates interface{}) (*Schema, error)
	ById(id string) (*Schema, error)
	Delete(container *Schema) error
}

func newSchemaClient(rancherClient *RancherClient) *SchemaClient {
	return &SchemaClient{
		rancherClient: rancherClient,
	}
}

func (self *SchemaClient) Create(container *Schema) (*Schema, error) {
	resp := &Schema{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *SchemaClient) Update(existing *Schema, updates interface{}) (*Schema, error) {
	resp := &Schema{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *SchemaClient) List(opts *ListOpts) (*SchemaCollection, error) {
	resp := &SchemaCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *SchemaClient) ById(id string) (*Schema, error) {
	resp := &Schema{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *SchemaClient) Delete(container *Schema) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
