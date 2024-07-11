package common

import "net/http"

type Method func(writer http.ResponseWriter, request *http.Request)

type httpHandler struct {
	get     Method
	post    Method
	put     Method
	patch   Method
	delete  Method
	head    Method
	options Method
}

func (handler *httpHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		handler.get(writer, request)
	case "POST":
		handler.post(writer, request)
	case "PUT":
		handler.put(writer, request)
	case "PATCH":
		handler.patch(writer, request)
	case "DELETE":
		handler.delete(writer, request)
	case "HEAD":
		handler.head(writer, request)
	case "OPTIONS":
		handler.options(writer, request)
	}
}

type httpHandlerBuilder struct {
	get     Method
	post    Method
	put     Method
	patch   Method
	delete  Method
	head    Method
	options Method
}

func NewHttpHandlerBuilder() *httpHandlerBuilder {
	return &httpHandlerBuilder{}
}

func (builder *httpHandlerBuilder) Get(get Method) *httpHandlerBuilder {
	builder.get = get
	return builder
}

func (builder *httpHandlerBuilder) Post(post Method) *httpHandlerBuilder {
	builder.post = post
	return builder
}

func (builder *httpHandlerBuilder) Put(put Method) *httpHandlerBuilder {
	builder.put = put
	return builder
}

func (builder *httpHandlerBuilder) Patch(patch Method) *httpHandlerBuilder {
	builder.patch = patch
	return builder
}

func (builder *httpHandlerBuilder) Delete(delete Method) *httpHandlerBuilder {
	builder.delete = delete
	return builder
}

func (builder *httpHandlerBuilder) Head(head Method) *httpHandlerBuilder {
	builder.head = head
	return builder
}

func (builder *httpHandlerBuilder) Options(options Method) *httpHandlerBuilder {
	builder.options = options
	return builder
}

func (builder *httpHandlerBuilder) Build() Method {
	handler := &httpHandler{
		get:     builder.get,
		post:    builder.post,
		put:     builder.put,
		patch:   builder.patch,
		delete:  builder.delete,
		head:    builder.head,
		options: builder.options,
	}
	return handler.Handle
}
