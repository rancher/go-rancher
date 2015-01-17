package client

const (
	IP_ASSOCIATION_TYPE = "ipAssociation"
)

type IpAssociation struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    ChildIpAddressId string `json:"childIpAddressId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    IpAddressId string `json:"ipAddressId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    Role string `json:"role,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type IpAssociationCollection struct {
	Collection
	Data []IpAssociation `json:"data,omitempty"`
}

type IpAssociationClient struct {
	rancherClient *RancherClient
}

type IpAssociationOperations interface {
	List(opts *ListOpts) (*IpAssociationCollection, error)
	Create(opts *IpAssociation) (*IpAssociation, error)
	Update(existing *IpAssociation, updates interface{}) (*IpAssociation, error)
	ById(id string) (*IpAssociation, error)
	Delete(container *IpAssociation) error
}

func newIpAssociationClient(rancherClient *RancherClient) *IpAssociationClient {
	return &IpAssociationClient{
		rancherClient: rancherClient,
	}
}

func (c *IpAssociationClient) Create(container *IpAssociation) (*IpAssociation, error) {
	resp := &IpAssociation{}
	err := c.rancherClient.doCreate(IP_ASSOCIATION_TYPE, container, resp)
	return resp, err
}

func (c *IpAssociationClient) Update(existing *IpAssociation, updates interface{}) (*IpAssociation, error) {
	resp := &IpAssociation{}
	err := c.rancherClient.doUpdate(IP_ASSOCIATION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *IpAssociationClient) List(opts *ListOpts) (*IpAssociationCollection, error) {
	resp := &IpAssociationCollection{}
	err := c.rancherClient.doList(IP_ASSOCIATION_TYPE, opts, resp)
	return resp, err
}

func (c *IpAssociationClient) ById(id string) (*IpAssociation, error) {
	resp := &IpAssociation{}
	err := c.rancherClient.doById(IP_ASSOCIATION_TYPE, id, resp)
	return resp, err
}

func (c *IpAssociationClient) Delete(container *IpAssociation) error {
	return c.rancherClient.doResourceDelete(IP_ASSOCIATION_TYPE, &container.Resource)
}
