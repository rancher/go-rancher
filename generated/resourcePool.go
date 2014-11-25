package client

const (
	CONTAINER_TYPE = "resourcePool"
)

type ResourcePool struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Item string `json:"Item,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    OwnerId int `json:"OwnerId,omitempty"`
    
    OwnerType string `json:"OwnerType,omitempty"`
    
    PoolId int `json:"PoolId,omitempty"`
    
    PoolType string `json:"PoolType,omitempty"`
    
    Qualifier string `json:"Qualifier,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State string `json:"State,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type ResourcePoolCollection struct {
	Collection
	Data []ResourcePool `json:"data,omitempty"`
}

type ResourcePoolClient struct {
	rancherClient *RancherClient
}

type ResourcePoolOperations interface {
	List(opts *ListOpts) (*ResourcePoolCollection, error)
	Create(opts *ResourcePool) (*ResourcePool, error)
	Update(existing *ResourcePool, updates interface{}) (*ResourcePool, error)
	ById(id string) (*ResourcePool, error)
	Delete(container *ResourcePool) error
}

func newResourcePoolClient(rancherClient *RancherClient) *ResourcePoolClient {
	return &ResourcePoolClient{
		rancherClient: rancherClient,
	}
}

func (self *ResourcePoolClient) Create(container *ResourcePool) (*ResourcePool, error) {
	resp := &ResourcePool{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ResourcePoolClient) Update(existing *ResourcePool, updates interface{}) (*ResourcePool, error) {
	resp := &ResourcePool{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ResourcePoolClient) List(opts *ListOpts) (*ResourcePoolCollection, error) {
	resp := &ResourcePoolCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ResourcePoolClient) ById(id string) (*ResourcePool, error) {
	resp := &ResourcePool{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ResourcePoolClient) Delete(container *ResourcePool) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
