// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/barkerd427/go-vo/models"
)

// PostAPIPublicV1TeamTeamMembersReader is a Reader for the PostAPIPublicV1TeamTeamMembers structure.
type PostAPIPublicV1TeamTeamMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIPublicV1TeamTeamMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAPIPublicV1TeamTeamMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAPIPublicV1TeamTeamMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostAPIPublicV1TeamTeamMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAPIPublicV1TeamTeamMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostAPIPublicV1TeamTeamMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostAPIPublicV1TeamTeamMembersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostAPIPublicV1TeamTeamMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAPIPublicV1TeamTeamMembersOK creates a PostAPIPublicV1TeamTeamMembersOK with default headers values
func NewPostAPIPublicV1TeamTeamMembersOK() *PostAPIPublicV1TeamTeamMembersOK {
	return &PostAPIPublicV1TeamTeamMembersOK{}
}

/*PostAPIPublicV1TeamTeamMembersOK handles this case with default header values.

Some details about the team that was added
*/
type PostAPIPublicV1TeamTeamMembersOK struct {
	Payload *models.ListTeamMembersResponse
}

func (o *PostAPIPublicV1TeamTeamMembersOK) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team/{team}/members][%d] postApiPublicV1TeamTeamMembersOK  %+v", 200, o.Payload)
}

func (o *PostAPIPublicV1TeamTeamMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListTeamMembersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIPublicV1TeamTeamMembersBadRequest creates a PostAPIPublicV1TeamTeamMembersBadRequest with default headers values
func NewPostAPIPublicV1TeamTeamMembersBadRequest() *PostAPIPublicV1TeamTeamMembersBadRequest {
	return &PostAPIPublicV1TeamTeamMembersBadRequest{}
}

/*PostAPIPublicV1TeamTeamMembersBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type PostAPIPublicV1TeamTeamMembersBadRequest struct {
}

func (o *PostAPIPublicV1TeamTeamMembersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team/{team}/members][%d] postApiPublicV1TeamTeamMembersBadRequest ", 400)
}

func (o *PostAPIPublicV1TeamTeamMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1TeamTeamMembersUnauthorized creates a PostAPIPublicV1TeamTeamMembersUnauthorized with default headers values
func NewPostAPIPublicV1TeamTeamMembersUnauthorized() *PostAPIPublicV1TeamTeamMembersUnauthorized {
	return &PostAPIPublicV1TeamTeamMembersUnauthorized{}
}

/*PostAPIPublicV1TeamTeamMembersUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type PostAPIPublicV1TeamTeamMembersUnauthorized struct {
}

func (o *PostAPIPublicV1TeamTeamMembersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team/{team}/members][%d] postApiPublicV1TeamTeamMembersUnauthorized ", 401)
}

func (o *PostAPIPublicV1TeamTeamMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1TeamTeamMembersForbidden creates a PostAPIPublicV1TeamTeamMembersForbidden with default headers values
func NewPostAPIPublicV1TeamTeamMembersForbidden() *PostAPIPublicV1TeamTeamMembersForbidden {
	return &PostAPIPublicV1TeamTeamMembersForbidden{}
}

/*PostAPIPublicV1TeamTeamMembersForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type PostAPIPublicV1TeamTeamMembersForbidden struct {
}

func (o *PostAPIPublicV1TeamTeamMembersForbidden) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team/{team}/members][%d] postApiPublicV1TeamTeamMembersForbidden ", 403)
}

func (o *PostAPIPublicV1TeamTeamMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1TeamTeamMembersNotFound creates a PostAPIPublicV1TeamTeamMembersNotFound with default headers values
func NewPostAPIPublicV1TeamTeamMembersNotFound() *PostAPIPublicV1TeamTeamMembersNotFound {
	return &PostAPIPublicV1TeamTeamMembersNotFound{}
}

/*PostAPIPublicV1TeamTeamMembersNotFound handles this case with default header values.

Team not found
*/
type PostAPIPublicV1TeamTeamMembersNotFound struct {
}

func (o *PostAPIPublicV1TeamTeamMembersNotFound) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team/{team}/members][%d] postApiPublicV1TeamTeamMembersNotFound ", 404)
}

func (o *PostAPIPublicV1TeamTeamMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1TeamTeamMembersUnprocessableEntity creates a PostAPIPublicV1TeamTeamMembersUnprocessableEntity with default headers values
func NewPostAPIPublicV1TeamTeamMembersUnprocessableEntity() *PostAPIPublicV1TeamTeamMembersUnprocessableEntity {
	return &PostAPIPublicV1TeamTeamMembersUnprocessableEntity{}
}

/*PostAPIPublicV1TeamTeamMembersUnprocessableEntity handles this case with default header values.

Team name or email is unavailable, or you have reached your team limit.

*/
type PostAPIPublicV1TeamTeamMembersUnprocessableEntity struct {
}

func (o *PostAPIPublicV1TeamTeamMembersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team/{team}/members][%d] postApiPublicV1TeamTeamMembersUnprocessableEntity ", 422)
}

func (o *PostAPIPublicV1TeamTeamMembersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1TeamTeamMembersInternalServerError creates a PostAPIPublicV1TeamTeamMembersInternalServerError with default headers values
func NewPostAPIPublicV1TeamTeamMembersInternalServerError() *PostAPIPublicV1TeamTeamMembersInternalServerError {
	return &PostAPIPublicV1TeamTeamMembersInternalServerError{}
}

/*PostAPIPublicV1TeamTeamMembersInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostAPIPublicV1TeamTeamMembersInternalServerError struct {
}

func (o *PostAPIPublicV1TeamTeamMembersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team/{team}/members][%d] postApiPublicV1TeamTeamMembersInternalServerError ", 500)
}

func (o *PostAPIPublicV1TeamTeamMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
