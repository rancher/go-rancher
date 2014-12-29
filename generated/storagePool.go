package client

const (
	CONTAINER_TYPE = "storagePool"
)

type StoragePool struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    AgentId string `json:"AgentId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    External boolean `json:"External,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
    ZoneId string `json:"ZoneId,omitempty"`
    
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

func (self *StoragePoolClient) Create(container *StoragePool) (*StoragePool, error) {
	resp := &StoragePool{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *StoragePoolClient) Update(existing *StoragePool, updates interface{}) (*StoragePool, error) {
	resp := &StoragePool{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *StoragePoolClient) List(opts *ListOpts) (*StoragePoolCollection, error) {
	resp := &StoragePoolCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *StoragePoolClient) ById(id string) (*StoragePool, error) {
	resp := &StoragePool{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *StoragePoolClient) Delete(container *StoragePool) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
