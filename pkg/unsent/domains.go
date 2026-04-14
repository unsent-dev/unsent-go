// @manual
package unsent

import "fmt"

// DomainsClient handles domain-related API operations
type DomainsClient struct {
	client *Client
}

// List retrieves all domains
func (c *DomainsClient) List() (*[]Domain, *APIError) {
	return Get[[]Domain](c.client, "/domains")
}

// Get retrieves a domain by ID
func (c *DomainsClient) Get(domainID string) (*Domain, *APIError) {
	return Get[Domain](c.client, fmt.Sprintf("/domains/%s", domainID))
}

// Create creates a new domain
func (c *DomainsClient) Create(payload CreateDomainJSONBody) (*DomainCreateResponse, *APIError) {
	return Post[DomainCreateResponse](c.client, "/domains", payload)
}

// Verify triggers domain verification
func (c *DomainsClient) Verify(domainID string) (*DomainVerifyResponse, *APIError) {
	return Put[DomainVerifyResponse](c.client, fmt.Sprintf("/domains/%s/verify", domainID), nil)
}

// Delete deletes a domain
func (c *DomainsClient) Delete(domainID string) (*DomainDeleteResponse, *APIError) {
	return Delete[DomainDeleteResponse](c.client, fmt.Sprintf("/domains/%s", domainID), nil)
}

// GetAnalytics retrieves analytics for a specific domain
func (c *DomainsClient) GetAnalytics(id string, params GetDomainAnalyticsParams) (*interface{}, *APIError) {
	path := fmt.Sprintf("/domains/%s/analytics", id)
	
	// Build query parameters
	query := buildQueryParams(map[string]interface{}{
		"period": params.Period,
	})
	
	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}
	
	return Get[interface{}](c.client, path)
}

// GetStats retrieves statistics for a specific domain
func (c *DomainsClient) GetStats(id string, params GetDomainStatsParams) (*interface{}, *APIError) {
	path := fmt.Sprintf("/domains/%s/stats", id)

	// Build query parameters
	query := buildQueryParams(map[string]interface{}{
		"startDate": params.StartDate,
		"endDate":   params.EndDate,
	})

	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}

	return Get[interface{}](c.client, path)
}

// ListRoutes retrieves all routes for a domain
func (c *DomainsClient) ListRoutes(domainID string) (*[]DomainRoute, *APIError) {
	return Get[[]DomainRoute](c.client, fmt.Sprintf("/domains/%s/routes", domainID))
}

// AddRoute adds a route to a domain
func (c *DomainsClient) AddRoute(domainID string, payload AddDomainRouteJSONBody) (*DomainRoute, *APIError) {
	return Post[DomainRoute](c.client, fmt.Sprintf("/domains/%s/routes", domainID), payload)
}

// UpdateRoute updates a route on a domain
func (c *DomainsClient) UpdateRoute(domainID string, routeID string, payload UpdateDomainRouteJSONBody) (*DomainRouteUpdateResponse, *APIError) {
	return Patch[DomainRouteUpdateResponse](c.client, fmt.Sprintf("/domains/%s/routes/%s", domainID, routeID), payload)
}

// DeleteRoute deletes a route from a domain
func (c *DomainsClient) DeleteRoute(domainID string, routeID string) (*DomainRouteDeleteResponse, *APIError) {
	return Delete[DomainRouteDeleteResponse](c.client, fmt.Sprintf("/domains/%s/routes/%s", domainID, routeID), nil)
}

