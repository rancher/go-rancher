package client

const (
	RESOURCE_POOL_TYPE = "resourcePool"
)

type ResourcePool struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Item string `json:"item,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    OwnerId int `json:"ownerId,omitempty"`
    
    OwnerType string `json:"ownerType,omitempty"`
    
    PoolId int `json:"poolId,omitempty"`
    
    PoolType string `json:"poolType,omitempty"`
    
    Qualifier string `json:"qualifier,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
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

func (c *ResourcePoolClient) Create(container *ResourcePool) (*ResourcePool, error) {
	resp := &ResourcePool{}
	err := c.rancherClient.doCreate(RESOURCE_POOL_TYPE, container, resp)
	return resp, err
}

func (c *ResourcePoolClient) Update(existing *ResourcePool, updates interface{}) (*ResourcePool, error) {
	resp := &ResourcePool{}
	err := c.rancherClient.doUpdate(RESOURCE_POOL_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ResourcePoolClient) List(opts *ListOpts) (*ResourcePoolCollection, error) {
	resp := &ResourcePoolCollection{}
	err := c.rancherClient.doList(RESOURCE_POOL_TYPE, opts, resp)
	return resp, err
}

func (c *ResourcePoolClient) ById(id string) (*ResourcePool, error) {
	resp := &ResourcePool{}
	err := c.rancherClient.doById(RESOURCE_POOL_TYPE, id, resp)
	return resp, err
}

func (c *ResourcePoolClient) Delete(container *ResourcePool) error {
	return c.rancherClient.doResourceDelete(RESOURCE_POOL_TYPE, &container.Resource)
}
