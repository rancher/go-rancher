package client

const (
	CONTAINER_TYPE = "databasechangelog"
)

type Databasechangelog struct {
	Resource
    
    Author string `json:"Author,omitempty"`
    
    Comments string `json:"Comments,omitempty"`
    
    Dateexecuted date `json:"Dateexecuted,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Exectype string `json:"Exectype,omitempty"`
    
    Filename string `json:"Filename,omitempty"`
    
    Id string `json:"Id,omitempty"`
    
    Liquibase string `json:"Liquibase,omitempty"`
    
    Md5sum string `json:"Md5sum,omitempty"`
    
    Orderexecuted int `json:"Orderexecuted,omitempty"`
    
    Tag string `json:"Tag,omitempty"`
    
}

type DatabasechangelogCollection struct {
	Collection
	Data []Databasechangelog `json:"data,omitempty"`
}

type DatabasechangelogClient struct {
	rancherClient *RancherClient
}

type DatabasechangelogOperations interface {
	List(opts *ListOpts) (*DatabasechangelogCollection, error)
	Create(opts *Databasechangelog) (*Databasechangelog, error)
	Update(existing *Databasechangelog, updates interface{}) (*Databasechangelog, error)
	ById(id string) (*Databasechangelog, error)
	Delete(container *Databasechangelog) error
}

func newDatabasechangelogClient(rancherClient *RancherClient) *DatabasechangelogClient {
	return &DatabasechangelogClient{
		rancherClient: rancherClient,
	}
}

func (self *DatabasechangelogClient) Create(container *Databasechangelog) (*Databasechangelog, error) {
	resp := &Databasechangelog{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *DatabasechangelogClient) Update(existing *Databasechangelog, updates interface{}) (*Databasechangelog, error) {
	resp := &Databasechangelog{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *DatabasechangelogClient) List(opts *ListOpts) (*DatabasechangelogCollection, error) {
	resp := &DatabasechangelogCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *DatabasechangelogClient) ById(id string) (*Databasechangelog, error) {
	resp := &Databasechangelog{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *DatabasechangelogClient) Delete(container *Databasechangelog) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
