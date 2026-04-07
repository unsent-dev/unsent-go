package unsent

import "fmt"

// ProviderConnectionsClient handles provider-connections API operations
type ProviderConnectionsClient struct {
	client *Client
}

func (c *ProviderConnectionsClient) List() (*interface{}, *APIError) {
	return Get[interface{}](c.client, "/provider-connections")
}

func (c *ProviderConnectionsClient) Create(payload interface{}) (*interface{}, *APIError) {
	return Post[interface{}](c.client, "/provider-connections", payload)
}

func (c *ProviderConnectionsClient) Delete(id string) (*interface{}, *APIError) {
	return Delete[interface{}](c.client, fmt.Sprintf("/provider-connections/%s", id), nil)
}
