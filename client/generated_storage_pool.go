package client

const (
	STORAGE_POOL_TYPE = "storagePool"
)

type StoragePool struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    AgentId string `json:"agentId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    External bool `json:"external,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    ZoneId string `json:"zoneId,omitempty"`
    
}

type StoragePoolCollection struct {
	Collection
	Data []StoragePool `json:"data,omitempty"`
}

type StoragePoolClient struct {
	rancherClient *RancherClient
}

type StoragePoolOperations interface {
	List(opts *ListOpts) (*StoragePoolCollection, error)
	Create(opts *StoragePool) (*StoragePool, error)
	Update(existing *StoragePool, updates interface{}) (*StoragePool, error)
	ById(id string) (*StoragePool, error)
	Delete(container *StoragePool) error
}

func newStoragePoolClient(rancherClient *RancherClient) *StoragePoolClient {
	return &StoragePoolClient{
		rancherClient: rancherClient,
	}
}

func (c *StoragePoolClient) Create(container *StoragePool) (*StoragePool, error) {
	resp := &StoragePool{}
	err := c.rancherClient.doCreate(STORAGE_POOL_TYPE, container, resp)
	return resp, err
}

func (c *StoragePoolClient) Update(existing *StoragePool, updates interface{}) (*StoragePool, error) {
	resp := &StoragePool{}
	err := c.rancherClient.doUpdate(STORAGE_POOL_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *StoragePoolClient) List(opts *ListOpts) (*StoragePoolCollection, error) {
	resp := &StoragePoolCollection{}
	err := c.rancherClient.doList(STORAGE_POOL_TYPE, opts, resp)
	return resp, err
}

func (c *StoragePoolClient) ById(id string) (*StoragePool, error) {
	resp := &StoragePool{}
	err := c.rancherClient.doById(STORAGE_POOL_TYPE, id, resp)
	return resp, err
}

func (c *StoragePoolClient) Delete(container *StoragePool) error {
	return c.rancherClient.doResourceDelete(STORAGE_POOL_TYPE, &container.Resource)
}
