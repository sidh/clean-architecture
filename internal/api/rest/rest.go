package rest

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sidh/clean-architecture/internal/logic"
)

// Server implements HTTP server
type Server struct {
	c logic.Core
}

// New constructs HTTP server
func New(c logic.Core) *Server {
	return &Server{c: c}
}

// Serve initiates HTTP server serving
func (s *Server) Serve() error {
	http.HandleFunc("/", s.handleRequest)
	return http.ListenAndServe(":8080", nil)
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleLoad(w, r)
	case "PUT":
		fallthrough
	case "POST":
		s.handleStore(w, r)
	}
}

func (s *Server) handleStore(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, ErrValueMissing.Error(), 400)
		return
	}
	defer r.Body.Close()

	user := r.URL.Query().Get(QueryParamUser)
	if user == "" {
		http.Error(w, ErrUserMissing.Error(), 400)
		return
	}

	key := r.URL.Query().Get(QueryParamKey)
	if key == "" {
		http.Error(w, ErrKeyMissing.Error(), 400)
		return
	}

	value, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, ErrValueMissing.Error(), 400)
		return
	}

	var uv userValue
	if err := uv.Unmarshal(value); err != nil {
		http.Error(w, fmt.Errorf(ErrValueInvalid.Error()+": %w", err).Error(), 400)
		return
	}

	if err := s.c.Store(r.Context(), user, key, toValueModel(uv)); err != nil {
		switch err {
		case logic.ErrAuthFailed:
			fallthrough
		case logic.ErrActionDenied:
			http.Error(w, ErrForbidden.Error(), 403)
			return
		case logic.ErrActionFailed:
			fallthrough
		default:
			http.Error(w, ErrInternalError.Error(), 500)
			return
		}
	}

	fmt.Fprintf(w, "Stored key '%s'\n", key)
}

func (s *Server) handleLoad(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get(QueryParamUser)
	if user == "" {
		http.Error(w, ErrUserMissing.Error(), 400)
		return
	}

	key := r.URL.Query().Get(QueryParamKey)
	if key == "" {
		http.Error(w, ErrKeyMissing.Error(), 400)
		return
	}

	value, err := s.c.Load(r.Context(), user, key)
	if err != nil {
		switch err {
		case logic.ErrAuthFailed:
			fallthrough
		case logic.ErrActionDenied:
			http.Error(w, ErrForbidden.Error(), 403)
			return
		case logic.ErrKeyNotFound:
			http.Error(w, ErrKeyNotFound.Error(), 404)
			return
		case logic.ErrActionFailed:
			fallthrough
		default:
			http.Error(w, ErrInternalError.Error(), 500)
			return
		}
	}

	sv := fromValueModel(value)
	v, err := sv.Marshal()
	if err != nil {
		http.Error(w, ErrValueInvalid.Error(), 400)
		return
	}

	fmt.Fprint(w, string(v))
}
