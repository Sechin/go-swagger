package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/models"
)

/*AddOneCreated Created

swagger:response addOneCreated
*/
type AddOneCreated struct {

	// In: body
	Payload *models.Item `json:"body,omitempty"`
}

// NewAddOneCreated creates AddOneCreated with default headers values
func NewAddOneCreated() *AddOneCreated {
	return &AddOneCreated{}
}

// WithPayload adds the payload to the add one created response
func (o *AddOneCreated) WithPayload(payload *models.Item) *AddOneCreated {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *AddOneCreated) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddOneDefault error

swagger:response addOneDefault
*/
type AddOneDefault struct {
	_statusCode int `json:"-"`

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddOneDefault creates AddOneDefault with default headers values
func NewAddOneDefault(code int) *AddOneDefault {
	if code <= 0 {
		code = 500
	}

	return &AddOneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add one default response
func (o *AddOneDefault) WithStatusCode(code int) *AddOneDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the add one default response
func (o *AddOneDefault) WithPayload(payload *models.Error) *AddOneDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *AddOneDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
