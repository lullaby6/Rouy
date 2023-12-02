package rouy

func (rouy *Rouy) Middleware(middleware Route) *Rouy {
	if middleware.Path == "" {
		middleware.Path = "/"
	}

	if middleware.Method == "" {
		middleware.Method = "GET"
	}

	rouy.Middlewares = append(rouy.Middlewares, middleware)
	return rouy
}

func (rouy *Rouy) Use(middleware Route) *Rouy {
	if middleware.Path == "" {
		middleware.Path = "/"
	}

	if middleware.Method == "" {
		middleware.Method = "GET"
	}

	rouy.Middlewares = append(rouy.Middlewares, middleware)
	return rouy
}
