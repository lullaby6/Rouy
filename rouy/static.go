package rouy

import "net/http"

func (rouy *Rouy) Static(path, dir_path string) *Rouy {

	path += "/"

	http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir_path))))

	return rouy
}
