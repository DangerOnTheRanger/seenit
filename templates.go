package seenit

import (
	"html/template"
	"io"
	"path"
)

const (
	templateDir     = "templates"
	baseTemplate    = "base.html"
	landingTemplate = "landing.html"
	uploadTemplate  = "upload.html"
	seenTemplate    = "seen.html"
	unseenTemplate  = "unseen.html"
)

func buildTemplate(filename string) *template.Template {
	fullPath := path.Join(templateDir, filename)
	baseTemplatePath := path.Join(templateDir, baseTemplate)
	return template.Must(template.ParseFiles(baseTemplatePath, fullPath))
}

func RenderLanding(communities []Community, w io.Writer) error {
	landing := buildTemplate(landingTemplate)
	data := struct {
		Communities []Community
	}{
		Communities: communities,
	}
	return landing.ExecuteTemplate(w, baseTemplate, data)
}

func RenderUpload(community string, w io.Writer) error {
	upload := buildTemplate(uploadTemplate)
	data := struct {
		Community string
	}{
		Community: community,
	}
	return upload.ExecuteTemplate(w, baseTemplate, data)
}

func RenderSeen(w io.Writer) error {
	seen := buildTemplate(seenTemplate)
	return seen.ExecuteTemplate(w, baseTemplate, nil)
}

func RenderUnseen(w io.Writer) error {
	unseen := buildTemplate(unseenTemplate)
	return unseen.ExecuteTemplate(w, baseTemplate, nil)
}
