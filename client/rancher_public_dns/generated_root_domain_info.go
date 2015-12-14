package rancher_public_dns

import (
		"fmt"
		"github.com/rancher/go-rancher/client"
)

const (
	ROOT_DOMAIN_INFO_TYPE = "rootDomainInfo"
)

type RootDomainInfo struct {
	client.Resource

	Actions map[string]interface{} `json:"actions,omitempty" yaml:"actions,omitempty"`

	Links map[string]interface{} `json:"links,omitempty" yaml:"links,omitempty"`

	RootDomain string `json:"rootDomain,omitempty" yaml:"root_domain,omitempty"`

	Token string `json:"token,omitempty" yaml:"token,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

type RootDomainInfoCollection struct {
	client.Collection
	Data []RootDomainInfo `json:"data,omitempty"`
}

type RootDomainInfoClient struct {
	rancherClient *RancherDNSClient
}

type RootDomainInfoOperations interface {
	Get(opts *client.ListOpts) (*RootDomainInfo, error)
	Create(opts *RootDomainInfo) (*RootDomainInfo, error)
	Update(existing *RootDomainInfo, updates interface{}) (*RootDomainInfo, error)
	ById(id string) (*RootDomainInfo, error)
	Delete(container *RootDomainInfo) error
}

func newRootDomainInfoClient(rancherClient *RancherDNSClient) *RootDomainInfoClient {
	return &RootDomainInfoClient{
		rancherClient: rancherClient,
	}
}

func (c *RootDomainInfoClient) Create(container *RootDomainInfo) (*RootDomainInfo, error) {
	resp := &RootDomainInfo{}
	err := c.rancherClient.Create(ROOT_DOMAIN_INFO_TYPE, container, resp)
	return resp, err
}

func (c *RootDomainInfoClient) Update(existing *RootDomainInfo, updates interface{}) (*RootDomainInfo, error) {
	resp := &RootDomainInfo{}
	err := c.rancherClient.Update(ROOT_DOMAIN_INFO_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RootDomainInfoClient) Get(opts *client.ListOpts) (*RootDomainInfo, error) {
	resp := &RootDomainInfo{}
	fmt.Printf("\nList called rootDI: %s", ROOT_DOMAIN_INFO_TYPE)
	err := c.rancherClient.List(ROOT_DOMAIN_INFO_TYPE, opts, resp)
	return resp, err
}

func (c *RootDomainInfoClient) ById(id string) (*RootDomainInfo, error) {
	resp := &RootDomainInfo{}
	err := c.rancherClient.ById(ROOT_DOMAIN_INFO_TYPE, id, resp)
	if apiError, ok := err.(*client.ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *RootDomainInfoClient) Delete(container *RootDomainInfo) error {
	return c.rancherClient.ResourceDelete(ROOT_DOMAIN_INFO_TYPE, &container.Resource)
}
