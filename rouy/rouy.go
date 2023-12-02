package rouy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type JSON map[string]interface{}

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

type Config struct {
	Host   string
	Port   string
	Logger bool
}

type Rouy struct {
	Middlewares []Route
	Routes      []Route
	NotFound    HandleFunc
	Config      Config
}

func New() *Rouy {
	return &Rouy{}
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

func requestHandler(route Route, w http.ResponseWriter, r *http.Request, context Context) bool {
	response := route.Handler(context)

	responseHandlerResult := respondeHandler(w, response)

	return responseHandlerResult
}

func (rouy Rouy) Listen(config Config) error {
	rouy.Config = config

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

		if rouy.Config.Logger {
			fmt.Printf("\n [Rouy Request]\n Method: %s,\n URL: %s,\n Host: %s,\n RemoteAddr: %s,\n Headers: %v,\n Body: %v\n\n",
				r.Method, r.URL, r.Host, r.RemoteAddr, r.Header, body)
		}

		context := Context{
			Request:  r,
			Response: w,
			Query:    r.URL.Query(),
			Method:   r.Method,
			Path:     r.URL.Path,
			Body:     body,
		}

		for _, middleware := range rouy.Middlewares {
			if r.Method == middleware.Method && r.URL.Path == middleware.Path {
				if requestHandler(middleware, w, r, context) {
					return
				}
			}
		}

		for _, route := range rouy.Routes {
			if r.Method == route.Method && r.URL.Path == route.Path {
				if requestHandler(route, w, r, context) {
					return
				}
			}
		}

		if rouy.NotFound != nil {
			if requestHandler(Route{
				Handler: rouy.NotFound,
			}, w, r, context) {
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
