// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.1-0.20240604070534-2f0ff757704b DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List a redirection link
	// (GET /links)
	LinkList(w http.ResponseWriter, r *http.Request)
	// Create a redirection link
	// (POST /links)
	LinkCreate(w http.ResponseWriter, r *http.Request)
	// Delete a redirection link
	// (DELETE /links/{linkId})
	LinkDelete(w http.ResponseWriter, r *http.Request, linkId LinkId)
	// Get a redirection link
	// (GET /links/{linkId})
	LinkGet(w http.ResponseWriter, r *http.Request, linkId LinkId)
	// Redirect
	// (GET /{hostname}/{slug})
	LinkRedirect(w http.ResponseWriter, r *http.Request, hostname Hostname, slug Slug)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// LinkList operation middleware
func (siw *ServerInterfaceWrapper) LinkList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LinkList(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LinkCreate operation middleware
func (siw *ServerInterfaceWrapper) LinkCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LinkCreate(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LinkDelete operation middleware
func (siw *ServerInterfaceWrapper) LinkDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "linkId" -------------
	var linkId LinkId

	err = runtime.BindStyledParameterWithOptions("simple", "linkId", mux.Vars(r)["linkId"], &linkId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "linkId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LinkDelete(w, r, linkId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LinkGet operation middleware
func (siw *ServerInterfaceWrapper) LinkGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "linkId" -------------
	var linkId LinkId

	err = runtime.BindStyledParameterWithOptions("simple", "linkId", mux.Vars(r)["linkId"], &linkId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "linkId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LinkGet(w, r, linkId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LinkRedirect operation middleware
func (siw *ServerInterfaceWrapper) LinkRedirect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "hostname" -------------
	var hostname Hostname

	err = runtime.BindStyledParameterWithOptions("simple", "hostname", mux.Vars(r)["hostname"], &hostname, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "hostname", Err: err})
		return
	}

	// ------------- Path parameter "slug" -------------
	var slug Slug

	err = runtime.BindStyledParameterWithOptions("simple", "slug", mux.Vars(r)["slug"], &slug, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "slug", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LinkRedirect(w, r, hostname, slug)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/links", wrapper.LinkList).Methods("GET")

	r.HandleFunc(options.BaseURL+"/links", wrapper.LinkCreate).Methods("POST")

	r.HandleFunc(options.BaseURL+"/links/{linkId}", wrapper.LinkDelete).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/links/{linkId}", wrapper.LinkGet).Methods("GET")

	r.HandleFunc(options.BaseURL+"/{hostname}/{slug}", wrapper.LinkRedirect).Methods("GET")

	return r
}

type ErrorBadRequestResponseApplicationProblemPlusJSONResponse Error

type ErrorConflictResponseApplicationProblemPlusJSONResponse Error

type ErrorForbiddenResponseApplicationProblemPlusJSONResponse Error

type ErrorInternalServerResponseApplicationProblemPlusJSONResponse Error

type ErrorNotFoundResponseApplicationProblemPlusJSONResponse Error

type ErrorUnauthorizedResponseApplicationProblemPlusJSONResponse Error

type LinkCreatedResponseJSONResponse Link

type LinkDeletedResponseResponse struct {
}

type LinkGetResponseJSONResponse Link

type LinkListResponseJSONResponse []Link

type LinkRedirectedResponseResponseHeaders struct {
	Location string
}
type LinkRedirectedResponseResponse struct {
	Headers LinkRedirectedResponseResponseHeaders
}

type LinkListRequestObject struct {
}

type LinkListResponseObject interface {
	VisitLinkListResponse(w http.ResponseWriter) error
}

type LinkList200JSONResponse struct{ LinkListResponseJSONResponse }

func (response LinkList200JSONResponse) VisitLinkListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type LinkList400ApplicationProblemPlusJSONResponse struct {
	ErrorBadRequestResponseApplicationProblemPlusJSONResponse
}

func (response LinkList400ApplicationProblemPlusJSONResponse) VisitLinkListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type LinkList401ApplicationProblemPlusJSONResponse struct {
	ErrorUnauthorizedResponseApplicationProblemPlusJSONResponse
}

func (response LinkList401ApplicationProblemPlusJSONResponse) VisitLinkListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type LinkList403ApplicationProblemPlusJSONResponse struct {
	ErrorForbiddenResponseApplicationProblemPlusJSONResponse
}

func (response LinkList403ApplicationProblemPlusJSONResponse) VisitLinkListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type LinkList409ApplicationProblemPlusJSONResponse struct {
	ErrorConflictResponseApplicationProblemPlusJSONResponse
}

func (response LinkList409ApplicationProblemPlusJSONResponse) VisitLinkListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type LinkList500ApplicationProblemPlusJSONResponse struct {
	ErrorInternalServerResponseApplicationProblemPlusJSONResponse
}

func (response LinkList500ApplicationProblemPlusJSONResponse) VisitLinkListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type LinkCreateRequestObject struct {
	Body *LinkCreateJSONRequestBody
}

type LinkCreateResponseObject interface {
	VisitLinkCreateResponse(w http.ResponseWriter) error
}

type LinkCreate201JSONResponse struct {
	LinkCreatedResponseJSONResponse
}

func (response LinkCreate201JSONResponse) VisitLinkCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type LinkCreate400ApplicationProblemPlusJSONResponse struct {
	ErrorBadRequestResponseApplicationProblemPlusJSONResponse
}

func (response LinkCreate400ApplicationProblemPlusJSONResponse) VisitLinkCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type LinkCreate401ApplicationProblemPlusJSONResponse struct {
	ErrorUnauthorizedResponseApplicationProblemPlusJSONResponse
}

func (response LinkCreate401ApplicationProblemPlusJSONResponse) VisitLinkCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type LinkCreate403ApplicationProblemPlusJSONResponse struct {
	ErrorForbiddenResponseApplicationProblemPlusJSONResponse
}

func (response LinkCreate403ApplicationProblemPlusJSONResponse) VisitLinkCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type LinkCreate409ApplicationProblemPlusJSONResponse struct {
	ErrorConflictResponseApplicationProblemPlusJSONResponse
}

func (response LinkCreate409ApplicationProblemPlusJSONResponse) VisitLinkCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type LinkCreate500ApplicationProblemPlusJSONResponse struct {
	ErrorInternalServerResponseApplicationProblemPlusJSONResponse
}

func (response LinkCreate500ApplicationProblemPlusJSONResponse) VisitLinkCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type LinkDeleteRequestObject struct {
	LinkId LinkId `json:"linkId"`
}

type LinkDeleteResponseObject interface {
	VisitLinkDeleteResponse(w http.ResponseWriter) error
}

type LinkDelete200Response = LinkDeletedResponseResponse

func (response LinkDelete200Response) VisitLinkDeleteResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type LinkDelete400ApplicationProblemPlusJSONResponse struct {
	ErrorBadRequestResponseApplicationProblemPlusJSONResponse
}

func (response LinkDelete400ApplicationProblemPlusJSONResponse) VisitLinkDeleteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type LinkDelete401ApplicationProblemPlusJSONResponse struct {
	ErrorUnauthorizedResponseApplicationProblemPlusJSONResponse
}

func (response LinkDelete401ApplicationProblemPlusJSONResponse) VisitLinkDeleteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type LinkDelete403ApplicationProblemPlusJSONResponse struct {
	ErrorForbiddenResponseApplicationProblemPlusJSONResponse
}

func (response LinkDelete403ApplicationProblemPlusJSONResponse) VisitLinkDeleteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type LinkDelete404ApplicationProblemPlusJSONResponse struct {
	ErrorNotFoundResponseApplicationProblemPlusJSONResponse
}

func (response LinkDelete404ApplicationProblemPlusJSONResponse) VisitLinkDeleteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type LinkDelete409ApplicationProblemPlusJSONResponse struct {
	ErrorConflictResponseApplicationProblemPlusJSONResponse
}

func (response LinkDelete409ApplicationProblemPlusJSONResponse) VisitLinkDeleteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type LinkDelete500ApplicationProblemPlusJSONResponse struct {
	ErrorInternalServerResponseApplicationProblemPlusJSONResponse
}

func (response LinkDelete500ApplicationProblemPlusJSONResponse) VisitLinkDeleteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type LinkGetRequestObject struct {
	LinkId LinkId `json:"linkId"`
}

type LinkGetResponseObject interface {
	VisitLinkGetResponse(w http.ResponseWriter) error
}

type LinkGet200JSONResponse struct{ LinkGetResponseJSONResponse }

func (response LinkGet200JSONResponse) VisitLinkGetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type LinkGet400ApplicationProblemPlusJSONResponse struct {
	ErrorBadRequestResponseApplicationProblemPlusJSONResponse
}

func (response LinkGet400ApplicationProblemPlusJSONResponse) VisitLinkGetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type LinkGet401ApplicationProblemPlusJSONResponse struct {
	ErrorUnauthorizedResponseApplicationProblemPlusJSONResponse
}

func (response LinkGet401ApplicationProblemPlusJSONResponse) VisitLinkGetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type LinkGet403ApplicationProblemPlusJSONResponse struct {
	ErrorForbiddenResponseApplicationProblemPlusJSONResponse
}

func (response LinkGet403ApplicationProblemPlusJSONResponse) VisitLinkGetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type LinkGet404ApplicationProblemPlusJSONResponse struct {
	ErrorNotFoundResponseApplicationProblemPlusJSONResponse
}

func (response LinkGet404ApplicationProblemPlusJSONResponse) VisitLinkGetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type LinkGet500ApplicationProblemPlusJSONResponse struct {
	ErrorInternalServerResponseApplicationProblemPlusJSONResponse
}

func (response LinkGet500ApplicationProblemPlusJSONResponse) VisitLinkGetResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type LinkRedirectRequestObject struct {
	Hostname Hostname `json:"hostname"`
	Slug     Slug     `json:"slug"`
}

type LinkRedirectResponseObject interface {
	VisitLinkRedirectResponse(w http.ResponseWriter) error
}

type LinkRedirect301Response = LinkRedirectedResponseResponse

func (response LinkRedirect301Response) VisitLinkRedirectResponse(w http.ResponseWriter) error {
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(301)
	return nil
}

type LinkRedirect400ApplicationProblemPlusJSONResponse struct {
	ErrorBadRequestResponseApplicationProblemPlusJSONResponse
}

func (response LinkRedirect400ApplicationProblemPlusJSONResponse) VisitLinkRedirectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type LinkRedirect401ApplicationProblemPlusJSONResponse struct {
	ErrorUnauthorizedResponseApplicationProblemPlusJSONResponse
}

func (response LinkRedirect401ApplicationProblemPlusJSONResponse) VisitLinkRedirectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type LinkRedirect403ApplicationProblemPlusJSONResponse struct {
	ErrorForbiddenResponseApplicationProblemPlusJSONResponse
}

func (response LinkRedirect403ApplicationProblemPlusJSONResponse) VisitLinkRedirectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type LinkRedirect409ApplicationProblemPlusJSONResponse struct {
	ErrorConflictResponseApplicationProblemPlusJSONResponse
}

func (response LinkRedirect409ApplicationProblemPlusJSONResponse) VisitLinkRedirectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type LinkRedirect500ApplicationProblemPlusJSONResponse struct {
	ErrorInternalServerResponseApplicationProblemPlusJSONResponse
}

func (response LinkRedirect500ApplicationProblemPlusJSONResponse) VisitLinkRedirectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// List a redirection link
	// (GET /links)
	LinkList(ctx context.Context, request LinkListRequestObject) (LinkListResponseObject, error)
	// Create a redirection link
	// (POST /links)
	LinkCreate(ctx context.Context, request LinkCreateRequestObject) (LinkCreateResponseObject, error)
	// Delete a redirection link
	// (DELETE /links/{linkId})
	LinkDelete(ctx context.Context, request LinkDeleteRequestObject) (LinkDeleteResponseObject, error)
	// Get a redirection link
	// (GET /links/{linkId})
	LinkGet(ctx context.Context, request LinkGetRequestObject) (LinkGetResponseObject, error)
	// Redirect
	// (GET /{hostname}/{slug})
	LinkRedirect(ctx context.Context, request LinkRedirectRequestObject) (LinkRedirectResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// LinkList operation middleware
func (sh *strictHandler) LinkList(w http.ResponseWriter, r *http.Request) {
	var request LinkListRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.LinkList(ctx, request.(LinkListRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "LinkList")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(LinkListResponseObject); ok {
		if err := validResponse.VisitLinkListResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// LinkCreate operation middleware
func (sh *strictHandler) LinkCreate(w http.ResponseWriter, r *http.Request) {
	var request LinkCreateRequestObject

	var body LinkCreateJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.LinkCreate(ctx, request.(LinkCreateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "LinkCreate")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(LinkCreateResponseObject); ok {
		if err := validResponse.VisitLinkCreateResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// LinkDelete operation middleware
func (sh *strictHandler) LinkDelete(w http.ResponseWriter, r *http.Request, linkId LinkId) {
	var request LinkDeleteRequestObject

	request.LinkId = linkId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.LinkDelete(ctx, request.(LinkDeleteRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "LinkDelete")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(LinkDeleteResponseObject); ok {
		if err := validResponse.VisitLinkDeleteResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// LinkGet operation middleware
func (sh *strictHandler) LinkGet(w http.ResponseWriter, r *http.Request, linkId LinkId) {
	var request LinkGetRequestObject

	request.LinkId = linkId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.LinkGet(ctx, request.(LinkGetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "LinkGet")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(LinkGetResponseObject); ok {
		if err := validResponse.VisitLinkGetResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// LinkRedirect operation middleware
func (sh *strictHandler) LinkRedirect(w http.ResponseWriter, r *http.Request, hostname Hostname, slug Slug) {
	var request LinkRedirectRequestObject

	request.Hostname = hostname
	request.Slug = slug

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.LinkRedirect(ctx, request.(LinkRedirectRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "LinkRedirect")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(LinkRedirectResponseObject); ok {
		if err := validResponse.VisitLinkRedirectResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"tzNfU+0jq3Ggg7H+2vbvAAAA//+NgxYjEx0AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
