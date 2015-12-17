package rancher_public_dns

import "github.com/rancher/go-rancher/client"

const (
	DNS_RECORD_TYPE = "dnsRecord"
)

type DnsRecord struct {
	client.Resource

	Actions map[string]interface{} `json:"actions,omitempty" yaml:"actions,omitempty"`

	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`

	Links map[string]interface{} `json:"links,omitempty" yaml:"links,omitempty"`

	Records []string `json:"records,omitempty" yaml:"records,omitempty"`

	Recordtype string `json:"recordtype,omitempty" yaml:"recordtype,omitempty"`

	Ttl int64 `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type DnsRecordCollection struct {
	client.Collection
	Data []DnsRecord `json:"data,omitempty"`
}

type DnsRecordClient struct {
	rancherClient *RancherDNSClient
}

type DnsRecordOperations interface {
	List(opts *client.ListOpts) (*DnsRecordCollection, error)
	Create(opts *DnsRecord) (*DnsRecord, error)
	Update(existing *DnsRecord, updates interface{}) (*DnsRecord, error)
	ById(id string) (*DnsRecord, error)
	Delete(container *DnsRecord) error
}

func newDnsRecordClient(rancherClient *RancherDNSClient) *DnsRecordClient {
	return &DnsRecordClient{
		rancherClient: rancherClient,
	}
}

func (c *DnsRecordClient) Create(container *DnsRecord) (*DnsRecord, error) {
	resp := &DnsRecord{}
	err := c.rancherClient.Create(DNS_RECORD_TYPE, container, resp)
	return resp, err
}

func (c *DnsRecordClient) Update(existing *DnsRecord, updates interface{}) (*DnsRecord, error) {
	_, ok := existing.Resource.Links[client.SELF]

	if !ok {
		existing.Resource.Links = make(map[string]string)
		for key, value := range existing.Links {
			if str, ok := value.(string); ok {
				existing.Resource.Links[key] = str
			}
		}
	}
	resp := &DnsRecord{}
	err := c.rancherClient.Update(DNS_RECORD_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DnsRecordClient) List(opts *client.ListOpts) (*DnsRecordCollection, error) {
	resp := &DnsRecordCollection{}
	err := c.rancherClient.List(DNS_RECORD_TYPE, opts, resp)
	return resp, err
}

func (c *DnsRecordClient) ById(id string) (*DnsRecord, error) {
	resp := &DnsRecord{}
	err := c.rancherClient.ById(DNS_RECORD_TYPE, id, resp)
	if apiError, ok := err.(*client.ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *DnsRecordClient) Delete(dnsRec *DnsRecord) error {

	_, ok := dnsRec.Resource.Links[client.SELF]
	if !ok {

		dnsRec.Resource.Links = make(map[string]string)
		for key, value := range dnsRec.Links {
			if str, ok := value.(string); ok {
				dnsRec.Resource.Links[key] = str
			}
		}
	}

	return c.rancherClient.ResourceDelete(DNS_RECORD_TYPE, &dnsRec.Resource)
}
