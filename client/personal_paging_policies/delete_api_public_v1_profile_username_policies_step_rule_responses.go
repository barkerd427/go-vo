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

// DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleReader is a Reader for the DeleteAPIPublicV1ProfileUsernamePoliciesStepRule structure.
type DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleOK creates a DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleOK with default headers values
func NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleOK() *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleOK {
	return &DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleOK{}
}

/*DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleOK handles this case with default header values.

The deleted rule from the paging policy step
*/
type DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleOK struct {
	Payload *models.InlineResponse2005
}

func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleOK) Error() string {
	return fmt.Sprintf("[DELETE /api-public/v1/profile/{username}/policies/{step}/{rule}][%d] deleteApiPublicV1ProfileUsernamePoliciesStepRuleOK  %+v", 200, o.Payload)
}

func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2005)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest creates a DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest with default headers values
func NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest() *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest {
	return &DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest{}
}

/*DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest struct {
}

func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api-public/v1/profile/{username}/policies/{step}/{rule}][%d] deleteApiPublicV1ProfileUsernamePoliciesStepRuleBadRequest ", 400)
}

func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized creates a DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized with default headers values
func NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized() *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized {
	return &DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized{}
}

/*DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized struct {
}

func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api-public/v1/profile/{username}/policies/{step}/{rule}][%d] deleteApiPublicV1ProfileUsernamePoliciesStepRuleUnauthorized ", 401)
}

func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden creates a DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden with default headers values
func NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden() *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden {
	return &DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden{}
}

/*DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden struct {
}

func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api-public/v1/profile/{username}/policies/{step}/{rule}][%d] deleteApiPublicV1ProfileUsernamePoliciesStepRuleForbidden ", 403)
}

func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError creates a DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError with default headers values
func NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError() *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError {
	return &DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError{}
}

/*DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError struct {
}

func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api-public/v1/profile/{username}/policies/{step}/{rule}][%d] deleteApiPublicV1ProfileUsernamePoliciesStepRuleInternalServerError ", 500)
}

func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
