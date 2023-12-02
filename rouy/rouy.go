package rouy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type HandleFunc func(context Context) *Response

type Response struct {
	ContentType string
	StatusCode  int
	Body        interface{}
}

func HandleResponse(statusCode int, contentType string, body interface{}) *Response {
	return &Response{
		ContentType: contentType,
		StatusCode:  statusCode,
		Body:        body,
	}
}

type Config struct {
	Host string
	Port string
}

type Rouy struct {
	routes []Route
	config Config
}

func (rouy Rouy) Listen(config Config) error {
	rouy.config = config

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, route := range rouy.routes {
			path := route.Path

			context := Context{
				Request:  r,
				Response: w,
				Query:    r.URL.Query(),
				Method:   r.Method,
				Path:     r.URL.Path,
			}

			if r.Method == route.Method && r.URL.Path == path {
				response := route.Handler(context)

				if response == nil {
					return
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
					return
				} else if response.ContentType == "application/json" {
					json.NewEncoder(w).Encode(response.Body)
					return
				} else if response.ContentType == "application/pdf" {
					w.Write(response.Body.([]byte))
					return
				} else if response.ContentType == "application/zip" {
					w.Write(response.Body.([]byte))
					return
				}

				w.Write([]byte(response.Body.(string)))
				return
			}
		}

		http.NotFound(w, r)
	})

	fullUrl := fmt.Sprintf("%s:%s", config.Host, config.Port)
	fmt.Printf("Listening on %s", fullUrl)
	err := http.ListenAndServe(fullUrl, nil)
	if err != nil {
		return err
	}
	return nil
}
