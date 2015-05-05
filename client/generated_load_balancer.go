package client

const (
	LOAD_BALANCER_TYPE = "loadBalancer"
)

type LoadBalancer struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    GlobalLoadBalancerId string `json:"globalLoadBalancerId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    LoadBalancerConfigId string `json:"loadBalancerConfigId,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    ServiceId string `json:"serviceId,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    Weight int `json:"weight,omitempty"`
    
}

type LoadBalancerCollection struct {
	Collection
	Data []LoadBalancer `json:"data,omitempty"`
}

type LoadBalancerClient struct {
	rancherClient *RancherClient
}

type LoadBalancerOperations interface {
	List(opts *ListOpts) (*LoadBalancerCollection, error)
	Create(opts *LoadBalancer) (*LoadBalancer, error)
	Update(existing *LoadBalancer, updates interface{}) (*LoadBalancer, error)
	ById(id string) (*LoadBalancer, error)
	Delete(container *LoadBalancer) error
    
    ActionAddhost (*LoadBalancer, *AddRemoveLoadBalancerHostInput) (*LoadBalancer, error)
    
    
    ActionAddtarget (*LoadBalancer, *AddRemoveLoadBalancerTargetInput) (*LoadBalancer, error)
    
    
    ActionCreate (*LoadBalancer) (*LoadBalancer, error)
    
    
    ActionRemove (*LoadBalancer) (*LoadBalancer, error)
    
    
    ActionRemovehost (*LoadBalancer, *AddRemoveLoadBalancerHostInput) (*LoadBalancer, error)
    
    
    ActionRemovetarget (*LoadBalancer, *AddRemoveLoadBalancerTargetInput) (*LoadBalancer, error)
    
    
    ActionSethosts (*LoadBalancer, *SetLoadBalancerHostsInput) (*LoadBalancer, error)
    
    
    ActionSettargets (*LoadBalancer, *SetLoadBalancerTargetsInput) (*LoadBalancer, error)
    
    
    ActionUpdate (*LoadBalancer) (*LoadBalancer, error)
    
}

func newLoadBalancerClient(rancherClient *RancherClient) *LoadBalancerClient {
	return &LoadBalancerClient{
		rancherClient: rancherClient,
	}
}

func (c *LoadBalancerClient) Create(container *LoadBalancer) (*LoadBalancer, error) {
	resp := &LoadBalancer{}
	err := c.rancherClient.doCreate(LOAD_BALANCER_TYPE, container, resp)
	return resp, err
}

func (c *LoadBalancerClient) Update(existing *LoadBalancer, updates interface{}) (*LoadBalancer, error) {
	resp := &LoadBalancer{}
	err := c.rancherClient.doUpdate(LOAD_BALANCER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LoadBalancerClient) List(opts *ListOpts) (*LoadBalancerCollection, error) {
	resp := &LoadBalancerCollection{}
	err := c.rancherClient.doList(LOAD_BALANCER_TYPE, opts, resp)
	return resp, err
}

func (c *LoadBalancerClient) ById(id string) (*LoadBalancer, error) {
	resp := &LoadBalancer{}
	err := c.rancherClient.doById(LOAD_BALANCER_TYPE, id, resp)
	return resp, err
}

func (c *LoadBalancerClient) Delete(container *LoadBalancer) error {
	return c.rancherClient.doResourceDelete(LOAD_BALANCER_TYPE, &container.Resource)
}
    
func (c *LoadBalancerClient) ActionAddhost (resource *LoadBalancer, input *AddRemoveLoadBalancerHostInput) (*LoadBalancer, error) {
    
	resp := &LoadBalancer{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_TYPE, "addhost", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *LoadBalancerClient) ActionAddtarget (resource *LoadBalancer, input *AddRemoveLoadBalancerTargetInput) (*LoadBalancer, error) {
    
	resp := &LoadBalancer{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_TYPE, "addtarget", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *LoadBalancerClient) ActionCreate (resource *LoadBalancer) (*LoadBalancer, error) {
    
	resp := &LoadBalancer{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_TYPE, "create", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *LoadBalancerClient) ActionRemove (resource *LoadBalancer) (*LoadBalancer, error) {
    
	resp := &LoadBalancer{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_TYPE, "remove", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *LoadBalancerClient) ActionRemovehost (resource *LoadBalancer, input *AddRemoveLoadBalancerHostInput) (*LoadBalancer, error) {
    
	resp := &LoadBalancer{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_TYPE, "removehost", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *LoadBalancerClient) ActionRemovetarget (resource *LoadBalancer, input *AddRemoveLoadBalancerTargetInput) (*LoadBalancer, error) {
    
	resp := &LoadBalancer{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_TYPE, "removetarget", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *LoadBalancerClient) ActionSethosts (resource *LoadBalancer, input *SetLoadBalancerHostsInput) (*LoadBalancer, error) {
    
	resp := &LoadBalancer{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_TYPE, "sethosts", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *LoadBalancerClient) ActionSettargets (resource *LoadBalancer, input *SetLoadBalancerTargetsInput) (*LoadBalancer, error) {
    
	resp := &LoadBalancer{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_TYPE, "settargets", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *LoadBalancerClient) ActionUpdate (resource *LoadBalancer) (*LoadBalancer, error) {
    
	resp := &LoadBalancer{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_TYPE, "update", &resource.Resource, nil, resp)
    
	return resp, err
}
