package client

const (
	CONTAINER_TYPE = "databasechangeloglock"
)

type Databasechangeloglock struct {
	Resource
    
    Id int `json:"Id,omitempty"`
    
    Locked boolean `json:"Locked,omitempty"`
    
    Lockedby string `json:"Lockedby,omitempty"`
    
    Lockgranted date `json:"Lockgranted,omitempty"`
    
}

type DatabasechangeloglockCollection struct {
	Collection
	Data []Databasechangeloglock `json:"data,omitempty"`
}

type DatabasechangeloglockClient struct {
	rancherClient *RancherClient
}

type DatabasechangeloglockOperations interface {
	List(opts *ListOpts) (*DatabasechangeloglockCollection, error)
	Create(opts *Databasechangeloglock) (*Databasechangeloglock, error)
	Update(existing *Databasechangeloglock, updates interface{}) (*Databasechangeloglock, error)
	ById(id string) (*Databasechangeloglock, error)
	Delete(container *Databasechangeloglock) error
}

func newDatabasechangeloglockClient(rancherClient *RancherClient) *DatabasechangeloglockClient {
	return &DatabasechangeloglockClient{
		rancherClient: rancherClient,
	}
}

func (self *DatabasechangeloglockClient) Create(container *Databasechangeloglock) (*Databasechangeloglock, error) {
	resp := &Databasechangeloglock{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *DatabasechangeloglockClient) Update(existing *Databasechangeloglock, updates interface{}) (*Databasechangeloglock, error) {
	resp := &Databasechangeloglock{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *DatabasechangeloglockClient) List(opts *ListOpts) (*DatabasechangeloglockCollection, error) {
	resp := &DatabasechangeloglockCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *DatabasechangeloglockClient) ById(id string) (*Databasechangeloglock, error) {
	resp := &Databasechangeloglock{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *DatabasechangeloglockClient) Delete(container *Databasechangeloglock) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
