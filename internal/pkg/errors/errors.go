package errors

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tidwall/sjson"
	"google.golang.org/grpc/status"
)

type Error interface {
	error
	HTTPStatus() int
	GRPCStatus() *status.Status
	Is(err error) bool
}

type Errors struct {
	domain      string
	title       string
	instance    string
	details     []string
	additionals map[string]interface{}
	status      *StatusCode
}

type errorJson struct {
	Type       string   `json:"type"`
	Title      string   `json:"title"`
	Instance   string   `json:"instance"`
	StatusCode int      `json:"statusCode"`
	Status     string   `json:"status"`
	Details    []string `json:"details,omitempty"`
	Detail     *string  `json:"detail,omitempty"`
}

func (e *Errors) Error() string {
	return e.domain + ": " + strings.Join(e.details, "\n\t")
}

func (e *Errors) Is(err error) bool {
	er, ok := err.(*Errors)
	if !ok {
		return false
	}
	return er.title == e.title
}

func (e *Errors) MarshalJSON() ([]byte, error) {
	m := &errorJson{
		Type:       basePathType + tosnakecase(e.title),
		Title:      e.title,
		Instance:   e.instance,
		StatusCode: e.HTTPStatus(),
		Status:     http.StatusText(e.HTTPStatus()),
	}
	switch {
	case len(e.details) > 1:
		m.Details = e.details
	case len(e.details) == 1:
		m.Detail = &e.details[0]
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	for key, val := range e.additionals {
		b, err = sjson.SetBytes(b, key, val)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}

func (e *Errors) GRPCStatus() *status.Status {
	return status.New(e.status.grpc, e.Error())
}

func (e *Errors) HTTPStatus() int {
	return e.status.http
}

type ErrorOption func(e *Errors)

var basePathType = "#"

func SetBaseURLType(base string) {
	base = strings.TrimRight(base, "/")
	basePathType = base + "#"
}

func New(status *StatusCode, domain, title, instance string, details []string, opts ...ErrorOption) error {
	err := &Errors{
		domain:      domain,
		title:       title,
		instance:    instance,
		details:     details,
		status:      status,
		additionals: make(map[string]interface{}),
	}

	for _, opt := range opts {
		opt(err)
	}
	return err
}
