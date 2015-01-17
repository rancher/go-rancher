package client

type RancherClient struct {
    RancherBaseClient
	
    ApiVersion *ApiVersionClient
    Error *ErrorClient
    AgentInstanceProvider *AgentInstanceProviderClient
    DnsService *DnsServiceClient
    DhcpService *DhcpServiceClient
    IpsecTunnelService *IpsecTunnelServiceClient
    LinkService *LinkServiceClient
    PortService *PortServiceClient
    MetadataService *MetadataServiceClient
    GatewayService *GatewayServiceClient
    HostNatGatewayService *HostNatGatewayServiceClient
    Subscribe *SubscribeClient
    Publish *PublishClient
    ConfigContent *ConfigContentClient
    VirtualMachine *VirtualMachineClient
    Container *ContainerClient
    ApiKey *ApiKeyClient
    SshKey *SshKeyClient
    InstanceStop *InstanceStopClient
    InstanceConsole *InstanceConsoleClient
    InstanceConsoleInput *InstanceConsoleInputClient
    IpAddressAssociateInput *IpAddressAssociateInputClient
    SubnetIpPool *SubnetIpPoolClient
    Account *AccountClient
    Agent *AgentClient
    AgentGroup *AgentGroupClient
    ConfigItem *ConfigItemClient
    ConfigItemStatus *ConfigItemStatusClient
    Credential *CredentialClient
    CredentialInstanceMap *CredentialInstanceMapClient
    Data *DataClient
    Databasechangelog *DatabasechangelogClient
    Databasechangeloglock *DatabasechangeloglockClient
    ExternalHandler *ExternalHandlerClient
    ExternalHandlerExternalHandlerProcessMap *ExternalHandlerExternalHandlerProcessMapClient
    ExternalHandlerProcess *ExternalHandlerProcessClient
    GenericObject *GenericObjectClient
    Host *HostClient
    HostIpAddressMap *HostIpAddressMapClient
    HostVnetMap *HostVnetMapClient
    Image *ImageClient
    ImageStoragePoolMap *ImageStoragePoolMapClient
    Instance *InstanceClient
    InstanceHostMap *InstanceHostMapClient
    InstanceLink *InstanceLinkClient
    IpAddress *IpAddressClient
    IpAddressNicMap *IpAddressNicMapClient
    IpAssociation *IpAssociationClient
    IpPool *IpPoolClient
    Mount *MountClient
    Network *NetworkClient
    NetworkService *NetworkServiceClient
    NetworkServiceProvider *NetworkServiceProviderClient
    NetworkServiceProviderInstanceMap *NetworkServiceProviderInstanceMapClient
    Nic *NicClient
    PhysicalHost *PhysicalHostClient
    Port *PortClient
    ProcessExecution *ProcessExecutionClient
    ProcessInstance *ProcessInstanceClient
    ResourcePool *ResourcePoolClient
    Setting *SettingClient
    StoragePool *StoragePoolClient
    StoragePoolHostMap *StoragePoolHostMapClient
    Subnet *SubnetClient
    SubnetVnetMap *SubnetVnetMapClient
    Task *TaskClient
    TaskInstance *TaskInstanceClient
    Vnet *VnetClient
    Volume *VolumeClient
    VolumeStoragePoolMap *VolumeStoragePoolMapClient
    Zone *ZoneClient
    TypeDocumentation *TypeDocumentationClient
    FieldDocumentation *FieldDocumentationClient
    ContainerExec *ContainerExecClient
    ContainerExecOutput *ContainerExecOutputClient
    ActiveSetting *ActiveSettingClient
    ExtensionImplementation *ExtensionImplementationClient
    ExtensionPoint *ExtensionPointClient
    ProcessDefinition *ProcessDefinitionClient
    ResourceDefinition *ResourceDefinitionClient
    StateTransition *StateTransitionClient
    StatsAccess *StatsAccessClient
    HostOnlyNetwork *HostOnlyNetworkClient
    Register *RegisterClient
    RegistrationToken *RegistrationTokenClient
    Authorized *AuthorizedClient
}

func constructClient() *RancherClient {
	client := &RancherClient{
		RancherBaseClient: RancherBaseClient{
			Types: map[string]Schema{},
		},
	}

    
    client.ApiVersion = newApiVersionClient(client)
    client.Error = newErrorClient(client)
    client.AgentInstanceProvider = newAgentInstanceProviderClient(client)
    client.DnsService = newDnsServiceClient(client)
    client.DhcpService = newDhcpServiceClient(client)
    client.IpsecTunnelService = newIpsecTunnelServiceClient(client)
    client.LinkService = newLinkServiceClient(client)
    client.PortService = newPortServiceClient(client)
    client.MetadataService = newMetadataServiceClient(client)
    client.GatewayService = newGatewayServiceClient(client)
    client.HostNatGatewayService = newHostNatGatewayServiceClient(client)
    client.Subscribe = newSubscribeClient(client)
    client.Publish = newPublishClient(client)
    client.ConfigContent = newConfigContentClient(client)
    client.VirtualMachine = newVirtualMachineClient(client)
    client.Container = newContainerClient(client)
    client.ApiKey = newApiKeyClient(client)
    client.SshKey = newSshKeyClient(client)
    client.InstanceStop = newInstanceStopClient(client)
    client.InstanceConsole = newInstanceConsoleClient(client)
    client.InstanceConsoleInput = newInstanceConsoleInputClient(client)
    client.IpAddressAssociateInput = newIpAddressAssociateInputClient(client)
    client.SubnetIpPool = newSubnetIpPoolClient(client)
    client.Account = newAccountClient(client)
    client.Agent = newAgentClient(client)
    client.AgentGroup = newAgentGroupClient(client)
    client.ConfigItem = newConfigItemClient(client)
    client.ConfigItemStatus = newConfigItemStatusClient(client)
    client.Credential = newCredentialClient(client)
    client.CredentialInstanceMap = newCredentialInstanceMapClient(client)
    client.Data = newDataClient(client)
    client.Databasechangelog = newDatabasechangelogClient(client)
    client.Databasechangeloglock = newDatabasechangeloglockClient(client)
    client.ExternalHandler = newExternalHandlerClient(client)
    client.ExternalHandlerExternalHandlerProcessMap = newExternalHandlerExternalHandlerProcessMapClient(client)
    client.ExternalHandlerProcess = newExternalHandlerProcessClient(client)
    client.GenericObject = newGenericObjectClient(client)
    client.Host = newHostClient(client)
    client.HostIpAddressMap = newHostIpAddressMapClient(client)
    client.HostVnetMap = newHostVnetMapClient(client)
    client.Image = newImageClient(client)
    client.ImageStoragePoolMap = newImageStoragePoolMapClient(client)
    client.Instance = newInstanceClient(client)
    client.InstanceHostMap = newInstanceHostMapClient(client)
    client.InstanceLink = newInstanceLinkClient(client)
    client.IpAddress = newIpAddressClient(client)
    client.IpAddressNicMap = newIpAddressNicMapClient(client)
    client.IpAssociation = newIpAssociationClient(client)
    client.IpPool = newIpPoolClient(client)
    client.Mount = newMountClient(client)
    client.Network = newNetworkClient(client)
    client.NetworkService = newNetworkServiceClient(client)
    client.NetworkServiceProvider = newNetworkServiceProviderClient(client)
    client.NetworkServiceProviderInstanceMap = newNetworkServiceProviderInstanceMapClient(client)
    client.Nic = newNicClient(client)
    client.PhysicalHost = newPhysicalHostClient(client)
    client.Port = newPortClient(client)
    client.ProcessExecution = newProcessExecutionClient(client)
    client.ProcessInstance = newProcessInstanceClient(client)
    client.ResourcePool = newResourcePoolClient(client)
    client.Setting = newSettingClient(client)
    client.StoragePool = newStoragePoolClient(client)
    client.StoragePoolHostMap = newStoragePoolHostMapClient(client)
    client.Subnet = newSubnetClient(client)
    client.SubnetVnetMap = newSubnetVnetMapClient(client)
    client.Task = newTaskClient(client)
    client.TaskInstance = newTaskInstanceClient(client)
    client.Vnet = newVnetClient(client)
    client.Volume = newVolumeClient(client)
    client.VolumeStoragePoolMap = newVolumeStoragePoolMapClient(client)
    client.Zone = newZoneClient(client)
    client.TypeDocumentation = newTypeDocumentationClient(client)
    client.FieldDocumentation = newFieldDocumentationClient(client)
    client.ContainerExec = newContainerExecClient(client)
    client.ContainerExecOutput = newContainerExecOutputClient(client)
    client.ActiveSetting = newActiveSettingClient(client)
    client.ExtensionImplementation = newExtensionImplementationClient(client)
    client.ExtensionPoint = newExtensionPointClient(client)
    client.ProcessDefinition = newProcessDefinitionClient(client)
    client.ResourceDefinition = newResourceDefinitionClient(client)
    client.StateTransition = newStateTransitionClient(client)
    client.StatsAccess = newStatsAccessClient(client)
    client.HostOnlyNetwork = newHostOnlyNetworkClient(client)
    client.Register = newRegisterClient(client)
    client.RegistrationToken = newRegistrationTokenClient(client)
    client.Authorized = newAuthorizedClient(client) 


	return client
}

func NewRancherClient(opts *ClientOpts) (*RancherClient, error) {
    client := constructClient()
        
    err := setupRancherBaseClient(&client.RancherBaseClient, opts)
    if err != nil {
        return nil, err
    }

    return client, nil
}
