package client

const (
	PUBLIC_ENDPOINT_TYPE = "publicEndpoint"
)

type PublicEndpoint struct {
	Resource

	AgentIpAddress string `json:"agentIpAddress,omitempty" yaml:"agent_ip_address,omitempty"`

	BindIpAddress string `json:"bindIpAddress,omitempty" yaml:"bind_ip_address,omitempty"`

	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`

	InstanceId string `json:"instanceId,omitempty" yaml:"instance_id,omitempty"`

	IpAddress string `json:"ipAddress,omitempty" yaml:"ip_address,omitempty"`

	PrivatePort int64 `json:"privatePort,omitempty" yaml:"private_port,omitempty"`

	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	PublicPort int64 `json:"publicPort,omitempty" yaml:"public_port,omitempty"`

	ServiceId string `json:"serviceId,omitempty" yaml:"service_id,omitempty"`
}

type PublicEndpointCollection struct {
	Collection
	Data   []PublicEndpoint `json:"data,omitempty"`
	client *PublicEndpointClient
}

type PublicEndpointClient struct {
	rancherClient *RancherClient
}

type PublicEndpointOperations interface {
	List(opts *ListOpts) (*PublicEndpointCollection, error)
	Create(opts *PublicEndpoint) (*PublicEndpoint, error)
	Update(existing *PublicEndpoint, updates interface{}) (*PublicEndpoint, error)
	ById(id string) (*PublicEndpoint, error)
	Delete(container *PublicEndpoint) error
}

func newPublicEndpointClient(rancherClient *RancherClient) *PublicEndpointClient {
	return &PublicEndpointClient{
		rancherClient: rancherClient,
	}
}

func (c *PublicEndpointClient) Create(container *PublicEndpoint) (*PublicEndpoint, error) {
	resp := &PublicEndpoint{}
	err := c.rancherClient.doCreate(PUBLIC_ENDPOINT_TYPE, container, resp)
	return resp, err
}

func (c *PublicEndpointClient) Update(existing *PublicEndpoint, updates interface{}) (*PublicEndpoint, error) {
	resp := &PublicEndpoint{}
	err := c.rancherClient.doUpdate(PUBLIC_ENDPOINT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PublicEndpointClient) List(opts *ListOpts) (*PublicEndpointCollection, error) {
	resp := &PublicEndpointCollection{}
	err := c.rancherClient.doList(PUBLIC_ENDPOINT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PublicEndpointCollection) Next() (*PublicEndpointCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PublicEndpointCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PublicEndpointClient) ById(id string) (*PublicEndpoint, error) {
	resp := &PublicEndpoint{}
	err := c.rancherClient.doById(PUBLIC_ENDPOINT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *PublicEndpointClient) Delete(container *PublicEndpoint) error {
	return c.rancherClient.doResourceDelete(PUBLIC_ENDPOINT_TYPE, &container.Resource)
}
