package client

const (
	CONTAINER_TYPE = "data"
)

type Data struct {
	Resource
    
    Id int `json:"Id,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    Value string `json:"Value,omitempty"`
    
    Visible boolean `json:"Visible,omitempty"`
    
}

type DataCollection struct {
	Collection
	Data []Data `json:"data,omitempty"`
}

type DataClient struct {
	rancherClient *RancherClient
}

type DataOperations interface {
	List(opts *ListOpts) (*DataCollection, error)
	Create(opts *Data) (*Data, error)
	Update(existing *Data, updates interface{}) (*Data, error)
	ById(id string) (*Data, error)
	Delete(container *Data) error
}

func newDataClient(rancherClient *RancherClient) *DataClient {
	return &DataClient{
		rancherClient: rancherClient,
	}
}

func (self *DataClient) Create(container *Data) (*Data, error) {
	resp := &Data{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *DataClient) Update(existing *Data, updates interface{}) (*Data, error) {
	resp := &Data{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *DataClient) List(opts *ListOpts) (*DataCollection, error) {
	resp := &DataCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *DataClient) ById(id string) (*Data, error) {
	resp := &Data{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *DataClient) Delete(container *Data) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
