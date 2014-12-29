package client

const (
	CONTAINER_TYPE = "ipAssociation"
)

type IpAssociation struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    ChildIpAddressId string `json:"ChildIpAddressId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    IpAddressId string `json:"IpAddressId,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    Role string `json:"Role,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
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

func (self *IpAssociationClient) Create(container *IpAssociation) (*IpAssociation, error) {
	resp := &IpAssociation{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *IpAssociationClient) Update(existing *IpAssociation, updates interface{}) (*IpAssociation, error) {
	resp := &IpAssociation{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *IpAssociationClient) List(opts *ListOpts) (*IpAssociationCollection, error) {
	resp := &IpAssociationCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *IpAssociationClient) ById(id string) (*IpAssociation, error) {
	resp := &IpAssociation{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *IpAssociationClient) Delete(container *IpAssociation) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
