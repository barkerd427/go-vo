// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/barkerd427/go-vo/models"
)

// PostAPIPublicV1IncidentsReader is a Reader for the PostAPIPublicV1Incidents structure.
type PostAPIPublicV1IncidentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIPublicV1IncidentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAPIPublicV1IncidentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAPIPublicV1IncidentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostAPIPublicV1IncidentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAPIPublicV1IncidentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostAPIPublicV1IncidentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostAPIPublicV1IncidentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAPIPublicV1IncidentsOK creates a PostAPIPublicV1IncidentsOK with default headers values
func NewPostAPIPublicV1IncidentsOK() *PostAPIPublicV1IncidentsOK {
	return &PostAPIPublicV1IncidentsOK{}
}

/*PostAPIPublicV1IncidentsOK handles this case with default header values.

Information about the incident created

*/
type PostAPIPublicV1IncidentsOK struct {
	Payload *models.CreatedIncident
}

func (o *PostAPIPublicV1IncidentsOK) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/incidents][%d] postApiPublicV1IncidentsOK  %+v", 200, o.Payload)
}

func (o *PostAPIPublicV1IncidentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreatedIncident)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIPublicV1IncidentsBadRequest creates a PostAPIPublicV1IncidentsBadRequest with default headers values
func NewPostAPIPublicV1IncidentsBadRequest() *PostAPIPublicV1IncidentsBadRequest {
	return &PostAPIPublicV1IncidentsBadRequest{}
}

/*PostAPIPublicV1IncidentsBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type PostAPIPublicV1IncidentsBadRequest struct {
}

func (o *PostAPIPublicV1IncidentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/incidents][%d] postApiPublicV1IncidentsBadRequest ", 400)
}

func (o *PostAPIPublicV1IncidentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1IncidentsUnauthorized creates a PostAPIPublicV1IncidentsUnauthorized with default headers values
func NewPostAPIPublicV1IncidentsUnauthorized() *PostAPIPublicV1IncidentsUnauthorized {
	return &PostAPIPublicV1IncidentsUnauthorized{}
}

/*PostAPIPublicV1IncidentsUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type PostAPIPublicV1IncidentsUnauthorized struct {
}

func (o *PostAPIPublicV1IncidentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/incidents][%d] postApiPublicV1IncidentsUnauthorized ", 401)
}

func (o *PostAPIPublicV1IncidentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1IncidentsForbidden creates a PostAPIPublicV1IncidentsForbidden with default headers values
func NewPostAPIPublicV1IncidentsForbidden() *PostAPIPublicV1IncidentsForbidden {
	return &PostAPIPublicV1IncidentsForbidden{}
}

/*PostAPIPublicV1IncidentsForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type PostAPIPublicV1IncidentsForbidden struct {
}

func (o *PostAPIPublicV1IncidentsForbidden) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/incidents][%d] postApiPublicV1IncidentsForbidden ", 403)
}

func (o *PostAPIPublicV1IncidentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1IncidentsNotFound creates a PostAPIPublicV1IncidentsNotFound with default headers values
func NewPostAPIPublicV1IncidentsNotFound() *PostAPIPublicV1IncidentsNotFound {
	return &PostAPIPublicV1IncidentsNotFound{}
}

/*PostAPIPublicV1IncidentsNotFound handles this case with default header values.

Path not found
*/
type PostAPIPublicV1IncidentsNotFound struct {
}

func (o *PostAPIPublicV1IncidentsNotFound) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/incidents][%d] postApiPublicV1IncidentsNotFound ", 404)
}

func (o *PostAPIPublicV1IncidentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1IncidentsInternalServerError creates a PostAPIPublicV1IncidentsInternalServerError with default headers values
func NewPostAPIPublicV1IncidentsInternalServerError() *PostAPIPublicV1IncidentsInternalServerError {
	return &PostAPIPublicV1IncidentsInternalServerError{}
}

/*PostAPIPublicV1IncidentsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostAPIPublicV1IncidentsInternalServerError struct {
}

func (o *PostAPIPublicV1IncidentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/incidents][%d] postApiPublicV1IncidentsInternalServerError ", 500)
}

func (o *PostAPIPublicV1IncidentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
