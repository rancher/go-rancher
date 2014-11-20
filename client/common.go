package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	SELF       = "self"
	COLLECTION = "collection"
)

type ClientOpts struct {
	Url       string
	AccessKey string
	SecretKey string
}

func contains(array []string, item string) bool {
	for _, check := range array {
		if check == item {
			return true
		}
	}

	return false
}

func appendFilters(urlString string, filters map[string]interface{}) (string, error) {
	if len(filters) == 0 {
		return urlString, nil
	}

	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range filters {
		q.Set(k, fmt.Sprintf("%v", v))
	}

	return u.String(), nil
}

func NewRancherClient(opts *ClientOpts) (*RancherClient, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", opts.Url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(opts.AccessKey, opts.SecretKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Bad response from [%s], go [%d]", opts.Url, resp.StatusCode))
	}

	schemasUrls := resp.Header.Get("X-API-Schemas")
	if len(schemasUrls) == 0 {
		return nil, errors.New("Failed to find schema at [" + opts.Url + "]")
	}

	if schemasUrls != opts.Url {
		req, err = http.NewRequest("GET", schemasUrls, nil)
		if err != nil {
			return nil, err
		}

		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return nil, errors.New(fmt.Sprintf("Bad response from [%s], go [%d]", opts.Url, resp.StatusCode))
		}
	}

	var schemas Schemas
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &schemas)
	if err != nil {
		return nil, err
	}

	rancherClient := constructClient()
	rancherClient.Opts = opts
	rancherClient.Schemas = &schemas

	for _, schema := range schemas.Data {
		rancherClient.Types[schema.Id] = schema
	}

	return rancherClient, nil
}

func NewListOpts() *ListOpts {
	return &ListOpts{
		Filters: map[string]interface{}{},
	}
}

func (rancherClient *RancherClient) setupRequest(req *http.Request) {
	req.SetBasicAuth(rancherClient.Opts.AccessKey, rancherClient.Opts.SecretKey)
}

func (rancherClient *RancherClient) newHttpClient() *http.Client {
	return &http.Client{}
}

func (rancherClient *RancherClient) doDelete(url string) error {
	client := rancherClient.newHttpClient()
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	rancherClient.setupRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("Bad response from [%s], go [%d]", url, resp.StatusCode))
	}

	return nil
}

func (rancherClient *RancherClient) doGet(url string, opts *ListOpts, respObject interface{}) error {
	if opts == nil {
		opts = NewListOpts()
	}

	url, err := appendFilters(url, opts.Filters)
	if err != nil {
		return err
	}

	client := rancherClient.newHttpClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	rancherClient.setupRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Bad response from [%s], go [%d]", url, resp.StatusCode))
	}

	byteContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteContent, respObject)
}

func (rancherClient *RancherClient) doList(schemaType string, opts *ListOpts, respObject interface{}) error {
	schema, ok := rancherClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	if !contains(schema.CollectionMethods, "GET") {
		return errors.New("Resource type [" + schemaType + "] is not listable")
	}

	collectionUrl, ok := schema.Links[COLLECTION]
	if !ok {
		return errors.New("Failed to find collection URL for [" + schemaType + "]")
	}

	return rancherClient.doGet(collectionUrl, opts, respObject)
}

func (rancherClient *RancherClient) doModify(method string, url string, createObj interface{}, respObject interface{}) error {
	bodyContent, err := json.Marshal(createObj)
	if err != nil {
		return err
	}

	client := rancherClient.newHttpClient()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyContent))
	if err != nil {
		return err
	}

	rancherClient.setupRequest(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", string(len(bodyContent)))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("Bad response from [%s], go [%d]", url, resp.StatusCode))
	}

	byteContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteContent, respObject)
}

func (rancherClient *RancherClient) doCreate(schemaType string, createObj interface{}, respObject interface{}) error {
	if createObj == nil {
		createObj = map[string]string{}
	}

	schema, ok := rancherClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	if !contains(schema.CollectionMethods, "POST") {
		return errors.New("Resource type [" + schemaType + "] is not creatable")
	}

	collectionUrl, ok := schema.Links[COLLECTION]
	if !ok {
		return errors.New("Failed to find collection URL for [" + schemaType + "]")
	}

	return rancherClient.doModify("POST", collectionUrl, createObj, respObject)
}

func (rancherClient *RancherClient) doUpdate(schemaType string, existing *Resource, updates interface{}, respObject interface{}) error {
	if existing == nil {
		return errors.New("Existing object is nil")
	}

	selfUrl, ok := existing.Links[SELF]
	if !ok {
		return errors.New(fmt.Sprintf("Failed to find self URL of [%v]", existing))
	}

	if updates == nil {
		updates = map[string]string{}
	}

	schema, ok := rancherClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	if !contains(schema.ResourceMethods, "PUT") {
		return errors.New("Resource type [" + schemaType + "] is not updatable")
	}

	return rancherClient.doModify("PUT", selfUrl, updates, respObject)
}

func (rancherClient *RancherClient) doById(schemaType string, id string, respObject interface{}) error {
	schema, ok := rancherClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	if !contains(schema.ResourceMethods, "GET") {
		return errors.New("Resource type [" + schemaType + "] can not be looked up by ID")
	}

	collectionUrl, ok := schema.Links[COLLECTION]
	if !ok {
		return errors.New("Failed to find collection URL for [" + schemaType + "]")
	}

	err := rancherClient.doGet(collectionUrl+"/"+id, nil, respObject)
	//TODO check for 404 and return nil, nil
	return err
}

func (rancherClient *RancherClient) doResourceDelete(schemaType string, existing *Resource) error {
	schema, ok := rancherClient.Types[schemaType]
	if !ok {
		return errors.New("Unknown schema type [" + schemaType + "]")
	}

	if !contains(schema.ResourceMethods, "DELETE") {
		return errors.New("Resource type [" + schemaType + "] can not be deleted")
	}

	selfUrl, ok := existing.Links[SELF]
	if !ok {
		return errors.New(fmt.Sprintf("Failed to find self URL of [%v]", existing))
	}

	return rancherClient.doDelete(selfUrl)
}
