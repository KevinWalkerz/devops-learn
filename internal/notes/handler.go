package notes

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/notes", h.list)
	mux.HandleFunc("POST /api/notes", h.create)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list notes")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

type createRequest struct {
	Text string `json:"text"`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	req.Text = strings.TrimSpace(req.Text)
	if req.Text == "" {
		writeError(w, http.StatusBadRequest, "text is required")
		return
	}

	note, err := h.store.Create(r.Context(), req.Text)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create note")
		return
	}
	writeJSON(w, http.StatusCreated, note)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
