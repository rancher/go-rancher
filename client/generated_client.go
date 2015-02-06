package client

type RancherClient struct {
    RancherBaseClient
	
    Subscribe *SubscribeClient
    Publish *PublishClient
    RestartPolicy *RestartPolicyClient
    LoadBalancerHealthCheck *LoadBalancerHealthCheckClient
    LoadBalancerPolicy *LoadBalancerPolicyClient
    GlobalLoadBalancerPolicy *GlobalLoadBalancerPolicyClient
    GlobalLoadBalancerHealthCheck *GlobalLoadBalancerHealthCheckClient
    Container *ContainerClient
    ApiKey *ApiKeyClient
    InstanceStop *InstanceStopClient
    InstanceConsole *InstanceConsoleClient
    InstanceConsoleInput *InstanceConsoleInputClient
    IpAddressAssociateInput *IpAddressAssociateInputClient
    Project *ProjectClient
    AddRemoveLoadBalancerListenerInput *AddRemoveLoadBalancerListenerInputClient
    AddRemoveLoadBalancerTargetInput *AddRemoveLoadBalancerTargetInputClient
    AddLoadBalancerInput *AddLoadBalancerInputClient
    RemoveLoadBalancerInput *RemoveLoadBalancerInputClient
    AddRemoveLoadBalancerHostInput *AddRemoveLoadBalancerHostInputClient
    Account *AccountClient
    Agent *AgentClient
    ConfigItem *ConfigItemClient
    ConfigItemStatus *ConfigItemStatusClient
    Credential *CredentialClient
    Databasechangelog *DatabasechangelogClient
    Databasechangeloglock *DatabasechangeloglockClient
    ExternalHandler *ExternalHandlerClient
    ExternalHandlerExternalHandlerProcessMap *ExternalHandlerExternalHandlerProcessMapClient
    ExternalHandlerProcess *ExternalHandlerProcessClient
    GlobalLoadBalancer *GlobalLoadBalancerClient
    Host *HostClient
    Image *ImageClient
    Instance *InstanceClient
    InstanceLink *InstanceLinkClient
    IpAddress *IpAddressClient
    LoadBalancer *LoadBalancerClient
    LoadBalancerConfig *LoadBalancerConfigClient
    LoadBalancerListener *LoadBalancerListenerClient
    LoadBalancerTarget *LoadBalancerTargetClient
    Mount *MountClient
    Network *NetworkClient
    PhysicalHost *PhysicalHostClient
    Port *PortClient
    ProcessExecution *ProcessExecutionClient
    ProcessInstance *ProcessInstanceClient
    Setting *SettingClient
    Task *TaskClient
    TaskInstance *TaskInstanceClient
    Volume *VolumeClient
    TypeDocumentation *TypeDocumentationClient
    ContainerExec *ContainerExecClient
    ContainerExecOutput *ContainerExecOutputClient
    ActiveSetting *ActiveSettingClient
    ExtensionImplementation *ExtensionImplementationClient
    ExtensionPoint *ExtensionPointClient
    ProcessDefinition *ProcessDefinitionClient
    ResourceDefinition *ResourceDefinitionClient
    Githubconfig *GithubconfigClient
    StatsAccess *StatsAccessClient
    VirtualboxConfig *VirtualboxConfigClient
    DigitaloceanConfig *DigitaloceanConfigClient
    MachineHost *MachineHostClient
    Register *RegisterClient
    RegistrationToken *RegistrationTokenClient
}

func constructClient() *RancherClient {
	client := &RancherClient{
		RancherBaseClient: RancherBaseClient{
			Types: map[string]Schema{},
		},
	}

    
    client.Subscribe = newSubscribeClient(client)
    client.Publish = newPublishClient(client)
    client.RestartPolicy = newRestartPolicyClient(client)
    client.LoadBalancerHealthCheck = newLoadBalancerHealthCheckClient(client)
    client.LoadBalancerPolicy = newLoadBalancerPolicyClient(client)
    client.GlobalLoadBalancerPolicy = newGlobalLoadBalancerPolicyClient(client)
    client.GlobalLoadBalancerHealthCheck = newGlobalLoadBalancerHealthCheckClient(client)
    client.Container = newContainerClient(client)
    client.ApiKey = newApiKeyClient(client)
    client.InstanceStop = newInstanceStopClient(client)
    client.InstanceConsole = newInstanceConsoleClient(client)
    client.InstanceConsoleInput = newInstanceConsoleInputClient(client)
    client.IpAddressAssociateInput = newIpAddressAssociateInputClient(client)
    client.Project = newProjectClient(client)
    client.AddRemoveLoadBalancerListenerInput = newAddRemoveLoadBalancerListenerInputClient(client)
    client.AddRemoveLoadBalancerTargetInput = newAddRemoveLoadBalancerTargetInputClient(client)
    client.AddLoadBalancerInput = newAddLoadBalancerInputClient(client)
    client.RemoveLoadBalancerInput = newRemoveLoadBalancerInputClient(client)
    client.AddRemoveLoadBalancerHostInput = newAddRemoveLoadBalancerHostInputClient(client)
    client.Account = newAccountClient(client)
    client.Agent = newAgentClient(client)
    client.ConfigItem = newConfigItemClient(client)
    client.ConfigItemStatus = newConfigItemStatusClient(client)
    client.Credential = newCredentialClient(client)
    client.Databasechangelog = newDatabasechangelogClient(client)
    client.Databasechangeloglock = newDatabasechangeloglockClient(client)
    client.ExternalHandler = newExternalHandlerClient(client)
    client.ExternalHandlerExternalHandlerProcessMap = newExternalHandlerExternalHandlerProcessMapClient(client)
    client.ExternalHandlerProcess = newExternalHandlerProcessClient(client)
    client.GlobalLoadBalancer = newGlobalLoadBalancerClient(client)
    client.Host = newHostClient(client)
    client.Image = newImageClient(client)
    client.Instance = newInstanceClient(client)
    client.InstanceLink = newInstanceLinkClient(client)
    client.IpAddress = newIpAddressClient(client)
    client.LoadBalancer = newLoadBalancerClient(client)
    client.LoadBalancerConfig = newLoadBalancerConfigClient(client)
    client.LoadBalancerListener = newLoadBalancerListenerClient(client)
    client.LoadBalancerTarget = newLoadBalancerTargetClient(client)
    client.Mount = newMountClient(client)
    client.Network = newNetworkClient(client)
    client.PhysicalHost = newPhysicalHostClient(client)
    client.Port = newPortClient(client)
    client.ProcessExecution = newProcessExecutionClient(client)
    client.ProcessInstance = newProcessInstanceClient(client)
    client.Setting = newSettingClient(client)
    client.Task = newTaskClient(client)
    client.TaskInstance = newTaskInstanceClient(client)
    client.Volume = newVolumeClient(client)
    client.TypeDocumentation = newTypeDocumentationClient(client)
    client.ContainerExec = newContainerExecClient(client)
    client.ContainerExecOutput = newContainerExecOutputClient(client)
    client.ActiveSetting = newActiveSettingClient(client)
    client.ExtensionImplementation = newExtensionImplementationClient(client)
    client.ExtensionPoint = newExtensionPointClient(client)
    client.ProcessDefinition = newProcessDefinitionClient(client)
    client.ResourceDefinition = newResourceDefinitionClient(client)
    client.Githubconfig = newGithubconfigClient(client)
    client.StatsAccess = newStatsAccessClient(client)
    client.VirtualboxConfig = newVirtualboxConfigClient(client)
    client.DigitaloceanConfig = newDigitaloceanConfigClient(client)
    client.MachineHost = newMachineHostClient(client)
    client.Register = newRegisterClient(client)
    client.RegistrationToken = newRegistrationTokenClient(client) 


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
