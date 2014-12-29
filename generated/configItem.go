package client

const (
	CONTAINER_TYPE = "configItem"
)

type ConfigItem struct {
	Resource
    
    Id int `json:"Id,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    SourceVersion string `json:"SourceVersion,omitempty"`
    
}

type ConfigItemCollection struct {
	Collection
	Data []ConfigItem `json:"data,omitempty"`
}

type ConfigItemClient struct {
	rancherClient *RancherClient
}

type ConfigItemOperations interface {
	List(opts *ListOpts) (*ConfigItemCollection, error)
	Create(opts *ConfigItem) (*ConfigItem, error)
	Update(existing *ConfigItem, updates interface{}) (*ConfigItem, error)
	ById(id string) (*ConfigItem, error)
	Delete(container *ConfigItem) error
}

func newConfigItemClient(rancherClient *RancherClient) *ConfigItemClient {
	return &ConfigItemClient{
		rancherClient: rancherClient,
	}
}

func (self *ConfigItemClient) Create(container *ConfigItem) (*ConfigItem, error) {
	resp := &ConfigItem{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ConfigItemClient) Update(existing *ConfigItem, updates interface{}) (*ConfigItem, error) {
	resp := &ConfigItem{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ConfigItemClient) List(opts *ListOpts) (*ConfigItemCollection, error) {
	resp := &ConfigItemCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ConfigItemClient) ById(id string) (*ConfigItem, error) {
	resp := &ConfigItem{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ConfigItemClient) Delete(container *ConfigItem) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
