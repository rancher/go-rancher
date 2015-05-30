package client

const (
	DNS_SERVICE_TYPE = "dnsService"
)

type DnsService struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    EnvironmentId string `json:"environmentId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type DnsServiceCollection struct {
	Collection
	Data []DnsService `json:"data,omitempty"`
}

type DnsServiceClient struct {
	rancherClient *RancherClient
}

type DnsServiceOperations interface {
	List(opts *ListOpts) (*DnsServiceCollection, error)
	Create(opts *DnsService) (*DnsService, error)
	Update(existing *DnsService, updates interface{}) (*DnsService, error)
	ById(id string) (*DnsService, error)
	Delete(container *DnsService) error
}

func newDnsServiceClient(rancherClient *RancherClient) *DnsServiceClient {
	return &DnsServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *DnsServiceClient) Create(container *DnsService) (*DnsService, error) {
	resp := &DnsService{}
	err := c.rancherClient.doCreate(DNS_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *DnsServiceClient) Update(existing *DnsService, updates interface{}) (*DnsService, error) {
	resp := &DnsService{}
	err := c.rancherClient.doUpdate(DNS_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DnsServiceClient) List(opts *ListOpts) (*DnsServiceCollection, error) {
	resp := &DnsServiceCollection{}
	err := c.rancherClient.doList(DNS_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *DnsServiceClient) ById(id string) (*DnsService, error) {
	resp := &DnsService{}
	err := c.rancherClient.doById(DNS_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *DnsServiceClient) Delete(container *DnsService) error {
	return c.rancherClient.doResourceDelete(DNS_SERVICE_TYPE, &container.Resource)
}
