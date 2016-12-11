package catalog

const (
	OUTPUT_TYPE = "output"
)

type Output struct {
	Resource

	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}

type OutputCollection struct {
	Collection
	Data   []Output `json:"data,omitempty"`
	client *OutputClient
}

type OutputClient struct {
	rancherClient *RancherClient
}

type OutputOperations interface {
	List(opts *ListOpts) (*OutputCollection, error)
	Create(opts *Output) (*Output, error)
	Update(existing *Output, updates interface{}) (*Output, error)
	ById(id string) (*Output, error)
	Delete(container *Output) error
}

func newOutputClient(rancherClient *RancherClient) *OutputClient {
	return &OutputClient{
		rancherClient: rancherClient,
	}
}

func (c *OutputClient) Create(container *Output) (*Output, error) {
	resp := &Output{}
	err := c.rancherClient.doCreate(OUTPUT_TYPE, container, resp)
	return resp, err
}

func (c *OutputClient) Update(existing *Output, updates interface{}) (*Output, error) {
	resp := &Output{}
	err := c.rancherClient.doUpdate(OUTPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *OutputClient) List(opts *ListOpts) (*OutputCollection, error) {
	resp := &OutputCollection{}
	err := c.rancherClient.doList(OUTPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *OutputCollection) Next() (*OutputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &OutputCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *OutputClient) ById(id string) (*Output, error) {
	resp := &Output{}
	err := c.rancherClient.doById(OUTPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *OutputClient) Delete(container *Output) error {
	return c.rancherClient.doResourceDelete(OUTPUT_TYPE, &container.Resource)
}
