package client

const (
	DATA_TYPE = "data"
)

type Data struct {
	Resource
    
    Name string `json:"name,omitempty"`
    
    Value string `json:"value,omitempty"`
    
    Visible bool `json:"visible,omitempty"`
    
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

func (c *DataClient) Create(container *Data) (*Data, error) {
	resp := &Data{}
	err := c.rancherClient.doCreate(DATA_TYPE, container, resp)
	return resp, err
}

func (c *DataClient) Update(existing *Data, updates interface{}) (*Data, error) {
	resp := &Data{}
	err := c.rancherClient.doUpdate(DATA_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DataClient) List(opts *ListOpts) (*DataCollection, error) {
	resp := &DataCollection{}
	err := c.rancherClient.doList(DATA_TYPE, opts, resp)
	return resp, err
}

func (c *DataClient) ById(id string) (*Data, error) {
	resp := &Data{}
	err := c.rancherClient.doById(DATA_TYPE, id, resp)
	return resp, err
}

func (c *DataClient) Delete(container *Data) error {
	return c.rancherClient.doResourceDelete(DATA_TYPE, &container.Resource)
}
