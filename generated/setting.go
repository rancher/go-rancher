package client

const (
	CONTAINER_TYPE = "setting"
)

type Setting struct {
	Resource
    
    Id int `json:"Id,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    Value string `json:"Value,omitempty"`
    
}

type SettingCollection struct {
	Collection
	Data []Setting `json:"data,omitempty"`
}

type SettingClient struct {
	rancherClient *RancherClient
}

type SettingOperations interface {
	List(opts *ListOpts) (*SettingCollection, error)
	Create(opts *Setting) (*Setting, error)
	Update(existing *Setting, updates interface{}) (*Setting, error)
	ById(id string) (*Setting, error)
	Delete(container *Setting) error
}

func newSettingClient(rancherClient *RancherClient) *SettingClient {
	return &SettingClient{
		rancherClient: rancherClient,
	}
}

func (self *SettingClient) Create(container *Setting) (*Setting, error) {
	resp := &Setting{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *SettingClient) Update(existing *Setting, updates interface{}) (*Setting, error) {
	resp := &Setting{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *SettingClient) List(opts *ListOpts) (*SettingCollection, error) {
	resp := &SettingCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *SettingClient) ById(id string) (*Setting, error) {
	resp := &Setting{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *SettingClient) Delete(container *Setting) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
