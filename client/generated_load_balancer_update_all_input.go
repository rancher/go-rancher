package client

const (
	LOAD_BALANCER_UPDATE_ALL_INPUT_TYPE = "loadBalancerUpdateAllInput"
)

type LoadBalancerUpdateAllInput struct {
	Resource
    
    HostIds []string `json:"hostIds,omitempty"`
    
    InstanceIds []string `json:"instanceIds,omitempty"`
    
    IpAddresses []string `json:"ipAddresses,omitempty"`
    
}

type LoadBalancerUpdateAllInputCollection struct {
	Collection
	Data []LoadBalancerUpdateAllInput `json:"data,omitempty"`
}

type LoadBalancerUpdateAllInputClient struct {
	rancherClient *RancherClient
}

type LoadBalancerUpdateAllInputOperations interface {
	List(opts *ListOpts) (*LoadBalancerUpdateAllInputCollection, error)
	Create(opts *LoadBalancerUpdateAllInput) (*LoadBalancerUpdateAllInput, error)
	Update(existing *LoadBalancerUpdateAllInput, updates interface{}) (*LoadBalancerUpdateAllInput, error)
	ById(id string) (*LoadBalancerUpdateAllInput, error)
	Delete(container *LoadBalancerUpdateAllInput) error
}

func newLoadBalancerUpdateAllInputClient(rancherClient *RancherClient) *LoadBalancerUpdateAllInputClient {
	return &LoadBalancerUpdateAllInputClient{
		rancherClient: rancherClient,
	}
}

func (c *LoadBalancerUpdateAllInputClient) Create(container *LoadBalancerUpdateAllInput) (*LoadBalancerUpdateAllInput, error) {
	resp := &LoadBalancerUpdateAllInput{}
	err := c.rancherClient.doCreate(LOAD_BALANCER_UPDATE_ALL_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *LoadBalancerUpdateAllInputClient) Update(existing *LoadBalancerUpdateAllInput, updates interface{}) (*LoadBalancerUpdateAllInput, error) {
	resp := &LoadBalancerUpdateAllInput{}
	err := c.rancherClient.doUpdate(LOAD_BALANCER_UPDATE_ALL_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LoadBalancerUpdateAllInputClient) List(opts *ListOpts) (*LoadBalancerUpdateAllInputCollection, error) {
	resp := &LoadBalancerUpdateAllInputCollection{}
	err := c.rancherClient.doList(LOAD_BALANCER_UPDATE_ALL_INPUT_TYPE, opts, resp)
	return resp, err
}

func (c *LoadBalancerUpdateAllInputClient) ById(id string) (*LoadBalancerUpdateAllInput, error) {
	resp := &LoadBalancerUpdateAllInput{}
	err := c.rancherClient.doById(LOAD_BALANCER_UPDATE_ALL_INPUT_TYPE, id, resp)
	return resp, err
}

func (c *LoadBalancerUpdateAllInputClient) Delete(container *LoadBalancerUpdateAllInput) error {
	return c.rancherClient.doResourceDelete(LOAD_BALANCER_UPDATE_ALL_INPUT_TYPE, &container.Resource)
}
