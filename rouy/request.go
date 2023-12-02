package rouy

import "net/http"

func requestHandler(route Route, w http.ResponseWriter, r *http.Request, context Context) bool {
	response := route.Handler(context)

	responseHandlerResult := respondeHandler(w, response)

	return responseHandlerResult
}
