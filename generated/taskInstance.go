package client

const (
	CONTAINER_TYPE = "taskInstance"
)

type TaskInstance struct {
	Resource
    
    EndTime date `json:"EndTime,omitempty"`
    
    Exception string `json:"Exception,omitempty"`
    
    Id int `json:"Id,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
    ServerId string `json:"ServerId,omitempty"`
    
    StartTime date `json:"StartTime,omitempty"`
    
    TaskId string `json:"TaskId,omitempty"`
    
}

type TaskInstanceCollection struct {
	Collection
	Data []TaskInstance `json:"data,omitempty"`
}

type TaskInstanceClient struct {
	rancherClient *RancherClient
}

type TaskInstanceOperations interface {
	List(opts *ListOpts) (*TaskInstanceCollection, error)
	Create(opts *TaskInstance) (*TaskInstance, error)
	Update(existing *TaskInstance, updates interface{}) (*TaskInstance, error)
	ById(id string) (*TaskInstance, error)
	Delete(container *TaskInstance) error
}

func newTaskInstanceClient(rancherClient *RancherClient) *TaskInstanceClient {
	return &TaskInstanceClient{
		rancherClient: rancherClient,
	}
}

func (self *TaskInstanceClient) Create(container *TaskInstance) (*TaskInstance, error) {
	resp := &TaskInstance{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *TaskInstanceClient) Update(existing *TaskInstance, updates interface{}) (*TaskInstance, error) {
	resp := &TaskInstance{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *TaskInstanceClient) List(opts *ListOpts) (*TaskInstanceCollection, error) {
	resp := &TaskInstanceCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *TaskInstanceClient) ById(id string) (*TaskInstance, error) {
	resp := &TaskInstance{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *TaskInstanceClient) Delete(container *TaskInstance) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
