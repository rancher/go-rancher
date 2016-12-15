package client

const (
	WEBHOOK_TYPE = "webhook"
)

type Webhook struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Config interface{} `json:"config,omitempty" yaml:"config,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Driver string `json:"driver,omitempty" yaml:"driver,omitempty"`

	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type WebhookCollection struct {
	Collection
	Data   []Webhook `json:"data,omitempty"`
	client *WebhookClient
}

type WebhookClient struct {
	rancherClient *RancherClient
}

type WebhookOperations interface {
	List(opts *ListOpts) (*WebhookCollection, error)
	Create(opts *Webhook) (*Webhook, error)
	Update(existing *Webhook, updates interface{}) (*Webhook, error)
	ById(id string) (*Webhook, error)
	Delete(container *Webhook) error
}

func newWebhookClient(rancherClient *RancherClient) *WebhookClient {
	return &WebhookClient{
		rancherClient: rancherClient,
	}
}

func (c *WebhookClient) Create(container *Webhook) (*Webhook, error) {
	resp := &Webhook{}
	err := c.rancherClient.doCreate(WEBHOOK_TYPE, container, resp)
	return resp, err
}

func (c *WebhookClient) Update(existing *Webhook, updates interface{}) (*Webhook, error) {
	resp := &Webhook{}
	err := c.rancherClient.doUpdate(WEBHOOK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *WebhookClient) List(opts *ListOpts) (*WebhookCollection, error) {
	resp := &WebhookCollection{}
	err := c.rancherClient.doList(WEBHOOK_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *WebhookCollection) Next() (*WebhookCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &WebhookCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *WebhookClient) ById(id string) (*Webhook, error) {
	resp := &Webhook{}
	err := c.rancherClient.doById(WEBHOOK_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *WebhookClient) Delete(container *Webhook) error {
	return c.rancherClient.doResourceDelete(WEBHOOK_TYPE, &container.Resource)
}
