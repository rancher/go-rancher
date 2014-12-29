package client

const (
	CONTAINER_TYPE = "processExecution"
)

type ProcessExecution struct {
	Resource
    
    Id int `json:"Id,omitempty"`
    
    Log map[json] `json:"Log,omitempty"`
    
    ProcessInstanceId string `json:"ProcessInstanceId,omitempty"`
    
    Uuid string `json:"Uuid,omitempty"`
    
}

type ProcessExecutionCollection struct {
	Collection
	Data []ProcessExecution `json:"data,omitempty"`
}

type ProcessExecutionClient struct {
	rancherClient *RancherClient
}

type ProcessExecutionOperations interface {
	List(opts *ListOpts) (*ProcessExecutionCollection, error)
	Create(opts *ProcessExecution) (*ProcessExecution, error)
	Update(existing *ProcessExecution, updates interface{}) (*ProcessExecution, error)
	ById(id string) (*ProcessExecution, error)
	Delete(container *ProcessExecution) error
}

func newProcessExecutionClient(rancherClient *RancherClient) *ProcessExecutionClient {
	return &ProcessExecutionClient{
		rancherClient: rancherClient,
	}
}

func (self *ProcessExecutionClient) Create(container *ProcessExecution) (*ProcessExecution, error) {
	resp := &ProcessExecution{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *ProcessExecutionClient) Update(existing *ProcessExecution, updates interface{}) (*ProcessExecution, error) {
	resp := &ProcessExecution{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *ProcessExecutionClient) List(opts *ListOpts) (*ProcessExecutionCollection, error) {
	resp := &ProcessExecutionCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *ProcessExecutionClient) ById(id string) (*ProcessExecution, error) {
	resp := &ProcessExecution{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *ProcessExecutionClient) Delete(container *ProcessExecution) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
