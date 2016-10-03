package api

import "github.com/rancher/go-rancher/client"

type IDFormatter interface {
	FormatID(id, idType string, schemas *client.Schemas) string
	ParseID(id string) string
}
