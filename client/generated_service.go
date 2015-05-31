package client

const (
	SERVICE_TYPE = "service"
)

type Service struct {
	Resource

	AccountId string `json:"accountId,omitempty"`

	Created string `json:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty"`

	DataVolumesFromService []string `json:"dataVolumesFromService,omitempty"`

	Description string `json:"description,omitempty"`

	EnvironmentId string `json:"environmentId,omitempty"`

	Kind string `json:"kind,omitempty"`

	LaunchConfig Container `json:"launchConfig,omitempty"`

	Name string `json:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty"`

	Removed string `json:"removed,omitempty"`

	Scale int `json:"scale,omitempty"`

	State string `json:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty"`

	TransitioningProgress int `json:"transitioningProgress,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}

type ServiceCollection struct {
	Collection
	Data []Service `json:"data,omitempty"`
}

type ServiceClient struct {
	rancherClient *RancherClient
}

type ServiceOperations interface {
	List(opts *ListOpts) (*ServiceCollection, error)
	Create(opts *Service) (*Service, error)
	Update(existing *Service, updates interface{}) (*Service, error)
	ById(id string) (*Service, error)
	Delete(container *Service) error

	ActionActivate(*Service) (*Service, error)

	ActionAddservicelink(*Service, *AddRemoveServiceLinkInput) (*Service, error)

	ActionCreate(*Service) (*Service, error)

	ActionDeactivate(*Service) (*Service, error)

	ActionRemove(*Service) (*Service, error)

	ActionRemoveservicelink(*Service, *AddRemoveServiceLinkInput) (*Service, error)

	ActionSetservicelinks(*Service, *SetServiceLinksInput) (*Service, error)

	ActionUpdate(*Service) (*Service, error)
}

func newServiceClient(rancherClient *RancherClient) *ServiceClient {
	return &ServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *ServiceClient) Create(container *Service) (*Service, error) {
	resp := &Service{}
	err := c.rancherClient.doCreate(SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *ServiceClient) Update(existing *Service, updates interface{}) (*Service, error) {
	resp := &Service{}
	err := c.rancherClient.doUpdate(SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceClient) List(opts *ListOpts) (*ServiceCollection, error) {
	resp := &ServiceCollection{}
	err := c.rancherClient.doList(SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *ServiceClient) ById(id string) (*Service, error) {
	resp := &Service{}
	err := c.rancherClient.doById(SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *ServiceClient) Delete(container *Service) error {
	return c.rancherClient.doResourceDelete(SERVICE_TYPE, &container.Resource)
}

func (c *ServiceClient) ActionActivate(resource *Service) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(SERVICE_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ServiceClient) ActionAddservicelink(resource *Service, input *AddRemoveServiceLinkInput) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(SERVICE_TYPE, "addservicelink", &resource.Resource, input, resp)

	return resp, err
}

func (c *ServiceClient) ActionCreate(resource *Service) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(SERVICE_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ServiceClient) ActionDeactivate(resource *Service) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(SERVICE_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ServiceClient) ActionRemove(resource *Service) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(SERVICE_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ServiceClient) ActionRemoveservicelink(resource *Service, input *AddRemoveServiceLinkInput) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(SERVICE_TYPE, "removeservicelink", &resource.Resource, input, resp)

	return resp, err
}

func (c *ServiceClient) ActionSetservicelinks(resource *Service, input *SetServiceLinksInput) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(SERVICE_TYPE, "setservicelinks", &resource.Resource, input, resp)

	return resp, err
}

func (c *ServiceClient) ActionUpdate(resource *Service) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(SERVICE_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
