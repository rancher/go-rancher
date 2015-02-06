package client

const (
	MACHINE_HOST_TYPE = "machineHost"
)

type MachineHost struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    AuthCertificateAuthority string `json:"authCertificateAuthority,omitempty"`
    
    AuthKey string `json:"authKey,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    DigitaloceanConfig DigitaloceanConfig `json:"digitaloceanConfig,omitempty"`
    
    Driver string `json:"driver,omitempty"`
    
    ExternalId string `json:"externalId,omitempty"`
    
    ExtractedConfig string `json:"extractedConfig,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    VirtualboxConfig VirtualboxConfig `json:"virtualboxConfig,omitempty"`
    
}

type MachineHostCollection struct {
	Collection
	Data []MachineHost `json:"data,omitempty"`
}

type MachineHostClient struct {
	rancherClient *RancherClient
}

type MachineHostOperations interface {
	List(opts *ListOpts) (*MachineHostCollection, error)
	Create(opts *MachineHost) (*MachineHost, error)
	Update(existing *MachineHost, updates interface{}) (*MachineHost, error)
	ById(id string) (*MachineHost, error)
	Delete(container *MachineHost) error
}

func newMachineHostClient(rancherClient *RancherClient) *MachineHostClient {
	return &MachineHostClient{
		rancherClient: rancherClient,
	}
}

func (c *MachineHostClient) Create(container *MachineHost) (*MachineHost, error) {
	resp := &MachineHost{}
	err := c.rancherClient.doCreate(MACHINE_HOST_TYPE, container, resp)
	return resp, err
}

func (c *MachineHostClient) Update(existing *MachineHost, updates interface{}) (*MachineHost, error) {
	resp := &MachineHost{}
	err := c.rancherClient.doUpdate(MACHINE_HOST_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MachineHostClient) List(opts *ListOpts) (*MachineHostCollection, error) {
	resp := &MachineHostCollection{}
	err := c.rancherClient.doList(MACHINE_HOST_TYPE, opts, resp)
	return resp, err
}

func (c *MachineHostClient) ById(id string) (*MachineHost, error) {
	resp := &MachineHost{}
	err := c.rancherClient.doById(MACHINE_HOST_TYPE, id, resp)
	return resp, err
}

func (c *MachineHostClient) Delete(container *MachineHost) error {
	return c.rancherClient.doResourceDelete(MACHINE_HOST_TYPE, &container.Resource)
}
