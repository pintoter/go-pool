package transport

import (
	"net/http"
)

func (h *Handler) adminHandler(w http.ResponseWriter, r *http.Request) {
	if !h.checkLoginState(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	renderTemplate(w, "internal/templates/admin.html", nil)
}
