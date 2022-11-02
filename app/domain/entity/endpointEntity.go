package entity

type GetEndpointsRequest struct {
	EndpointType string
	Name         string
}

type Endpoint struct {
	Id           string
	Name         string
	EndpointType string
	Description  string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    *string
}

type GetEndpointsResponse struct {
	Endpoints []Endpoint
}

type InsertEndpointRequest struct {
	Name         string
	EndpointType string
	Description  string
}

type InsertEndpointResponse struct {
	Id           string
	Name         string
	EndpointType string
	Description  string
}

type UpdateEndpointRequest struct {
	Id           string
	Name         string
	EndpointType string
	Description  string
}

type DeleteEndpointRequest struct {
	Id         string
	SoftDelete bool
}
