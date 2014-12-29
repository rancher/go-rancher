package client

const (
	CONTAINER_TYPE = "statsAccess"
)

type StatsAccess struct {
	Resource
    
    Token string `json:"Token,omitempty"`
    
    Url string `json:"Url,omitempty"`
    
}

type StatsAccessCollection struct {
	Collection
	Data []StatsAccess `json:"data,omitempty"`
}

type StatsAccessClient struct {
	rancherClient *RancherClient
}

type StatsAccessOperations interface {
	List(opts *ListOpts) (*StatsAccessCollection, error)
	Create(opts *StatsAccess) (*StatsAccess, error)
	Update(existing *StatsAccess, updates interface{}) (*StatsAccess, error)
	ById(id string) (*StatsAccess, error)
	Delete(container *StatsAccess) error
}

func newStatsAccessClient(rancherClient *RancherClient) *StatsAccessClient {
	return &StatsAccessClient{
		rancherClient: rancherClient,
	}
}

func (self *StatsAccessClient) Create(container *StatsAccess) (*StatsAccess, error) {
	resp := &StatsAccess{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *StatsAccessClient) Update(existing *StatsAccess, updates interface{}) (*StatsAccess, error) {
	resp := &StatsAccess{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *StatsAccessClient) List(opts *ListOpts) (*StatsAccessCollection, error) {
	resp := &StatsAccessCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *StatsAccessClient) ById(id string) (*StatsAccess, error) {
	resp := &StatsAccess{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *StatsAccessClient) Delete(container *StatsAccess) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
