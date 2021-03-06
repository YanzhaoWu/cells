// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new workspace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workspace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Changes Sends back all changes since a given sequence ID.

This plugin requires **meta.syncable** active on the workspace.

*/
func (a *Client) Changes(params *ChangesParams, authInfo runtime.ClientAuthInfoWriter) (*ChangesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changes",
		Method:             "GET",
		PathPattern:        "/workspaces/{workspaceId}/changes/{sequenceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangesOK), nil

}

/*
GetWorkspaceInfo Get information about the given workspace.
Info can be gathered via various plugins. Pass the expected metadata type via the X-Pydio-Ws-Info header. Currently supported values are quota|info|changes

*/
func (a *Client) GetWorkspaceInfo(params *GetWorkspaceInfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkspaceInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkspaceInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkspaceInfo",
		Method:             "GET",
		PathPattern:        "/workspaces/{workspaceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkspaceInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkspaceInfoOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
