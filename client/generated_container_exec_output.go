package client

const (
	CONTAINER_EXEC_OUTPUT_TYPE = "containerExecOutput"
)

type ContainerExecOutput struct {
	Resource
    
    Token string `json:"token,omitempty"`
    
    Url string `json:"url,omitempty"`
    
}

type ContainerExecOutputCollection struct {
	Collection
	Data []ContainerExecOutput `json:"data,omitempty"`
}

type ContainerExecOutputClient struct {
	rancherClient *RancherClient
}

type ContainerExecOutputOperations interface {
	List(opts *ListOpts) (*ContainerExecOutputCollection, error)
	Create(opts *ContainerExecOutput) (*ContainerExecOutput, error)
	Update(existing *ContainerExecOutput, updates interface{}) (*ContainerExecOutput, error)
	ById(id string) (*ContainerExecOutput, error)
	Delete(container *ContainerExecOutput) error
}

func newContainerExecOutputClient(rancherClient *RancherClient) *ContainerExecOutputClient {
	return &ContainerExecOutputClient{
		rancherClient: rancherClient,
	}
}

func (c *ContainerExecOutputClient) Create(container *ContainerExecOutput) (*ContainerExecOutput, error) {
	resp := &ContainerExecOutput{}
	err := c.rancherClient.doCreate(CONTAINER_EXEC_OUTPUT_TYPE, container, resp)
	return resp, err
}

func (c *ContainerExecOutputClient) Update(existing *ContainerExecOutput, updates interface{}) (*ContainerExecOutput, error) {
	resp := &ContainerExecOutput{}
	err := c.rancherClient.doUpdate(CONTAINER_EXEC_OUTPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ContainerExecOutputClient) List(opts *ListOpts) (*ContainerExecOutputCollection, error) {
	resp := &ContainerExecOutputCollection{}
	err := c.rancherClient.doList(CONTAINER_EXEC_OUTPUT_TYPE, opts, resp)
	return resp, err
}

func (c *ContainerExecOutputClient) ById(id string) (*ContainerExecOutput, error) {
	resp := &ContainerExecOutput{}
	err := c.rancherClient.doById(CONTAINER_EXEC_OUTPUT_TYPE, id, resp)
	return resp, err
}

func (c *ContainerExecOutputClient) Delete(container *ContainerExecOutput) error {
	return c.rancherClient.doResourceDelete(CONTAINER_EXEC_OUTPUT_TYPE, &container.Resource)
}
