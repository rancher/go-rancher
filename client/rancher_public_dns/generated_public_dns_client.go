package rancher_public_dns

import "github.com/rancher/go-rancher/client"

type RancherDNSClient struct {
    client.RancherClient
    client.RancherBaseClient
	
    ApiVersion ApiVersionOperations
    RootDomainInfo RootDomainInfoOperations
    DnsRecord DnsRecordOperations
}

func constructDNSClient() *RancherDNSClient {
	dnsClient := &RancherDNSClient{
		RancherBaseClient: client.RancherBaseClient{
			Types: map[string]client.Schema{},
		},
	}

    
    dnsClient.ApiVersion = newApiVersionClient(dnsClient)
    dnsClient.RootDomainInfo = newRootDomainInfoClient(dnsClient)
    dnsClient.DnsRecord = newDnsRecordClient(dnsClient) 


	return dnsClient
}

func NewRancherDNSClient(opts *client.ClientOpts) (*RancherDNSClient, error) {
    dnsClient := constructDNSClient()
        
    err := client.SetupRancherBaseClient(&dnsClient.RancherBaseClient, opts)
    if err != nil {
        return nil, err
    }

    return dnsClient, nil
}
