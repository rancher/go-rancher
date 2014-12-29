package client

const (
	CONTAINER_TYPE = "error"
)

type Error struct {
	Resource
    
    Code string `json:"Code,omitempty"`
    
    Detail string `json:"Detail,omitempty"`
    
    Message string `json:"Message,omitempty"`
    
    Status int `json:"Status,omitempty"`
    
}

type ErrorCollection struct {
	Collection
	Data []Error `json:"data,omitempty"`
}

type ErrorClient struct {
	rancherClient *RancherClient
}

type ErrorOperations interface {
	List(opts *ListOpts) (*ErrorCollection, error)
	Create(opts *Error) (*Error, error)
	Update(existing *Error, updates interface{}) (*Error, error)
	ById(id string) (*Error, error)
	Delete(container *Error) error
}

func newErrorClient(rancherClient *RancherClient) *ErrorClient {
	return &ErrorClient{
		rancherClient: rancherClient,
	}
}

func (self *ErrorClient) Create(container *Error) (*Error, error) {
	resp := &Error{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ErrorClient) Update(existing *Error, updates interface{}) (*Error, error) {
	resp := &Error{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ErrorClient) List(opts *ListOpts) (*ErrorCollection, error) {
	resp := &ErrorCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ErrorClient) ById(id string) (*Error, error) {
	resp := &Error{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ErrorClient) Delete(container *Error) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
