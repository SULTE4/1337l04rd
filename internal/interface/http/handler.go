package http

import (
	"1337b04rd/internal/domain/post"
	"net/http"
	"text/template"
)

type Handler struct {
	postService   *post.Service
	templateCache map[string]*template.Template
}

func NewHandler(postService *post.Service, templateCache map[string]*template.Template) *Handler {
	return &Handler{
		postService: postService,
	}
}

func (h *Handler) Catalog(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Showing posts"))
}

func (h *Handler) Archive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing archived posts"))

}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create page"))
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create page submission"))

}

func (h *Handler) ViewPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing post"))

}

func (h *Handler) Error(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing error page"))

}
