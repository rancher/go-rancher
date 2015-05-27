package client

const (
	AMAZONEC2CONFIG_TYPE = "amazonec2Config"
)

type Amazonec2Config struct {
	Resource
    
    AccessKey string `json:"accessKey,omitempty"`
    
    Ami string `json:"ami,omitempty"`
    
    IamInstanceProfile string `json:"iamInstanceProfile,omitempty"`
    
    InstanceType string `json:"instanceType,omitempty"`
    
    Monitoring bool `json:"monitoring,omitempty"`
    
    PrivateAddressOnly bool `json:"privateAddressOnly,omitempty"`
    
    Region string `json:"region,omitempty"`
    
    RequestSpotInstance bool `json:"requestSpotInstance,omitempty"`
    
    RootSize string `json:"rootSize,omitempty"`
    
    SecretKey string `json:"secretKey,omitempty"`
    
    SecurityGroup string `json:"securityGroup,omitempty"`
    
    SessionToken string `json:"sessionToken,omitempty"`
    
    SpotPrice string `json:"spotPrice,omitempty"`
    
    SshUser string `json:"sshUser,omitempty"`
    
    SubnetId string `json:"subnetId,omitempty"`
    
    VpcId string `json:"vpcId,omitempty"`
    
    Zone string `json:"zone,omitempty"`
    
}

type Amazonec2ConfigCollection struct {
	Collection
	Data []Amazonec2Config `json:"data,omitempty"`
}

type Amazonec2ConfigClient struct {
	rancherClient *RancherClient
}

type Amazonec2ConfigOperations interface {
	List(opts *ListOpts) (*Amazonec2ConfigCollection, error)
	Create(opts *Amazonec2Config) (*Amazonec2Config, error)
	Update(existing *Amazonec2Config, updates interface{}) (*Amazonec2Config, error)
	ById(id string) (*Amazonec2Config, error)
	Delete(container *Amazonec2Config) error
}

func newAmazonec2ConfigClient(rancherClient *RancherClient) *Amazonec2ConfigClient {
	return &Amazonec2ConfigClient{
		rancherClient: rancherClient,
	}
}

func (c *Amazonec2ConfigClient) Create(container *Amazonec2Config) (*Amazonec2Config, error) {
	resp := &Amazonec2Config{}
	err := c.rancherClient.doCreate(AMAZONEC2CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *Amazonec2ConfigClient) Update(existing *Amazonec2Config, updates interface{}) (*Amazonec2Config, error) {
	resp := &Amazonec2Config{}
	err := c.rancherClient.doUpdate(AMAZONEC2CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *Amazonec2ConfigClient) List(opts *ListOpts) (*Amazonec2ConfigCollection, error) {
	resp := &Amazonec2ConfigCollection{}
	err := c.rancherClient.doList(AMAZONEC2CONFIG_TYPE, opts, resp)
	return resp, err
}

func (c *Amazonec2ConfigClient) ById(id string) (*Amazonec2Config, error) {
	resp := &Amazonec2Config{}
	err := c.rancherClient.doById(AMAZONEC2CONFIG_TYPE, id, resp)
	return resp, err
}

func (c *Amazonec2ConfigClient) Delete(container *Amazonec2Config) error {
	return c.rancherClient.doResourceDelete(AMAZONEC2CONFIG_TYPE, &container.Resource)
}
