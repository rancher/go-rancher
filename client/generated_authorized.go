package client

const (
	AUTHORIZED_TYPE = "authorized"
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

func (c *AuthorizedClient) Create(container *Authorized) (*Authorized, error) {
	resp := &Authorized{}
	err := c.rancherClient.doCreate(AUTHORIZED_TYPE, container, resp)
	return resp, err
}

func (c *AuthorizedClient) Update(existing *Authorized, updates interface{}) (*Authorized, error) {
	resp := &Authorized{}
	err := c.rancherClient.doUpdate(AUTHORIZED_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AuthorizedClient) List(opts *ListOpts) (*AuthorizedCollection, error) {
	resp := &AuthorizedCollection{}
	err := c.rancherClient.doList(AUTHORIZED_TYPE, opts, resp)
	return resp, err
}

func (c *AuthorizedClient) ById(id string) (*Authorized, error) {
	resp := &Authorized{}
	err := c.rancherClient.doById(AUTHORIZED_TYPE, id, resp)
	return resp, err
}

func (c *AuthorizedClient) Delete(container *Authorized) error {
	return c.rancherClient.doResourceDelete(AUTHORIZED_TYPE, &container.Resource)
}
