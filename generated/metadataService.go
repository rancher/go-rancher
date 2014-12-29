package client

const (
	CONTAINER_TYPE = "metadataService"
)

type MetadataService struct {
	Resource
    
    AccountId string `json:"AccountId,omitempty"`
    
    ConfigDrive boolean `json:"ConfigDrive,omitempty"`
    
    Created date `json:"Created,omitempty"`
    
    Data map[json] `json:"Data,omitempty"`
    
    Description string `json:"Description,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Kind string `json:"Kind,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    NetworkId string `json:"NetworkId,omitempty"`
    
    NetworkServiceProviderId string `json:"NetworkServiceProviderId,omitempty"`
    
    RemoveTime date `json:"RemoveTime,omitempty"`
    
    Removed date `json:"Removed,omitempty"`
    
    State enum `json:"State,omitempty"`
    
    Transitioning enum `json:"Transitioning,omitempty"`
    
    TransitioningMessage string `json:"TransitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"TransitioningProgress,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
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

func (self *MetadataServiceClient) Create(container *MetadataService) (*MetadataService, error) {
	resp := &MetadataService{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *MetadataServiceClient) Update(existing *MetadataService, updates interface{}) (*MetadataService, error) {
	resp := &MetadataService{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *MetadataServiceClient) List(opts *ListOpts) (*MetadataServiceCollection, error) {
	resp := &MetadataServiceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *MetadataServiceClient) ById(id string) (*MetadataService, error) {
	resp := &MetadataService{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *MetadataServiceClient) Delete(container *MetadataService) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
