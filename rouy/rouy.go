package rouy

import (
	"encoding/json"
	"fmt"
	"io"
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
	Logger bool
}

func (rouy Rouy) Listen(config Config) error {
	rouy.config = config

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bodyRead, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		var body map[string]interface{}

		if len(bodyRead) > 0 {
			if err := json.Unmarshal(bodyRead, &body); err != nil {
				http.Error(w, "Error parsing request body", http.StatusBadRequest)
				return
			}
		}

		if rouy.Logger {
			fmt.Printf("\n [Rouy Request]\n Method: %s,\n URL: %s,\n Host: %s,\n RemoteAddr: %s,\n Headers: %v,\n Body: %v\n\n",
				r.Method, r.URL, r.Host, r.RemoteAddr, r.Header, body)
		}

		for _, route := range rouy.routes {
			path := route.Path

			context := Context{
				Request:  r,
				Response: w,
				Query:    r.URL.Query(),
				Method:   r.Method,
				Path:     r.URL.Path,
				Body:     body,
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

	fmt.Printf("Listening on %s\n", fullUrl)

	err := http.ListenAndServe(fullUrl, nil)
	if err != nil {
		return err
	}

	return nil
}
