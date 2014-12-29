package client

const (
	CONTAINER_TYPE = "task"
)

type Task struct {
	Resource
    
    Id int `json:"Id,omitempty"`
    
    Name string `json:"Name,omitempty"`
    
}

type TaskCollection struct {
	Collection
	Data []Task `json:"data,omitempty"`
}

type TaskClient struct {
	rancherClient *RancherClient
}

type TaskOperations interface {
	List(opts *ListOpts) (*TaskCollection, error)
	Create(opts *Task) (*Task, error)
	Update(existing *Task, updates interface{}) (*Task, error)
	ById(id string) (*Task, error)
	Delete(container *Task) error
}

func newTaskClient(rancherClient *RancherClient) *TaskClient {
	return &TaskClient{
		rancherClient: rancherClient,
	}
}

func (self *TaskClient) Create(container *Task) (*Task, error) {
	resp := &Task{}
	err := self.rancherClient.doCreate(CONTAINER_TYPE, container, resp)
	return resp, err
}

func (self *TaskClient) Update(existing *Task, updates interface{}) (*Task, error) {
	resp := &Task{}
	err := self.rancherClient.doUpdate(CONTAINER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (self *TaskClient) List(opts *ListOpts) (*TaskCollection, error) {
	resp := &TaskCollection{}
	err := self.rancherClient.doList(CONTAINER_TYPE, opts, resp)
	return resp, err
}

func (self *TaskClient) ById(id string) (*Task, error) {
	resp := &Task{}
	err := self.rancherClient.doById(CONTAINER_TYPE, id, resp)
	return resp, err
}

func (self *TaskClient) Delete(container *Task) error {
	return self.rancherClient.doResourceDelete(CONTAINER_TYPE, &container.Resource)
}
