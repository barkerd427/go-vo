// Code generated by go-swagger; DO NOT EDIT.

package personal_paging_policy_values

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/barkerd427/go-vo/models"
)

// GetAPIPublicV1PoliciesTypesTimeoutsReader is a Reader for the GetAPIPublicV1PoliciesTypesTimeouts structure.
type GetAPIPublicV1PoliciesTypesTimeoutsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIPublicV1PoliciesTypesTimeoutsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIPublicV1PoliciesTypesTimeoutsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAPIPublicV1PoliciesTypesTimeoutsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAPIPublicV1PoliciesTypesTimeoutsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIPublicV1PoliciesTypesTimeoutsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIPublicV1PoliciesTypesTimeoutsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIPublicV1PoliciesTypesTimeoutsOK creates a GetAPIPublicV1PoliciesTypesTimeoutsOK with default headers values
func NewGetAPIPublicV1PoliciesTypesTimeoutsOK() *GetAPIPublicV1PoliciesTypesTimeoutsOK {
	return &GetAPIPublicV1PoliciesTypesTimeoutsOK{}
}

/*GetAPIPublicV1PoliciesTypesTimeoutsOK handles this case with default header values.

All the available timeout types.
*/
type GetAPIPublicV1PoliciesTypesTimeoutsOK struct {
	Payload *models.InlineResponse2002
}

func (o *GetAPIPublicV1PoliciesTypesTimeoutsOK) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/policies/types/timeouts][%d] getApiPublicV1PoliciesTypesTimeoutsOK  %+v", 200, o.Payload)
}

func (o *GetAPIPublicV1PoliciesTypesTimeoutsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2002)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIPublicV1PoliciesTypesTimeoutsBadRequest creates a GetAPIPublicV1PoliciesTypesTimeoutsBadRequest with default headers values
func NewGetAPIPublicV1PoliciesTypesTimeoutsBadRequest() *GetAPIPublicV1PoliciesTypesTimeoutsBadRequest {
	return &GetAPIPublicV1PoliciesTypesTimeoutsBadRequest{}
}

/*GetAPIPublicV1PoliciesTypesTimeoutsBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type GetAPIPublicV1PoliciesTypesTimeoutsBadRequest struct {
}

func (o *GetAPIPublicV1PoliciesTypesTimeoutsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/policies/types/timeouts][%d] getApiPublicV1PoliciesTypesTimeoutsBadRequest ", 400)
}

func (o *GetAPIPublicV1PoliciesTypesTimeoutsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1PoliciesTypesTimeoutsUnauthorized creates a GetAPIPublicV1PoliciesTypesTimeoutsUnauthorized with default headers values
func NewGetAPIPublicV1PoliciesTypesTimeoutsUnauthorized() *GetAPIPublicV1PoliciesTypesTimeoutsUnauthorized {
	return &GetAPIPublicV1PoliciesTypesTimeoutsUnauthorized{}
}

/*GetAPIPublicV1PoliciesTypesTimeoutsUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type GetAPIPublicV1PoliciesTypesTimeoutsUnauthorized struct {
}

func (o *GetAPIPublicV1PoliciesTypesTimeoutsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/policies/types/timeouts][%d] getApiPublicV1PoliciesTypesTimeoutsUnauthorized ", 401)
}

func (o *GetAPIPublicV1PoliciesTypesTimeoutsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1PoliciesTypesTimeoutsForbidden creates a GetAPIPublicV1PoliciesTypesTimeoutsForbidden with default headers values
func NewGetAPIPublicV1PoliciesTypesTimeoutsForbidden() *GetAPIPublicV1PoliciesTypesTimeoutsForbidden {
	return &GetAPIPublicV1PoliciesTypesTimeoutsForbidden{}
}

/*GetAPIPublicV1PoliciesTypesTimeoutsForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type GetAPIPublicV1PoliciesTypesTimeoutsForbidden struct {
}

func (o *GetAPIPublicV1PoliciesTypesTimeoutsForbidden) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/policies/types/timeouts][%d] getApiPublicV1PoliciesTypesTimeoutsForbidden ", 403)
}

func (o *GetAPIPublicV1PoliciesTypesTimeoutsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1PoliciesTypesTimeoutsInternalServerError creates a GetAPIPublicV1PoliciesTypesTimeoutsInternalServerError with default headers values
func NewGetAPIPublicV1PoliciesTypesTimeoutsInternalServerError() *GetAPIPublicV1PoliciesTypesTimeoutsInternalServerError {
	return &GetAPIPublicV1PoliciesTypesTimeoutsInternalServerError{}
}

/*GetAPIPublicV1PoliciesTypesTimeoutsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAPIPublicV1PoliciesTypesTimeoutsInternalServerError struct {
}

func (o *GetAPIPublicV1PoliciesTypesTimeoutsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/policies/types/timeouts][%d] getApiPublicV1PoliciesTypesTimeoutsInternalServerError ", 500)
}

func (o *GetAPIPublicV1PoliciesTypesTimeoutsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
