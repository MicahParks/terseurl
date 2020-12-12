// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/MicahParks/terse-URL/models"
)

// NewTerseURLAPI creates a new TerseURL instance
func NewTerseURLAPI(spec *loads.Document) *TerseURLAPI {
	return &TerseURLAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		AliveHandler: AliveHandlerFunc(func(params AliveParams) middleware.Responder {
			return middleware.NotImplemented("operation Alive has not yet been implemented")
		}),
		URLDeleteHandler: URLDeleteHandlerFunc(func(params URLDeleteParams, principal *models.JWTInfo) middleware.Responder {
			return middleware.NotImplemented("operation URLDelete has not yet been implemented")
		}),
		URLDumpHandler: URLDumpHandlerFunc(func(params URLDumpParams, principal *models.JWTInfo) middleware.Responder {
			return middleware.NotImplemented("operation URLDump has not yet been implemented")
		}),
		URLDumpShortenedHandler: URLDumpShortenedHandlerFunc(func(params URLDumpShortenedParams, principal *models.JWTInfo) middleware.Responder {
			return middleware.NotImplemented("operation URLDumpShortened has not yet been implemented")
		}),
		URLGetHandler: URLGetHandlerFunc(func(params URLGetParams) middleware.Responder {
			return middleware.NotImplemented("operation URLGet has not yet been implemented")
		}),
		URLNewHandler: URLNewHandlerFunc(func(params URLNewParams, principal *models.JWTInfo) middleware.Responder {
			return middleware.NotImplemented("operation URLNew has not yet been implemented")
		}),
		URLTrackHandler: URLTrackHandlerFunc(func(params URLTrackParams, principal *models.JWTInfo) middleware.Responder {
			return middleware.NotImplemented("operation URLTrack has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		JWTAuth: func(token string) (*models.JWTInfo, error) {
			return nil, errors.NotImplemented("api key auth (JWT) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*TerseURLAPI the terse URL API */
type TerseURLAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// JWTAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	JWTAuth func(string) (*models.JWTInfo, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// AliveHandler sets the operation handler for the alive operation
	AliveHandler AliveHandler
	// URLDeleteHandler sets the operation handler for the url delete operation
	URLDeleteHandler URLDeleteHandler
	// URLDumpHandler sets the operation handler for the url dump operation
	URLDumpHandler URLDumpHandler
	// URLDumpShortenedHandler sets the operation handler for the url dump shortened operation
	URLDumpShortenedHandler URLDumpShortenedHandler
	// URLGetHandler sets the operation handler for the url get operation
	URLGetHandler URLGetHandler
	// URLNewHandler sets the operation handler for the url new operation
	URLNewHandler URLNewHandler
	// URLTrackHandler sets the operation handler for the url track operation
	URLTrackHandler URLTrackHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *TerseURLAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *TerseURLAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *TerseURLAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TerseURLAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TerseURLAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TerseURLAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TerseURLAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TerseURLAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TerseURLAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TerseURLAPI
func (o *TerseURLAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.JWTAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.AliveHandler == nil {
		unregistered = append(unregistered, "AliveHandler")
	}
	if o.URLDeleteHandler == nil {
		unregistered = append(unregistered, "URLDeleteHandler")
	}
	if o.URLDumpHandler == nil {
		unregistered = append(unregistered, "URLDumpHandler")
	}
	if o.URLDumpShortenedHandler == nil {
		unregistered = append(unregistered, "URLDumpShortenedHandler")
	}
	if o.URLGetHandler == nil {
		unregistered = append(unregistered, "URLGetHandler")
	}
	if o.URLNewHandler == nil {
		unregistered = append(unregistered, "URLNewHandler")
	}
	if o.URLTrackHandler == nil {
		unregistered = append(unregistered, "URLTrackHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TerseURLAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TerseURLAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "JWT":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.JWTAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *TerseURLAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *TerseURLAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *TerseURLAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TerseURLAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the terse URL API
func (o *TerseURLAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TerseURLAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/alive"] = NewAlive(o.context, o.AliveHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/api/delete"] = NewURLDelete(o.context, o.URLDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/dump"] = NewURLDump(o.context, o.URLDumpHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/dump/{shortened}"] = NewURLDumpShortened(o.context, o.URLDumpShortenedHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/{shortened}"] = NewURLGet(o.context, o.URLGetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/new"] = NewURLNew(o.context, o.URLNewHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/track/{shortened}"] = NewURLTrack(o.context, o.URLTrackHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TerseURLAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TerseURLAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TerseURLAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TerseURLAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *TerseURLAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
