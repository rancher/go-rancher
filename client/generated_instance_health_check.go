package client

const (
	INSTANCE_HEALTH_CHECK_TYPE = "instanceHealthCheck"
)

type InstanceHealthCheck struct {
	Resource

	HealthyThreshold int `json:"healthyThreshold,omitempty"`

	Interval int `json:"interval,omitempty"`

	Name string `json:"name,omitempty"`

	Port int `json:"port,omitempty"`

	RequestLine string `json:"requestLine,omitempty"`

	ResponseTimeout int `json:"responseTimeout,omitempty"`

	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty"`
}

type InstanceHealthCheckCollection struct {
	Collection
	Data []InstanceHealthCheck `json:"data,omitempty"`
}

type InstanceHealthCheckClient struct {
	rancherClient *RancherClient
}

type InstanceHealthCheckOperations interface {
	List(opts *ListOpts) (*InstanceHealthCheckCollection, error)
	Create(opts *InstanceHealthCheck) (*InstanceHealthCheck, error)
	Update(existing *InstanceHealthCheck, updates interface{}) (*InstanceHealthCheck, error)
	ById(id string) (*InstanceHealthCheck, error)
	Delete(container *InstanceHealthCheck) error
}

func newInstanceHealthCheckClient(rancherClient *RancherClient) *InstanceHealthCheckClient {
	return &InstanceHealthCheckClient{
		rancherClient: rancherClient,
	}
}

func (c *InstanceHealthCheckClient) Create(container *InstanceHealthCheck) (*InstanceHealthCheck, error) {
	resp := &InstanceHealthCheck{}
	err := c.rancherClient.doCreate(INSTANCE_HEALTH_CHECK_TYPE, container, resp)
	return resp, err
}

func (c *InstanceHealthCheckClient) Update(existing *InstanceHealthCheck, updates interface{}) (*InstanceHealthCheck, error) {
	resp := &InstanceHealthCheck{}
	err := c.rancherClient.doUpdate(INSTANCE_HEALTH_CHECK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *InstanceHealthCheckClient) List(opts *ListOpts) (*InstanceHealthCheckCollection, error) {
	resp := &InstanceHealthCheckCollection{}
	err := c.rancherClient.doList(INSTANCE_HEALTH_CHECK_TYPE, opts, resp)
	return resp, err
}

func (c *InstanceHealthCheckClient) ById(id string) (*InstanceHealthCheck, error) {
	resp := &InstanceHealthCheck{}
	err := c.rancherClient.doById(INSTANCE_HEALTH_CHECK_TYPE, id, resp)
	return resp, err
}

func (c *InstanceHealthCheckClient) Delete(container *InstanceHealthCheck) error {
	return c.rancherClient.doResourceDelete(INSTANCE_HEALTH_CHECK_TYPE, &container.Resource)
}
