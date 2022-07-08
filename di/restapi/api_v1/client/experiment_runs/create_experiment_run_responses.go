// Code generated by go-swagger; DO NOT EDIT.

package experiment_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	restmodels "webank/DI/restapi/api_v1/restmodels"
)

// CreateExperimentRunReader is a Reader for the CreateExperimentRun structure.
type CreateExperimentRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateExperimentRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateExperimentRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateExperimentRunUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateExperimentRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateExperimentRunOK creates a CreateExperimentRunOK with default headers values
func NewCreateExperimentRunOK() *CreateExperimentRunOK {
	return &CreateExperimentRunOK{}
}

/*CreateExperimentRunOK handles this case with default header values.

OK
*/
type CreateExperimentRunOK struct {
	Payload *restmodels.ProphecisExperimentRun
}

func (o *CreateExperimentRunOK) Error() string {
	return fmt.Sprintf("[POST /di/v1/experimentRun/{exp_id}][%d] createExperimentRunOK  %+v", 200, o.Payload)
}

func (o *CreateExperimentRunOK) GetPayload() *restmodels.ProphecisExperimentRun {
	return o.Payload
}

func (o *CreateExperimentRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.ProphecisExperimentRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExperimentRunUnauthorized creates a CreateExperimentRunUnauthorized with default headers values
func NewCreateExperimentRunUnauthorized() *CreateExperimentRunUnauthorized {
	return &CreateExperimentRunUnauthorized{}
}

/*CreateExperimentRunUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateExperimentRunUnauthorized struct {
	Payload *restmodels.Error
}

func (o *CreateExperimentRunUnauthorized) Error() string {
	return fmt.Sprintf("[POST /di/v1/experimentRun/{exp_id}][%d] createExperimentRunUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateExperimentRunUnauthorized) GetPayload() *restmodels.Error {
	return o.Payload
}

func (o *CreateExperimentRunUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExperimentRunNotFound creates a CreateExperimentRunNotFound with default headers values
func NewCreateExperimentRunNotFound() *CreateExperimentRunNotFound {
	return &CreateExperimentRunNotFound{}
}

/*CreateExperimentRunNotFound handles this case with default header values.

The Models cannot be found
*/
type CreateExperimentRunNotFound struct {
	Payload *restmodels.Error
}

func (o *CreateExperimentRunNotFound) Error() string {
	return fmt.Sprintf("[POST /di/v1/experimentRun/{exp_id}][%d] createExperimentRunNotFound  %+v", 404, o.Payload)
}

func (o *CreateExperimentRunNotFound) GetPayload() *restmodels.Error {
	return o.Payload
}

func (o *CreateExperimentRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
