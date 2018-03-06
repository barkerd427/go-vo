// Code generated by go-swagger; DO NOT EDIT.

package personal_paging_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/barkerd427/go-vo/models"
)

// PutAPIPublicV1ProfileUsernamePoliciesStepReader is a Reader for the PutAPIPublicV1ProfileUsernamePoliciesStep structure.
type PutAPIPublicV1ProfileUsernamePoliciesStepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAPIPublicV1ProfileUsernamePoliciesStepOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutAPIPublicV1ProfileUsernamePoliciesStepBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPutAPIPublicV1ProfileUsernamePoliciesStepUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutAPIPublicV1ProfileUsernamePoliciesStepForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutAPIPublicV1ProfileUsernamePoliciesStepInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAPIPublicV1ProfileUsernamePoliciesStepOK creates a PutAPIPublicV1ProfileUsernamePoliciesStepOK with default headers values
func NewPutAPIPublicV1ProfileUsernamePoliciesStepOK() *PutAPIPublicV1ProfileUsernamePoliciesStepOK {
	return &PutAPIPublicV1ProfileUsernamePoliciesStepOK{}
}

/*PutAPIPublicV1ProfileUsernamePoliciesStepOK handles this case with default header values.

The updated paging policy step
*/
type PutAPIPublicV1ProfileUsernamePoliciesStepOK struct {
	Payload *models.InlineResponse2004
}

func (o *PutAPIPublicV1ProfileUsernamePoliciesStepOK) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/profile/{username}/policies/{step}][%d] putApiPublicV1ProfileUsernamePoliciesStepOK  %+v", 200, o.Payload)
}

func (o *PutAPIPublicV1ProfileUsernamePoliciesStepOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2004)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIPublicV1ProfileUsernamePoliciesStepBadRequest creates a PutAPIPublicV1ProfileUsernamePoliciesStepBadRequest with default headers values
func NewPutAPIPublicV1ProfileUsernamePoliciesStepBadRequest() *PutAPIPublicV1ProfileUsernamePoliciesStepBadRequest {
	return &PutAPIPublicV1ProfileUsernamePoliciesStepBadRequest{}
}

/*PutAPIPublicV1ProfileUsernamePoliciesStepBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type PutAPIPublicV1ProfileUsernamePoliciesStepBadRequest struct {
}

func (o *PutAPIPublicV1ProfileUsernamePoliciesStepBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/profile/{username}/policies/{step}][%d] putApiPublicV1ProfileUsernamePoliciesStepBadRequest ", 400)
}

func (o *PutAPIPublicV1ProfileUsernamePoliciesStepBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIPublicV1ProfileUsernamePoliciesStepUnauthorized creates a PutAPIPublicV1ProfileUsernamePoliciesStepUnauthorized with default headers values
func NewPutAPIPublicV1ProfileUsernamePoliciesStepUnauthorized() *PutAPIPublicV1ProfileUsernamePoliciesStepUnauthorized {
	return &PutAPIPublicV1ProfileUsernamePoliciesStepUnauthorized{}
}

/*PutAPIPublicV1ProfileUsernamePoliciesStepUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type PutAPIPublicV1ProfileUsernamePoliciesStepUnauthorized struct {
}

func (o *PutAPIPublicV1ProfileUsernamePoliciesStepUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/profile/{username}/policies/{step}][%d] putApiPublicV1ProfileUsernamePoliciesStepUnauthorized ", 401)
}

func (o *PutAPIPublicV1ProfileUsernamePoliciesStepUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIPublicV1ProfileUsernamePoliciesStepForbidden creates a PutAPIPublicV1ProfileUsernamePoliciesStepForbidden with default headers values
func NewPutAPIPublicV1ProfileUsernamePoliciesStepForbidden() *PutAPIPublicV1ProfileUsernamePoliciesStepForbidden {
	return &PutAPIPublicV1ProfileUsernamePoliciesStepForbidden{}
}

/*PutAPIPublicV1ProfileUsernamePoliciesStepForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type PutAPIPublicV1ProfileUsernamePoliciesStepForbidden struct {
}

func (o *PutAPIPublicV1ProfileUsernamePoliciesStepForbidden) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/profile/{username}/policies/{step}][%d] putApiPublicV1ProfileUsernamePoliciesStepForbidden ", 403)
}

func (o *PutAPIPublicV1ProfileUsernamePoliciesStepForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIPublicV1ProfileUsernamePoliciesStepInternalServerError creates a PutAPIPublicV1ProfileUsernamePoliciesStepInternalServerError with default headers values
func NewPutAPIPublicV1ProfileUsernamePoliciesStepInternalServerError() *PutAPIPublicV1ProfileUsernamePoliciesStepInternalServerError {
	return &PutAPIPublicV1ProfileUsernamePoliciesStepInternalServerError{}
}

/*PutAPIPublicV1ProfileUsernamePoliciesStepInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutAPIPublicV1ProfileUsernamePoliciesStepInternalServerError struct {
}

func (o *PutAPIPublicV1ProfileUsernamePoliciesStepInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/profile/{username}/policies/{step}][%d] putApiPublicV1ProfileUsernamePoliciesStepInternalServerError ", 500)
}

func (o *PutAPIPublicV1ProfileUsernamePoliciesStepInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
