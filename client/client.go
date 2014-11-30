package client

type RancherClient struct {
	Opts    *ClientOpts
	Schemas *Schemas
	Types   map[string]Schema

	Container *ContainerClient
	/*
		Generate the below for volume

		Volume *VolumeClient
	*/
}

func constructClient() *RancherClient {
	client := &RancherClient{
		Types: map[string]Schema{},
	}

	client.Container = newContainerClient(client)
	/*
		Generate the below for volume

		client.Volume = newVolumeClient(client)
	*/

	return client
}

func (rancherClient *RancherClient) Delete(schemaType string, existing *Resource) error {
	return rancherClient.doResourceDelete(schemaType, existing)
}

func (rancherClient *RancherClient) Create(schemaType string, createObj interface{}, respObject interface{}) error {
	return rancherClient.doCreate(schemaType, createObj, respObject)
}

func (rancherClient *RancherClient) Update(schemaType string, existing *Resource, updates interface{}, respObject interface{}) error {
	return rancherClient.doUpdate(schemaType, existing, updates, respObject)
}

func (rancherClient *RancherClient) List(schemaType string, opts *ListOpts, respObject interface{}) error {
	return rancherClient.doList(schemaType, opts, respObject)
}

func (rancherClient *RancherClient) Action(schemaType string, action string, reqObj interface{}, respObj interface{}) error {
	return rancherClient.doAction(schemaType, action, reqObj, respObj)
}
