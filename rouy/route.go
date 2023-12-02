package rouy

type Route struct {
	Method  string
	Path    string
	Handler HandleFunc
}

func (rouy *Rouy) Route(route Route) *Rouy {
	rouy.Routes = append(rouy.Routes, route)
	return rouy
}

func (rouy *Rouy) GET(path string, handler HandleFunc) *Rouy {
	route := Route{
		Method:  "GET",
		Path:    path,
		Handler: handler,
	}

	rouy.Routes = append(rouy.Routes, route)
	return rouy
}

func (rouy *Rouy) POST(path string, handler HandleFunc) *Rouy {
	route := Route{
		Method:  "POST",
		Path:    path,
		Handler: handler,
	}

	rouy.Routes = append(rouy.Routes, route)
	return rouy
}

func (rouy *Rouy) PUT(path string, handler HandleFunc) *Rouy {
	route := Route{
		Method:  "PUT",
		Path:    path,
		Handler: handler,
	}

	rouy.Routes = append(rouy.Routes, route)
	return rouy
}

func (rouy *Rouy) DELETE(path string, handler HandleFunc) *Rouy {
	route := Route{
		Method:  "DELETE",
		Path:    path,
		Handler: handler,
	}

	rouy.Routes = append(rouy.Routes, route)
	return rouy
}

func (rouy *Rouy) PATCH(path string, handler HandleFunc) *Rouy {
	route := Route{
		Method:  "PATCH",
		Path:    path,
		Handler: handler,
	}

	rouy.Routes = append(rouy.Routes, route)
	return rouy
}

func (rouy *Rouy) HEAD(path string, handler HandleFunc) *Rouy {
	route := Route{
		Method:  "HEAD",
		Path:    path,
		Handler: handler,
	}

	rouy.Routes = append(rouy.Routes, route)
	return rouy
}

func (rouy *Rouy) OPTIONS(path string, handler HandleFunc) *Rouy {
	route := Route{
		Method:  "OPTIONS",
		Path:    path,
		Handler: handler,
	}

	rouy.Routes = append(rouy.Routes, route)
	return rouy
}

func (rouy *Rouy) TRACE(path string, handler HandleFunc) *Rouy {
	route := Route{
		Method:  "TRACE",
		Path:    path,
		Handler: handler,
	}

	rouy.Routes = append(rouy.Routes, route)
	return rouy
}

func (rouy *Rouy) CONNECT(path string, handler HandleFunc) *Rouy {
	route := Route{
		Method:  "CONNECT",
		Path:    path,
		Handler: handler,
	}

	rouy.Routes = append(rouy.Routes, route)
	return rouy
}
