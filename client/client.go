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
