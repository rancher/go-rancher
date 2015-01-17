package client

const (
	HOST_ONLY_NETWORK_TYPE = "hostOnlyNetwork"
)

type HostOnlyNetwork struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Domain string `json:"domain,omitempty"`
    
    DynamicCreateVnet bool `json:"dynamicCreateVnet,omitempty"`
    
    HostVnetUri string `json:"hostVnetUri,omitempty"`
    
    IsPublic bool `json:"isPublic,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    MacPrefix string `json:"macPrefix,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type HostOnlyNetworkCollection struct {
	Collection
	Data []HostOnlyNetwork `json:"data,omitempty"`
}

type HostOnlyNetworkClient struct {
	rancherClient *RancherClient
}

type HostOnlyNetworkOperations interface {
	List(opts *ListOpts) (*HostOnlyNetworkCollection, error)
	Create(opts *HostOnlyNetwork) (*HostOnlyNetwork, error)
	Update(existing *HostOnlyNetwork, updates interface{}) (*HostOnlyNetwork, error)
	ById(id string) (*HostOnlyNetwork, error)
	Delete(container *HostOnlyNetwork) error
}

func newHostOnlyNetworkClient(rancherClient *RancherClient) *HostOnlyNetworkClient {
	return &HostOnlyNetworkClient{
		rancherClient: rancherClient,
	}
}

func (c *HostOnlyNetworkClient) Create(container *HostOnlyNetwork) (*HostOnlyNetwork, error) {
	resp := &HostOnlyNetwork{}
	err := c.rancherClient.doCreate(HOST_ONLY_NETWORK_TYPE, container, resp)
	return resp, err
}

func (c *HostOnlyNetworkClient) Update(existing *HostOnlyNetwork, updates interface{}) (*HostOnlyNetwork, error) {
	resp := &HostOnlyNetwork{}
	err := c.rancherClient.doUpdate(HOST_ONLY_NETWORK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HostOnlyNetworkClient) List(opts *ListOpts) (*HostOnlyNetworkCollection, error) {
	resp := &HostOnlyNetworkCollection{}
	err := c.rancherClient.doList(HOST_ONLY_NETWORK_TYPE, opts, resp)
	return resp, err
}

func (c *HostOnlyNetworkClient) ById(id string) (*HostOnlyNetwork, error) {
	resp := &HostOnlyNetwork{}
	err := c.rancherClient.doById(HOST_ONLY_NETWORK_TYPE, id, resp)
	return resp, err
}

func (c *HostOnlyNetworkClient) Delete(container *HostOnlyNetwork) error {
	return c.rancherClient.doResourceDelete(HOST_ONLY_NETWORK_TYPE, &container.Resource)
}
