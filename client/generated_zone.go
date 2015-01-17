package client

const (
	ZONE_TYPE = "zone"
)

type Zone struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type ZoneCollection struct {
	Collection
	Data []Zone `json:"data,omitempty"`
}

type ZoneClient struct {
	rancherClient *RancherClient
}

type ZoneOperations interface {
	List(opts *ListOpts) (*ZoneCollection, error)
	Create(opts *Zone) (*Zone, error)
	Update(existing *Zone, updates interface{}) (*Zone, error)
	ById(id string) (*Zone, error)
	Delete(container *Zone) error
}

func newZoneClient(rancherClient *RancherClient) *ZoneClient {
	return &ZoneClient{
		rancherClient: rancherClient,
	}
}

func (c *ZoneClient) Create(container *Zone) (*Zone, error) {
	resp := &Zone{}
	err := c.rancherClient.doCreate(ZONE_TYPE, container, resp)
	return resp, err
}

func (c *ZoneClient) Update(existing *Zone, updates interface{}) (*Zone, error) {
	resp := &Zone{}
	err := c.rancherClient.doUpdate(ZONE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ZoneClient) List(opts *ListOpts) (*ZoneCollection, error) {
	resp := &ZoneCollection{}
	err := c.rancherClient.doList(ZONE_TYPE, opts, resp)
	return resp, err
}

func (c *ZoneClient) ById(id string) (*Zone, error) {
	resp := &Zone{}
	err := c.rancherClient.doById(ZONE_TYPE, id, resp)
	return resp, err
}

func (c *ZoneClient) Delete(container *Zone) error {
	return c.rancherClient.doResourceDelete(ZONE_TYPE, &container.Resource)
}
