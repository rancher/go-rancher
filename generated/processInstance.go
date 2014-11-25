package client

const (
	CONTAINER_TYPE = "processInstance"
)

type ProcessInstance struct {
	Resource
    
    Data map[json] `json:"Data,omitempty"`
    
    EndTime date `json:"EndTime,omitempty"`
    
    ExitReason string `json:"ExitReason,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Phase string `json:"Phase,omitempty"`
    
    Priority int `json:"Priority,omitempty"`
    
    ProcessName string `json:"ProcessName,omitempty"`
    
    ResourceId string `json:"ResourceId,omitempty"`
    
    ResourceType string `json:"ResourceType,omitempty"`
    
    Result string `json:"Result,omitempty"`
    
    RunningProcessServerId string `json:"RunningProcessServerId,omitempty"`
    
    StartProcessServerId string `json:"StartProcessServerId,omitempty"`
    
    StartTime date `json:"StartTime,omitempty"`
    
}

type ProcessInstanceCollection struct {
	Collection
	Data []ProcessInstance `json:"data,omitempty"`
}

type ProcessInstanceClient struct {
	rancherClient *RancherClient
}

type ProcessInstanceOperations interface {
	List(opts *ListOpts) (*ProcessInstanceCollection, error)
	Create(opts *ProcessInstance) (*ProcessInstance, error)
	Update(existing *ProcessInstance, updates interface{}) (*ProcessInstance, error)
	ById(id string) (*ProcessInstance, error)
	Delete(container *ProcessInstance) error
}

func newProcessInstanceClient(rancherClient *RancherClient) *ProcessInstanceClient {
	return &ProcessInstanceClient{
		rancherClient: rancherClient,
	}
}

func (self *ProcessInstanceClient) Create(container *ProcessInstance) (*ProcessInstance, error) {
	resp := &ProcessInstance{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ProcessInstanceClient) Update(existing *ProcessInstance, updates interface{}) (*ProcessInstance, error) {
	resp := &ProcessInstance{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ProcessInstanceClient) List(opts *ListOpts) (*ProcessInstanceCollection, error) {
	resp := &ProcessInstanceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ProcessInstanceClient) ById(id string) (*ProcessInstance, error) {
	resp := &ProcessInstance{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ProcessInstanceClient) Delete(container *ProcessInstance) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
