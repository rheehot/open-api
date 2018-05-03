// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSubmissionParams creates a new DeleteSubmissionParams object
// with the default values initialized.
func NewDeleteSubmissionParams() *DeleteSubmissionParams {
	var ()
	return &DeleteSubmissionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubmissionParamsWithTimeout creates a new DeleteSubmissionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSubmissionParamsWithTimeout(timeout time.Duration) *DeleteSubmissionParams {
	var ()
	return &DeleteSubmissionParams{

		timeout: timeout,
	}
}

// NewDeleteSubmissionParamsWithContext creates a new DeleteSubmissionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSubmissionParamsWithContext(ctx context.Context) *DeleteSubmissionParams {
	var ()
	return &DeleteSubmissionParams{

		Context: ctx,
	}
}

// NewDeleteSubmissionParamsWithHTTPClient creates a new DeleteSubmissionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSubmissionParamsWithHTTPClient(client *http.Client) *DeleteSubmissionParams {
	var ()
	return &DeleteSubmissionParams{
		HTTPClient: client,
	}
}

/*DeleteSubmissionParams contains all the parameters to send to the API endpoint
for the delete submission operation typically these are written to a http.Request
*/
type DeleteSubmissionParams struct {

	/*SubmissionID*/
	SubmissionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete submission params
func (o *DeleteSubmissionParams) WithTimeout(timeout time.Duration) *DeleteSubmissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete submission params
func (o *DeleteSubmissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete submission params
func (o *DeleteSubmissionParams) WithContext(ctx context.Context) *DeleteSubmissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete submission params
func (o *DeleteSubmissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete submission params
func (o *DeleteSubmissionParams) WithHTTPClient(client *http.Client) *DeleteSubmissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete submission params
func (o *DeleteSubmissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubmissionID adds the submissionID to the delete submission params
func (o *DeleteSubmissionParams) WithSubmissionID(submissionID string) *DeleteSubmissionParams {
	o.SetSubmissionID(submissionID)
	return o
}

// SetSubmissionID adds the submissionId to the delete submission params
func (o *DeleteSubmissionParams) SetSubmissionID(submissionID string) {
	o.SubmissionID = submissionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubmissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param submission_id
	if err := r.SetPathParam("submission_id", o.SubmissionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
