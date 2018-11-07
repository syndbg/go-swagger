// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/syndbg/go-swagger/examples/task-tracker/models"
)

// UploadTaskFileReader is a Reader for the UploadTaskFile structure.
type UploadTaskFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadTaskFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUploadTaskFileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUploadTaskFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadTaskFileCreated creates a UploadTaskFileCreated with default headers values
func NewUploadTaskFileCreated() *UploadTaskFileCreated {
	return &UploadTaskFileCreated{}
}

/*UploadTaskFileCreated handles this case with default header values.

File added
*/
type UploadTaskFileCreated struct {
}

func (o *UploadTaskFileCreated) Error() string {
	return fmt.Sprintf("[POST /tasks/{id}/files][%d] uploadTaskFileCreated ", 201)
}

func (o *UploadTaskFileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadTaskFileDefault creates a UploadTaskFileDefault with default headers values
func NewUploadTaskFileDefault(code int) *UploadTaskFileDefault {
	return &UploadTaskFileDefault{
		_statusCode: code,
	}
}

/*UploadTaskFileDefault handles this case with default header values.

Error response
*/
type UploadTaskFileDefault struct {
	_statusCode int

	XErrorCode string

	Payload *models.Error
}

// Code gets the status code for the upload task file default response
func (o *UploadTaskFileDefault) Code() int {
	return o._statusCode
}

func (o *UploadTaskFileDefault) Error() string {
	return fmt.Sprintf("[POST /tasks/{id}/files][%d] uploadTaskFile default  %+v", o._statusCode, o.Payload)
}

func (o *UploadTaskFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Error-Code
	o.XErrorCode = response.GetHeader("X-Error-Code")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
