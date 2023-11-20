package transport

import (
	"log"
	"net/http"
)

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	if h.checkLoginState(r) {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}

	if r.Method == http.MethodGet {
		renderTemplate(w, "internal/templates/auth.html", nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == h.credentials.Login && password == h.credentials.Password {
		session, err := h.sessionStore.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["isAdmin"] = true
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin", http.StatusFound)
	} else {
		http.Error(w, "invalid login credentials", http.StatusUnauthorized)
		return
	}
}

func (h *Handler) checkLoginState(r *http.Request) bool {
	session, err := h.sessionStore.Get(r, "session")
	if err != nil {
		log.Println(err)
		return false
	}
	value := session.Values["isAdmin"]
	if value == nil {
		return false
	}

	return value.(bool)
}
