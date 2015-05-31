package client

const (
	SERVICE_EXPOSE_MAP_TYPE = "serviceExposeMap"
)

type ServiceExposeMap struct {
	Resource

	AccountId string `json:"accountId,omitempty"`

	Created string `json:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty"`

	Description string `json:"description,omitempty"`

	InstanceId string `json:"instanceId,omitempty"`

	IpAddress string `json:"ipAddress,omitempty"`

	Kind string `json:"kind,omitempty"`

	Name string `json:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty"`

	Removed string `json:"removed,omitempty"`

	ServiceId string `json:"serviceId,omitempty"`

	State string `json:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty"`

	TransitioningProgress int `json:"transitioningProgress,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}

type ServiceExposeMapCollection struct {
	Collection
	Data []ServiceExposeMap `json:"data,omitempty"`
}

type ServiceExposeMapClient struct {
	rancherClient *RancherClient
}

type ServiceExposeMapOperations interface {
	List(opts *ListOpts) (*ServiceExposeMapCollection, error)
	Create(opts *ServiceExposeMap) (*ServiceExposeMap, error)
	Update(existing *ServiceExposeMap, updates interface{}) (*ServiceExposeMap, error)
	ById(id string) (*ServiceExposeMap, error)
	Delete(container *ServiceExposeMap) error

	ActionCreate(*ServiceExposeMap) (*ServiceExposeMap, error)

	ActionRemove(*ServiceExposeMap) (*ServiceExposeMap, error)
}

func newServiceExposeMapClient(rancherClient *RancherClient) *ServiceExposeMapClient {
	return &ServiceExposeMapClient{
		rancherClient: rancherClient,
	}
}

func (c *ServiceExposeMapClient) Create(container *ServiceExposeMap) (*ServiceExposeMap, error) {
	resp := &ServiceExposeMap{}
	err := c.rancherClient.doCreate(SERVICE_EXPOSE_MAP_TYPE, container, resp)
	return resp, err
}

func (c *ServiceExposeMapClient) Update(existing *ServiceExposeMap, updates interface{}) (*ServiceExposeMap, error) {
	resp := &ServiceExposeMap{}
	err := c.rancherClient.doUpdate(SERVICE_EXPOSE_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceExposeMapClient) List(opts *ListOpts) (*ServiceExposeMapCollection, error) {
	resp := &ServiceExposeMapCollection{}
	err := c.rancherClient.doList(SERVICE_EXPOSE_MAP_TYPE, opts, resp)
	return resp, err
}

func (c *ServiceExposeMapClient) ById(id string) (*ServiceExposeMap, error) {
	resp := &ServiceExposeMap{}
	err := c.rancherClient.doById(SERVICE_EXPOSE_MAP_TYPE, id, resp)
	return resp, err
}

func (c *ServiceExposeMapClient) Delete(container *ServiceExposeMap) error {
	return c.rancherClient.doResourceDelete(SERVICE_EXPOSE_MAP_TYPE, &container.Resource)
}

func (c *ServiceExposeMapClient) ActionCreate(resource *ServiceExposeMap) (*ServiceExposeMap, error) {

	resp := &ServiceExposeMap{}

	err := c.rancherClient.doAction(SERVICE_EXPOSE_MAP_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ServiceExposeMapClient) ActionRemove(resource *ServiceExposeMap) (*ServiceExposeMap, error) {

	resp := &ServiceExposeMap{}

	err := c.rancherClient.doAction(SERVICE_EXPOSE_MAP_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
