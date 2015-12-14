package rancher_public_dns

import (
	"testing"
	"time"

	b64 "encoding/base64"
	"github.com/rancher/go-rancher/client"
)

const (
	PROJECT_URL = "http://localhost:8089/v1-rancher-dns/schemas"
	URL         = "http://localhost:8089/v1-rancher-dns"
	//ACCESS_KEY  = "admin"
	//SECRET_KEY  = "adminpass"
	MAX_WAIT = time.Duration(time.Second * 10)
)

var (
	TOKEN                string
	rootDomainInfoCached *RootDomainInfo
)

func newClient(t *testing.T, url string, authHeader string) *RancherDNSClient {
	dnsClient, err := NewRancherDNSClient(&client.ClientOpts{
		Url:             url,
		CustomAuthToken: authHeader,
		//SecretKey: SECRET_KEY,
	})

	if err != nil {
		t.Fatal("Failed to create client", err)
	}

	return dnsClient
}

func TestClientLoad(t *testing.T) {
	dnsClient := newClient(t, URL, "")
	if dnsClient.Schemas == nil {
		t.Fatal("Failed to load schema")
	}

	if len(dnsClient.Schemas.Data) == 0 {
		t.Fatal("Schemas is empty")
	}

	if _, ok := dnsClient.Types["rootDomainInfo"]; !ok {
		t.Fatal("Failed to find rootDomainInfo type")
	}
}

func TestGetRootDomainInfo(t *testing.T) {
	dnsClient := newClient(t, PROJECT_URL, "")

	rootDomainInfo, err := dnsClient.RootDomainInfo.Get(nil)
	if err != nil {
		t.Fatal("Failed to list ", err)
	}

	dnsClientNext := newClient(t, PROJECT_URL, b64.StdEncoding.EncodeToString([]byte(rootDomainInfo.Token)))

	rootDomainInfoAgain, err := dnsClientNext.RootDomainInfo.Get(nil)
	if err != nil {
		t.Fatal("Failed to list again", err)
	}

	if rootDomainInfo.Id != rootDomainInfoAgain.Id {
		t.Fatal("Got different ids:  %s, %s ", rootDomainInfo.Id, rootDomainInfoAgain.Id)
	}

	TOKEN = rootDomainInfo.Token
	rootDomainInfoCached = rootDomainInfo

}

func TestAddDNSRecords(t *testing.T) {
	if TOKEN != "" {
		dnsClient := newClient(t, PROJECT_URL, b64.StdEncoding.EncodeToString([]byte(TOKEN)))

		dnsRec, err := dnsClient.DnsRecord.Create(&DnsRecord{
			Fqdn:       "lb2.default.default." + rootDomainInfoCached.RootDomain,
			Records:    []string{"127.0.0.1"},
			Recordtype: "A",
			Ttl:        1900,
		})

		dnsRec.Resource.Links = make(map[string]string)
		for key, value := range dnsRec.Links {

			if str, ok := value.(string); ok {
				dnsRec.Resource.Links[key] = str
			}
		}

		defer dnsClient.DnsRecord.Delete(dnsRec)

		if err != nil {
			t.Fatal("Failed to create DnsRecord ", err)
		}

		if dnsRec.Fqdn != "lb2.default.default."+rootDomainInfoCached.RootDomain {
			t.Fatal("FQDN name is wrong [" + dnsRec.Fqdn + "]")
		}

	}
}

func TestGetDNSRecords(t *testing.T) {
	if TOKEN != "" {
		dnsClient := newClient(t, PROJECT_URL, b64.StdEncoding.EncodeToString([]byte(TOKEN)))

		dnsRec, err := dnsClient.DnsRecord.Create(&DnsRecord{
			Fqdn:       "lb2.default.default." + rootDomainInfoCached.RootDomain,
			Records:    []string{"127.0.0.1"},
			Recordtype: "A",
			Ttl:        1900,
		})

		dnsRec.Resource.Links = make(map[string]string)
		for key, value := range dnsRec.Links {

			if str, ok := value.(string); ok {
				dnsRec.Resource.Links[key] = str
			}
		}

		defer dnsClient.DnsRecord.Delete(dnsRec)

		dnsrecords, err := dnsClient.DnsRecord.List(nil)
		if err != nil {
			t.Fatal("Failed to list ", err)
		}

		if len(dnsrecords.Data) == 0 {
			t.Fatal("No dnsRecords found")
		}

		//getById
		lookedupRec, err := dnsClient.DnsRecord.ById(dnsRec.Id)
		if err != nil {
			t.Fatal("Failed to GetRecord ById ", err)
		}

		if lookedupRec.Fqdn != dnsRec.Fqdn {
			t.Fatal("GetRecord ById failed")
		}

	}
}

func TestUpdateDNSRecords(t *testing.T) {
	if TOKEN != "" {
		dnsClient := newClient(t, PROJECT_URL, b64.StdEncoding.EncodeToString([]byte(TOKEN)))

		dnsRec, err := dnsClient.DnsRecord.Create(&DnsRecord{
			Fqdn:       "lb2.default.default." + rootDomainInfoCached.RootDomain,
			Records:    []string{"127.0.0.1"},
			Recordtype: "A",
			Ttl:        1900,
		})
		if err != nil {
			t.Fatal("Failed to create DnsRecord ", err)
		}

		if dnsRec.Fqdn != "lb2.default.default."+rootDomainInfoCached.RootDomain {
			t.Fatal("FQDN name is wrong [" + dnsRec.Fqdn + "]")
		}

		dnsRec.Resource.Links = make(map[string]string)
		for key, value := range dnsRec.Links {

			if str, ok := value.(string); ok {
				dnsRec.Resource.Links[key] = str
			}
		}

		updatedDnsRec, err := dnsClient.DnsRecord.Update(dnsRec, &DnsRecord{
			Fqdn:       "lb2.default.default." + rootDomainInfoCached.RootDomain,
			Records:    []string{"127.0.0.1"},
			Recordtype: "A",
			Ttl:        2100,
		})

		updatedDnsRec.Resource.Links = make(map[string]string)
		for key, value := range updatedDnsRec.Links {

			if str, ok := value.(string); ok {
				updatedDnsRec.Resource.Links[key] = str
			}
		}

		defer dnsClient.DnsRecord.Delete(updatedDnsRec)

		if err != nil {
			t.Fatal("Failed to update DnsRecord ", err)
		}

		if updatedDnsRec.Ttl != 2100 {
			t.Fatal("Ttl  is wrong ")
		}

	}
}

func TestDeleteDNSRecords(t *testing.T) {
	if TOKEN != "" {
		dnsClient := newClient(t, PROJECT_URL, b64.StdEncoding.EncodeToString([]byte(TOKEN)))

		dnsRec, err := dnsClient.DnsRecord.Create(&DnsRecord{
			Fqdn:       "lb3.default.default." + rootDomainInfoCached.RootDomain,
			Records:    []string{"127.0.0.1"},
			Recordtype: "A",
			Ttl:        1900,
		})
		if err != nil {
			t.Fatal("Failed to create DnsRecord ", err)
		}

		if dnsRec.Fqdn != "lb3.default.default."+rootDomainInfoCached.RootDomain {
			t.Fatal("FQDN name is wrong [" + dnsRec.Fqdn + "]")
		}

		dnsRec.Resource.Links = make(map[string]string)
		for key, value := range dnsRec.Links {

			if str, ok := value.(string); ok {
				dnsRec.Resource.Links[key] = str
			}
		}
		err = dnsClient.DnsRecord.Delete(dnsRec)
		if err != nil {
			t.Fatal("Failed to delete DnsRecord ", err)
		}

	}
}
