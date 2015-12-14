package rancher_public_dns

import "github.com/rancher/go-rancher/client"

const (
	API_VERSION_TYPE = "apiVersion"
)

type ApiVersion struct {
	client.Resource
    
    Actions map[string]interface{} `json:"actions,omitempty" yaml:"actions,omitempty"`
    
    Links map[string]interface{} `json:"links,omitempty" yaml:"links,omitempty"`
    
    Type string `json:"type,omitempty" yaml:"type,omitempty"`
    
}

type ApiVersionCollection struct {
	client.Collection
	Data []ApiVersion `json:"data,omitempty"`
}

type ApiVersionClient struct {
	rancherClient *RancherDNSClient
}

type ApiVersionOperations interface {
	List(opts *client.ListOpts) (*ApiVersionCollection, error)
	Create(opts *ApiVersion) (*ApiVersion, error)
	Update(existing *ApiVersion, updates interface{}) (*ApiVersion, error)
	ById(id string) (*ApiVersion, error)
	Delete(container *ApiVersion) error
}

func newApiVersionClient(rancherClient *RancherDNSClient) *ApiVersionClient {
	return &ApiVersionClient{
		rancherClient: rancherClient,
	}
}

func (c *ApiVersionClient) Create(container *ApiVersion) (*ApiVersion, error) {
	resp := &ApiVersion{}
	err := c.rancherClient.Create(API_VERSION_TYPE, container, resp)
	return resp, err
}

func (c *ApiVersionClient) Update(existing *ApiVersion, updates interface{}) (*ApiVersion, error) {
	resp := &ApiVersion{}
	err := c.rancherClient.Update(API_VERSION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ApiVersionClient) List(opts *client.ListOpts) (*ApiVersionCollection, error) {
	resp := &ApiVersionCollection{}
	err := c.rancherClient.List(API_VERSION_TYPE, opts, resp)
	return resp, err
}

func (c *ApiVersionClient) ById(id string) (*ApiVersion, error) {
	resp := &ApiVersion{}
	err := c.rancherClient.ById(API_VERSION_TYPE, id, resp)
	if apiError, ok := err.(*client.ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ApiVersionClient) Delete(container *ApiVersion) error {
	return c.rancherClient.ResourceDelete(API_VERSION_TYPE, &container.Resource)
}
