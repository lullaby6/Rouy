package rouy

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Response struct {
	ContentType string
	StatusCode  int
	Body        interface{}
}

func FuncResponse(statusCode int, contentType string, body interface{}) *Response {
	return &Response{
		ContentType: contentType,
		StatusCode:  statusCode,
		Body:        body,
	}
}

func respondeHandler(w http.ResponseWriter, response *Response) bool {
	if response == nil {
		return false
	}

	if response.StatusCode == 0 {
		response.StatusCode = 200
	}

	if response.ContentType == "" {
		response.ContentType = "text/plain"
	}

	w.WriteHeader(response.StatusCode)
	w.Header().Set("Content-Type", response.ContentType)

	if strings.Contains(response.ContentType, "image") {
		w.Write(response.Body.([]byte))
		return true
	} else if response.ContentType == "application/json" {
		json.NewEncoder(w).Encode(response.Body)
		return true
	} else if response.ContentType == "application/pdf" {
		w.Write(response.Body.([]byte))
		return true
	} else if response.ContentType == "application/zip" {
		w.Write(response.Body.([]byte))
		return true
	}

	w.Write([]byte(response.Body.(string)))
	return true
}
