package client

const (
	METADATA_SERVICE_TYPE = "metadataService"
)

type MetadataService struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    ConfigDrive bool `json:"configDrive,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    NetworkId string `json:"networkId,omitempty"`
    
    NetworkServiceProviderId string `json:"networkServiceProviderId,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type MetadataServiceCollection struct {
	Collection
	Data []MetadataService `json:"data,omitempty"`
}

type MetadataServiceClient struct {
	rancherClient *RancherClient
}

type MetadataServiceOperations interface {
	List(opts *ListOpts) (*MetadataServiceCollection, error)
	Create(opts *MetadataService) (*MetadataService, error)
	Update(existing *MetadataService, updates interface{}) (*MetadataService, error)
	ById(id string) (*MetadataService, error)
	Delete(container *MetadataService) error
}

func newMetadataServiceClient(rancherClient *RancherClient) *MetadataServiceClient {
	return &MetadataServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *MetadataServiceClient) Create(container *MetadataService) (*MetadataService, error) {
	resp := &MetadataService{}
	err := c.rancherClient.doCreate(METADATA_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *MetadataServiceClient) Update(existing *MetadataService, updates interface{}) (*MetadataService, error) {
	resp := &MetadataService{}
	err := c.rancherClient.doUpdate(METADATA_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MetadataServiceClient) List(opts *ListOpts) (*MetadataServiceCollection, error) {
	resp := &MetadataServiceCollection{}
	err := c.rancherClient.doList(METADATA_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *MetadataServiceClient) ById(id string) (*MetadataService, error) {
	resp := &MetadataService{}
	err := c.rancherClient.doById(METADATA_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *MetadataServiceClient) Delete(container *MetadataService) error {
	return c.rancherClient.doResourceDelete(METADATA_SERVICE_TYPE, &container.Resource)
}
