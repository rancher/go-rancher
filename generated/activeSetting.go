package client

const (
	CONTAINER_TYPE = "activeSetting"
)

type ActiveSetting struct {
	Resource
    
    ActiveValue json `json:"ActiveValue,omitempty"`
    
    Id string `json:"Id,omitempty"`
    
    InDb boolean `json:"InDb,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    Source string `json:"Source,omitempty"`
    
    Value string `json:"Value,omitempty"`
    
}

type ActiveSettingCollection struct {
	Collection
	Data []ActiveSetting `json:"data,omitempty"`
}

type ActiveSettingClient struct {
	rancherClient *RancherClient
}

type ActiveSettingOperations interface {
	List(opts *ListOpts) (*ActiveSettingCollection, error)
	Create(opts *ActiveSetting) (*ActiveSetting, error)
	Update(existing *ActiveSetting, updates interface{}) (*ActiveSetting, error)
	ById(id string) (*ActiveSetting, error)
	Delete(container *ActiveSetting) error
}

func newActiveSettingClient(rancherClient *RancherClient) *ActiveSettingClient {
	return &ActiveSettingClient{
		rancherClient: rancherClient,
	}
}

func (self *ActiveSettingClient) Create(container *ActiveSetting) (*ActiveSetting, error) {
	resp := &ActiveSetting{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ActiveSettingClient) Update(existing *ActiveSetting, updates interface{}) (*ActiveSetting, error) {
	resp := &ActiveSetting{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ActiveSettingClient) List(opts *ListOpts) (*ActiveSettingCollection, error) {
	resp := &ActiveSettingCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ActiveSettingClient) ById(id string) (*ActiveSetting, error) {
	resp := &ActiveSetting{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ActiveSettingClient) Delete(container *ActiveSetting) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
