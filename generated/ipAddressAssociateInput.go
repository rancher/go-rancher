package client

const (
	CONTAINER_TYPE = "ipAddressAssociateInput"
)

type IpAddressAssociateInput struct {
	Resource
    
    IpAddressId string `json:"IpAddressId,omitempty"`
    
}

type IpAddressAssociateInputCollection struct {
	Collection
	Data []IpAddressAssociateInput `json:"data,omitempty"`
}

type IpAddressAssociateInputClient struct {
	rancherClient *RancherClient
}

type IpAddressAssociateInputOperations interface {
	List(opts *ListOpts) (*IpAddressAssociateInputCollection, error)
	Create(opts *IpAddressAssociateInput) (*IpAddressAssociateInput, error)
	Update(existing *IpAddressAssociateInput, updates interface{}) (*IpAddressAssociateInput, error)
	ById(id string) (*IpAddressAssociateInput, error)
	Delete(container *IpAddressAssociateInput) error
}

func newIpAddressAssociateInputClient(rancherClient *RancherClient) *IpAddressAssociateInputClient {
	return &IpAddressAssociateInputClient{
		rancherClient: rancherClient,
	}
}

func (self *IpAddressAssociateInputClient) Create(container *IpAddressAssociateInput) (*IpAddressAssociateInput, error) {
	resp := &IpAddressAssociateInput{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *IpAddressAssociateInputClient) Update(existing *IpAddressAssociateInput, updates interface{}) (*IpAddressAssociateInput, error) {
	resp := &IpAddressAssociateInput{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *IpAddressAssociateInputClient) List(opts *ListOpts) (*IpAddressAssociateInputCollection, error) {
	resp := &IpAddressAssociateInputCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *IpAddressAssociateInputClient) ById(id string) (*IpAddressAssociateInput, error) {
	resp := &IpAddressAssociateInput{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *IpAddressAssociateInputClient) Delete(container *IpAddressAssociateInput) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
