package client

const (
	SNAPSHOT_TYPE = "snapshot"
)

type Snapshot struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    ManagedSnapshotUUID string `json:"managedSnapshotUUID,omitempty"`
    
    ManagedVolumeUUID string `json:"managedVolumeUUID,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    VolumeId string `json:"volumeId,omitempty"`
    
}

type SnapshotCollection struct {
	Collection
	Data []Snapshot `json:"data,omitempty"`
}

type SnapshotClient struct {
	rancherClient *RancherClient
}

type SnapshotOperations interface {
	List(opts *ListOpts) (*SnapshotCollection, error)
	Create(opts *Snapshot) (*Snapshot, error)
	Update(existing *Snapshot, updates interface{}) (*Snapshot, error)
	ById(id string) (*Snapshot, error)
	Delete(container *Snapshot) error
    
    ActionBackup (*Snapshot) (*Snapshot, error)
    
    
    ActionCreate (*Snapshot) (*Snapshot, error)
    
    
    ActionRemove (*Snapshot) (*Snapshot, error)
    
}

func newSnapshotClient(rancherClient *RancherClient) *SnapshotClient {
	return &SnapshotClient{
		rancherClient: rancherClient,
	}
}

func (c *SnapshotClient) Create(container *Snapshot) (*Snapshot, error) {
	resp := &Snapshot{}
	err := c.rancherClient.doCreate(SNAPSHOT_TYPE, container, resp)
	return resp, err
}

func (c *SnapshotClient) Update(existing *Snapshot, updates interface{}) (*Snapshot, error) {
	resp := &Snapshot{}
	err := c.rancherClient.doUpdate(SNAPSHOT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SnapshotClient) List(opts *ListOpts) (*SnapshotCollection, error) {
	resp := &SnapshotCollection{}
	err := c.rancherClient.doList(SNAPSHOT_TYPE, opts, resp)
	return resp, err
}

func (c *SnapshotClient) ById(id string) (*Snapshot, error) {
	resp := &Snapshot{}
	err := c.rancherClient.doById(SNAPSHOT_TYPE, id, resp)
	return resp, err
}

func (c *SnapshotClient) Delete(container *Snapshot) error {
	return c.rancherClient.doResourceDelete(SNAPSHOT_TYPE, &container.Resource)
}
    
func (c *SnapshotClient) ActionBackup (resource *Snapshot) (*Snapshot, error) {
    
	resp := &Snapshot{}
    
	err := c.rancherClient.doAction(SNAPSHOT_TYPE, "backup", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *SnapshotClient) ActionCreate (resource *Snapshot) (*Snapshot, error) {
    
	resp := &Snapshot{}
    
	err := c.rancherClient.doAction(SNAPSHOT_TYPE, "create", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *SnapshotClient) ActionRemove (resource *Snapshot) (*Snapshot, error) {
    
	resp := &Snapshot{}
    
	err := c.rancherClient.doAction(SNAPSHOT_TYPE, "remove", &resource.Resource, nil, resp)
    
	return resp, err
}
