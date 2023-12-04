package rouy

import "net/http"

func (rouy *Rouy) Static(path, dir_path string) *Rouy {

	http.Handle(path, http.FileServer(http.Dir(dir_path)))

	return rouy
}
