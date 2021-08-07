// Code generated by go-swagger; DO NOT EDIT.

package evaluate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/openmock/openmock/swagger_gen/models"
)

// NewEvaluateParams creates a new EvaluateParams object
// no default values defined in spec.
func NewEvaluateParams() EvaluateParams {

	return EvaluateParams{}
}

// EvaluateParams contains all the bound params for the evaluate operation
// typically these are obtained from a http.Request
//
// swagger:parameters evaluate
type EvaluateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*request to process
	  Required: true
	  In: body
	*/
	EvalRequest *models.MockEvalRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewEvaluateParams() beforehand.
func (o *EvaluateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.MockEvalRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("evalRequest", "body"))
			} else {
				res = append(res, errors.NewParseError("evalRequest", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.EvalRequest = &body
			}
		}
	} else {
		res = append(res, errors.Required("evalRequest", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
