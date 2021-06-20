package seenit

import (
	"image"
	"net/http"
)

type Handler func(db Database, w http.ResponseWriter, r *http.Request)
type BoundHandler func(w http.ResponseWriter, r *http.Request)


func BindHandler(h Handler, db Database) BoundHandler {
	return func (w http.ResponseWriter, r *http.Request) {
		h(db, w, r)
	}
}

func ServeLanding(w http.ResponseWriter, r *http.Request) {
	err := RenderLanding(w)
	if err != nil {
		// TODO: how to handle errors?
	}
}


func ServeUpload(w http.ResponseWriter, r *http.Request) {
	community := r.FormValue("community")
	err := RenderUpload(community, w)
	if err != nil {
		// TODO: how to handle errors?
	}
}


func ServeResult(db Database, w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("image")
	if err != nil {
		// TODO: how to handle errors?
	}
	defer file.Close()
	image, _, err := image.Decode(file)
	if err != nil {
		// TODO: how to handle errors?
	}
	hash, err := ImageToHash(image)
	if err != nil {
		// TODO: how to handle errors?
	}
	community := r.FormValue("community")
	seen, err := HaveSeen(community, hash, db)
	if err != nil {
		// TODO: how to handle errors?
	}
	if seen {
		err = RenderSeen(w)
		if err != nil {
			// TODO: how to handle errors?
		}
	} else {
		err = RecordHash(community, hash, db)
		if err != nil {
			// TODO: how to handle errors?
		}
		err = RenderUnseen(w)
		if err != nil {
			// TODO: how to handle errors?
		}
	}
}
