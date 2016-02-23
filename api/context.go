package api

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/rancher/go-rancher/client"
)

type key int

const (
	contextKey key = 0
)

type ApiContext struct {
	r                 *http.Request
	Schemas           *client.Schemas
	UrlBuilder        UrlBuilder
	apiResponseWriter ApiResponseWriter
	responseWriter    http.ResponseWriter
	Policy            *Policy
	IDFormatter       IDFormatter
}


type Policy struct {
	AccountID int64 `json:"accountId"`
	AuthenticatedAsAccountID int64 `json:"authenticatedAsAccountId"`
	Identities []Identity `json:"identities"`
	Username string `json:"username"`
}

type Identity struct {
	ExternalID string `json:"externalId"`
	ExternalIDType string `json:"externalIdType"`
	ID string `json:"id"`
}


func GetApiContext(r *http.Request) *ApiContext {
	if rv := context.Get(r, contextKey); rv != nil {
		return rv.(*ApiContext)
	}
	return nil
}

func CreateApiContext(rw http.ResponseWriter, r *http.Request, schemas *client.Schemas) error {
	urlBuilder, err := NewUrlBuilder(r, schemas)
	if err != nil {
		return err
	}

	apiContext := &ApiContext{
		r:                 r,
		Schemas:           schemas,
		UrlBuilder:        urlBuilder,
		apiResponseWriter: parseResponseType(r),
		responseWriter:    rw,
	}

	context.Set(r, contextKey, apiContext)
	return nil
}
