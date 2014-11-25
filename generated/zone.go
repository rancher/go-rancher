package client

const (
	CONTAINER_TYPE = "zone"
)

type Zone struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
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

func (self *ZoneClient) Create(container *Zone) (*Zone, error) {
	resp := &Zone{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ZoneClient) Update(existing *Zone, updates interface{}) (*Zone, error) {
	resp := &Zone{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ZoneClient) List(opts *ListOpts) (*ZoneCollection, error) {
	resp := &ZoneCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ZoneClient) ById(id string) (*Zone, error) {
	resp := &Zone{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ZoneClient) Delete(container *Zone) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
