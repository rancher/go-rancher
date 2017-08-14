package client

const (
	MOUNT_POINT_TYPE = "mountPoint"
)

type MountPoint struct {
	Resource

	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	Driver string `json:"driver,omitempty" yaml:"driver,omitempty"`

	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	MountType string `json:"mountType,omitempty" yaml:"mount_type,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Propagation string `json:"propagation,omitempty" yaml:"propagation,omitempty"`

	Rw bool `json:"rw,omitempty" yaml:"rw,omitempty"`

	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}

type MountPointCollection struct {
	Collection
	Data   []MountPoint `json:"data,omitempty"`
	client *MountPointClient
}

type MountPointClient struct {
	rancherClient *RancherClient
}

type MountPointOperations interface {
	List(opts *ListOpts) (*MountPointCollection, error)
	Create(opts *MountPoint) (*MountPoint, error)
	Update(existing *MountPoint, updates interface{}) (*MountPoint, error)
	ById(id string) (*MountPoint, error)
	Delete(container *MountPoint) error
}

func newMountPointClient(rancherClient *RancherClient) *MountPointClient {
	return &MountPointClient{
		rancherClient: rancherClient,
	}
}

func (c *MountPointClient) Create(container *MountPoint) (*MountPoint, error) {
	resp := &MountPoint{}
	err := c.rancherClient.doCreate(MOUNT_POINT_TYPE, container, resp)
	return resp, err
}

func (c *MountPointClient) Update(existing *MountPoint, updates interface{}) (*MountPoint, error) {
	resp := &MountPoint{}
	err := c.rancherClient.doUpdate(MOUNT_POINT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MountPointClient) List(opts *ListOpts) (*MountPointCollection, error) {
	resp := &MountPointCollection{}
	err := c.rancherClient.doList(MOUNT_POINT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *MountPointCollection) Next() (*MountPointCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &MountPointCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *MountPointClient) ById(id string) (*MountPoint, error) {
	resp := &MountPoint{}
	err := c.rancherClient.doById(MOUNT_POINT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *MountPointClient) Delete(container *MountPoint) error {
	return c.rancherClient.doResourceDelete(MOUNT_POINT_TYPE, &container.Resource)
}
