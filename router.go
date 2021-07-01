package seenit

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"image"
	"net/http"
	"time"
)

const (
	commCookieName = "saved-communities"
)

type Handler func(db Database, w http.ResponseWriter, r *http.Request)
type BoundHandler func(w http.ResponseWriter, r *http.Request)

func BindHandler(h Handler, db Database) BoundHandler {
	return func (w http.ResponseWriter, r *http.Request) {
		h(db, w, r)
	}
}

func ServeLanding(w http.ResponseWriter, r *http.Request) {
	commCookie, err := r.Cookie(commCookieName)
	var commCookieJson string
	if errors.Is(err, http.ErrNoCookie) {
		commCookieJson = ""
	} else if err == nil {
		commCookieJson = commCookie.Value
	} else {
		// TODO: how to handle errors?
		return
	}
	var communities []Community
	decodedComms, err := base64.StdEncoding.DecodeString(commCookieJson)
	if err != nil {
		// TODO: how to handle erros?
		return
	}
	json.Unmarshal(decodedComms, &communities)
	err = RenderLanding(communities, w)
	if err != nil {
		// TODO: how to handle errors?
	}
}


func ServeUpload(w http.ResponseWriter, r *http.Request) {
	community := r.FormValue("community")
	commCookie, err := r.Cookie(commCookieName)
	var commCookieJson string
	if errors.Is(err, http.ErrNoCookie) {
		commCookieJson = ""
	} else if err == nil {
		commCookieJson = commCookie.Value
	} else {
		// TODO: how to handle errors?
		return
	}
	var communities []Community
	decodedComms, err := base64.StdEncoding.DecodeString(commCookieJson)
	if err != nil {
		// TODO: how to handle erros?
		return
	}
	json.Unmarshal(decodedComms, &communities)
	alreadyAdded := false
	for _, c := range communities {
		if c.Name == community {
			alreadyAdded = true
			break
		}
	}
	if !alreadyAdded {
		communities = append(communities, Community{Name: community})
		serializedComms, err := json.Marshal(communities)
		if err != nil {
			// TODO: how to handle errors?
		}
		expiration := time.Now().AddDate(1, 0, 0)
		http.SetCookie(w, &http.Cookie{Name: commCookieName,
			Value: base64.StdEncoding.EncodeToString(serializedComms),
			Expires: expiration,
		})
	}
	err = RenderUpload(community, w)
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
