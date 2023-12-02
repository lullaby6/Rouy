package rouy

import (
	"net/http"
	"net/url"
)

type Context struct {
	Query    url.Values
	Request  *http.Request
	Response http.ResponseWriter
	Method   string
	Path     string
	Body     interface{}
}

func (context Context) StatusCode(statusCode int) {
	context.Response.WriteHeader(statusCode)
}

func (context Context) Status(statusCode int) {
	context.Response.WriteHeader(statusCode)
}

func (context Context) Header(key string, value string) {
	context.Response.Header().Set(key, value)
}

func (context Context) Write(data []byte) {
	context.Response.Write(data)
}

func (context Context) Send(data []byte) {
	context.Response.Write(data)
}

func (context Context) ContentType(contentType string) {
	context.Response.Header().Set("Content-Type", contentType)
}

func (context Context) Type(contentType string) {
	context.Response.Header().Set("Content-Type", contentType)
}

func (context Context) JSON(StatusCode int, body interface{}) *Response {
	return &Response{
		ContentType: "application/json",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) Text(StatusCode int, body string) *Response {
	return &Response{
		ContentType: "text/plain",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) Image(StatusCode int, body []byte) *Response {
	return &Response{
		ContentType: "image/png",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) PDF(StatusCode int, body []byte) *Response {
	return &Response{
		ContentType: "application/pdf",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) HTML(StatusCode int, body string) *Response {
	return &Response{
		ContentType: "text/html",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) ZIP(StatusCode int, body []byte) *Response {
	return &Response{
		ContentType: "application/zip",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) XML(StatusCode int, body string) *Response {
	return &Response{
		ContentType: "application/xml",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) JS(StatusCode int, body string) *Response {
	return &Response{
		ContentType: "application/javascript",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) CSS(StatusCode int, body string) *Response {
	return &Response{
		ContentType: "text/css",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) GIF(StatusCode int, body []byte) *Response {
	return &Response{
		ContentType: "image/gif",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) MP3(StatusCode int, body []byte) *Response {
	return &Response{
		ContentType: "audio/mpeg",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) MP4(StatusCode int, body []byte) *Response {
	return &Response{
		ContentType: "video/mp4",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) CSV(StatusCode int, body string) *Response {
	return &Response{
		ContentType: "text/csv",
		StatusCode:  StatusCode,
		Body:        body,
	}
}

func (context Context) SVG(StatusCode int, body string) *Response {
	return &Response{
		ContentType: "image/svg+xml",
		StatusCode:  StatusCode,
		Body:        body,
	}
}
