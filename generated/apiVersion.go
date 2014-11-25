package client

const (
	CONTAINER_TYPE = "apiVersion"
)

type ApiVersion struct {
	Resource
    
}

type ApiVersionCollection struct {
	Collection
	Data []ApiVersion `json:"data,omitempty"`
}

type ApiVersionClient struct {
	rancherClient *RancherClient
}

type ApiVersionOperations interface {
	List(opts *ListOpts) (*ApiVersionCollection, error)
	Create(opts *ApiVersion) (*ApiVersion, error)
	Update(existing *ApiVersion, updates interface{}) (*ApiVersion, error)
	ById(id string) (*ApiVersion, error)
	Delete(container *ApiVersion) error
}

func newApiVersionClient(rancherClient *RancherClient) *ApiVersionClient {
	return &ApiVersionClient{
		rancherClient: rancherClient,
	}
}

func (self *ApiVersionClient) Create(container *ApiVersion) (*ApiVersion, error) {
	resp := &ApiVersion{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ApiVersionClient) Update(existing *ApiVersion, updates interface{}) (*ApiVersion, error) {
	resp := &ApiVersion{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ApiVersionClient) List(opts *ListOpts) (*ApiVersionCollection, error) {
	resp := &ApiVersionCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ApiVersionClient) ById(id string) (*ApiVersion, error) {
	resp := &ApiVersion{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ApiVersionClient) Delete(container *ApiVersion) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
