package client

const (
	CONTAINER_TYPE = "authorized"
)

type Authorized struct {
	Resource
    
}

type AuthorizedCollection struct {
	Collection
	Data []Authorized `json:"data,omitempty"`
}

type AuthorizedClient struct {
	rancherClient *RancherClient
}

type AuthorizedOperations interface {
	List(opts *ListOpts) (*AuthorizedCollection, error)
	Create(opts *Authorized) (*Authorized, error)
	Update(existing *Authorized, updates interface{}) (*Authorized, error)
	ById(id string) (*Authorized, error)
	Delete(container *Authorized) error
}

func newAuthorizedClient(rancherClient *RancherClient) *AuthorizedClient {
	return &AuthorizedClient{
		rancherClient: rancherClient,
	}
}

func (self *AuthorizedClient) Create(container *Authorized) (*Authorized, error) {
	resp := &Authorized{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *AuthorizedClient) Update(existing *Authorized, updates interface{}) (*Authorized, error) {
	resp := &Authorized{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *AuthorizedClient) List(opts *ListOpts) (*AuthorizedCollection, error) {
	resp := &AuthorizedCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *AuthorizedClient) ById(id string) (*Authorized, error) {
	resp := &Authorized{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *AuthorizedClient) Delete(container *Authorized) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
