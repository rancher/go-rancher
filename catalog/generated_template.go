package catalog

const (
	TEMPLATE_TYPE = "template"
)

type Template struct {
	Resource

	Actions map[string]interface{} `json:"actions,omitempty" yaml:"actions,omitempty"`

	Bindings map[string]interface{} `json:"bindings,omitempty" yaml:"bindings,omitempty"`

	CatalogId string `json:"catalogId,omitempty" yaml:"catalog_id,omitempty"`

	Category string `json:"category,omitempty" yaml:"category,omitempty"`

	DefaultVersion string `json:"defaultVersion,omitempty" yaml:"default_version,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Files map[string]interface{} `json:"files,omitempty" yaml:"files,omitempty"`

	IsSystem string `json:"isSystem,omitempty" yaml:"is_system,omitempty"`

	Labels map[string]interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`

	License string `json:"license,omitempty" yaml:"license,omitempty"`

	Links map[string]interface{} `json:"links,omitempty" yaml:"links,omitempty"`

	Maintainer string `json:"maintainer,omitempty" yaml:"maintainer,omitempty"`

	MaximumRancherVersion string `json:"maximumRancherVersion,omitempty" yaml:"maximum_rancher_version,omitempty"`

	MinimumRancherVersion string `json:"minimumRancherVersion,omitempty" yaml:"minimum_rancher_version,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Output Model.Output `json:"output,omitempty" yaml:"output,omitempty"`

	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	TemplateBase string `json:"templateBase,omitempty" yaml:"template_base,omitempty"`

	TemplateVersionRancherVersion map[string]interface{} `json:"templateVersionRancherVersion,omitempty" yaml:"template_version_rancher_version,omitempty"`

	TemplateVersionRancherVersionGte map[string]interface{} `json:"templateVersionRancherVersionGte,omitempty" yaml:"template_version_rancher_version_gte,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	UpgradeFrom string `json:"upgradeFrom,omitempty" yaml:"upgrade_from,omitempty"`

	UpgradeVersionLinks map[string]interface{} `json:"upgradeVersionLinks,omitempty" yaml:"upgrade_version_links,omitempty"`

	VersionLinks map[string]interface{} `json:"versionLinks,omitempty" yaml:"version_links,omitempty"`
}

type TemplateCollection struct {
	Collection
	Data   []Template `json:"data,omitempty"`
	client *TemplateClient
}

type TemplateClient struct {
	rancherClient *RancherClient
}

type TemplateOperations interface {
	List(opts *ListOpts) (*TemplateCollection, error)
	Create(opts *Template) (*Template, error)
	Update(existing *Template, updates interface{}) (*Template, error)
	ById(id string) (*Template, error)
	Delete(container *Template) error
}

func newTemplateClient(rancherClient *RancherClient) *TemplateClient {
	return &TemplateClient{
		rancherClient: rancherClient,
	}
}

func (c *TemplateClient) Create(container *Template) (*Template, error) {
	resp := &Template{}
	err := c.rancherClient.doCreate(TEMPLATE_TYPE, container, resp)
	return resp, err
}

func (c *TemplateClient) Update(existing *Template, updates interface{}) (*Template, error) {
	resp := &Template{}
	err := c.rancherClient.doUpdate(TEMPLATE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TemplateClient) List(opts *ListOpts) (*TemplateCollection, error) {
	resp := &TemplateCollection{}
	err := c.rancherClient.doList(TEMPLATE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *TemplateCollection) Next() (*TemplateCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TemplateCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TemplateClient) ById(id string) (*Template, error) {
	resp := &Template{}
	err := c.rancherClient.doById(TEMPLATE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *TemplateClient) Delete(container *Template) error {
	return c.rancherClient.doResourceDelete(TEMPLATE_TYPE, &container.Resource)
}
