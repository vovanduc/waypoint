// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewWaypointUIListReleases3Params creates a new WaypointUIListReleases3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointUIListReleases3Params() *WaypointUIListReleases3Params {
	return &WaypointUIListReleases3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointUIListReleases3ParamsWithTimeout creates a new WaypointUIListReleases3Params object
// with the ability to set a timeout on a request.
func NewWaypointUIListReleases3ParamsWithTimeout(timeout time.Duration) *WaypointUIListReleases3Params {
	return &WaypointUIListReleases3Params{
		timeout: timeout,
	}
}

// NewWaypointUIListReleases3ParamsWithContext creates a new WaypointUIListReleases3Params object
// with the ability to set a context for a request.
func NewWaypointUIListReleases3ParamsWithContext(ctx context.Context) *WaypointUIListReleases3Params {
	return &WaypointUIListReleases3Params{
		Context: ctx,
	}
}

// NewWaypointUIListReleases3ParamsWithHTTPClient creates a new WaypointUIListReleases3Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointUIListReleases3ParamsWithHTTPClient(client *http.Client) *WaypointUIListReleases3Params {
	return &WaypointUIListReleases3Params{
		HTTPClient: client,
	}
}

/*
WaypointUIListReleases3Params contains all the parameters to send to the API endpoint

	for the waypoint UI list releases3 operation.

	Typically these are written to a http.Request.
*/
type WaypointUIListReleases3Params struct {

	// ApplicationApplication.
	ApplicationApplication *string

	// ApplicationProject.
	ApplicationProject string

	// OrderDesc.
	OrderDesc *bool

	/* OrderLimit.

	   Limit the number of results.

	   Format: int64
	*/
	OrderLimit *int64

	/* OrderOrder.

	   Order for the results.

	   Default: "UNSET"
	*/
	OrderOrder *string

	/* PhysicalState.

	     The physical state to filter for. If this is zero or unset then no
	filtering on physical state will be done.

	     Default: "UNKNOWN"
	*/
	PhysicalState *string

	// WorkspaceWorkspace.
	WorkspaceWorkspace *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint UI list releases3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUIListReleases3Params) WithDefaults() *WaypointUIListReleases3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint UI list releases3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointUIListReleases3Params) SetDefaults() {
	var (
		orderOrderDefault = string("UNSET")

		physicalStateDefault = string("UNKNOWN")
	)

	val := WaypointUIListReleases3Params{
		OrderOrder:    &orderOrderDefault,
		PhysicalState: &physicalStateDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) WithTimeout(timeout time.Duration) *WaypointUIListReleases3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) WithContext(ctx context.Context) *WaypointUIListReleases3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) WithHTTPClient(client *http.Client) *WaypointUIListReleases3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationApplication adds the applicationApplication to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) WithApplicationApplication(applicationApplication *string) *WaypointUIListReleases3Params {
	o.SetApplicationApplication(applicationApplication)
	return o
}

// SetApplicationApplication adds the applicationApplication to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) SetApplicationApplication(applicationApplication *string) {
	o.ApplicationApplication = applicationApplication
}

// WithApplicationProject adds the applicationProject to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) WithApplicationProject(applicationProject string) *WaypointUIListReleases3Params {
	o.SetApplicationProject(applicationProject)
	return o
}

// SetApplicationProject adds the applicationProject to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) SetApplicationProject(applicationProject string) {
	o.ApplicationProject = applicationProject
}

// WithOrderDesc adds the orderDesc to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) WithOrderDesc(orderDesc *bool) *WaypointUIListReleases3Params {
	o.SetOrderDesc(orderDesc)
	return o
}

// SetOrderDesc adds the orderDesc to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) SetOrderDesc(orderDesc *bool) {
	o.OrderDesc = orderDesc
}

// WithOrderLimit adds the orderLimit to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) WithOrderLimit(orderLimit *int64) *WaypointUIListReleases3Params {
	o.SetOrderLimit(orderLimit)
	return o
}

// SetOrderLimit adds the orderLimit to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) SetOrderLimit(orderLimit *int64) {
	o.OrderLimit = orderLimit
}

// WithOrderOrder adds the orderOrder to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) WithOrderOrder(orderOrder *string) *WaypointUIListReleases3Params {
	o.SetOrderOrder(orderOrder)
	return o
}

// SetOrderOrder adds the orderOrder to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) SetOrderOrder(orderOrder *string) {
	o.OrderOrder = orderOrder
}

// WithPhysicalState adds the physicalState to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) WithPhysicalState(physicalState *string) *WaypointUIListReleases3Params {
	o.SetPhysicalState(physicalState)
	return o
}

// SetPhysicalState adds the physicalState to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) SetPhysicalState(physicalState *string) {
	o.PhysicalState = physicalState
}

// WithWorkspaceWorkspace adds the workspaceWorkspace to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) WithWorkspaceWorkspace(workspaceWorkspace *string) *WaypointUIListReleases3Params {
	o.SetWorkspaceWorkspace(workspaceWorkspace)
	return o
}

// SetWorkspaceWorkspace adds the workspaceWorkspace to the waypoint UI list releases3 params
func (o *WaypointUIListReleases3Params) SetWorkspaceWorkspace(workspaceWorkspace *string) {
	o.WorkspaceWorkspace = workspaceWorkspace
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointUIListReleases3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationApplication != nil {

		// query param application.application
		var qrApplicationApplication string

		if o.ApplicationApplication != nil {
			qrApplicationApplication = *o.ApplicationApplication
		}
		qApplicationApplication := qrApplicationApplication
		if qApplicationApplication != "" {

			if err := r.SetQueryParam("application.application", qApplicationApplication); err != nil {
				return err
			}
		}
	}

	// path param application.project
	if err := r.SetPathParam("application.project", o.ApplicationProject); err != nil {
		return err
	}

	if o.OrderDesc != nil {

		// query param order.desc
		var qrOrderDesc bool

		if o.OrderDesc != nil {
			qrOrderDesc = *o.OrderDesc
		}
		qOrderDesc := swag.FormatBool(qrOrderDesc)
		if qOrderDesc != "" {

			if err := r.SetQueryParam("order.desc", qOrderDesc); err != nil {
				return err
			}
		}
	}

	if o.OrderLimit != nil {

		// query param order.limit
		var qrOrderLimit int64

		if o.OrderLimit != nil {
			qrOrderLimit = *o.OrderLimit
		}
		qOrderLimit := swag.FormatInt64(qrOrderLimit)
		if qOrderLimit != "" {

			if err := r.SetQueryParam("order.limit", qOrderLimit); err != nil {
				return err
			}
		}
	}

	if o.OrderOrder != nil {

		// query param order.order
		var qrOrderOrder string

		if o.OrderOrder != nil {
			qrOrderOrder = *o.OrderOrder
		}
		qOrderOrder := qrOrderOrder
		if qOrderOrder != "" {

			if err := r.SetQueryParam("order.order", qOrderOrder); err != nil {
				return err
			}
		}
	}

	if o.PhysicalState != nil {

		// query param physical_state
		var qrPhysicalState string

		if o.PhysicalState != nil {
			qrPhysicalState = *o.PhysicalState
		}
		qPhysicalState := qrPhysicalState
		if qPhysicalState != "" {

			if err := r.SetQueryParam("physical_state", qPhysicalState); err != nil {
				return err
			}
		}
	}

	if o.WorkspaceWorkspace != nil {

		// query param workspace.workspace
		var qrWorkspaceWorkspace string

		if o.WorkspaceWorkspace != nil {
			qrWorkspaceWorkspace = *o.WorkspaceWorkspace
		}
		qWorkspaceWorkspace := qrWorkspaceWorkspace
		if qWorkspaceWorkspace != "" {

			if err := r.SetQueryParam("workspace.workspace", qWorkspaceWorkspace); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}