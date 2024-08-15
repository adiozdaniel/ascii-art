package handlers

import (
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/models"
	"github.com/adiozdaniel/ascii-art/internals/renders"
	"github.com/adiozdaniel/ascii-art/pkg/helpers"
)

const sessionCookieName = "session_id"

// Repository handles HTTP requests and application state
type Repository struct {
	app *models.StateManager
}

// NewRepository creates a new Repository instance
func NewRepo(sm *models.StateManager) *Repository {
	return &Repository{app: sm}
}

// HomeHandler handles the homepage route '/'
func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		renders.RenderTemplate(w, "login.page.html", m.app.GetTemplateData())
		return
	}

	data := m.app.GetSessionManager().GetSessionData(cookie.Value)
	renders.RenderTemplate(w, "home.page.html", data)
}

// SubmitHandler handles the output route '/ascii-art'
func (m *Repository) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		renders.RenderTemplate(w, "login.page.html", m.app.GetTemplateData())
		return
	}

	data := m.app.GetSessionManager().GetSessionData(cookie.Value)

	if r.Method == http.MethodGet {
		renders.RenderTemplate(w, "ascii.page.html", data)
		return
	}

	err = r.ParseForm()
	if err != nil {
		m.BadRequestHandler(w, r)
		return
	}

	if r.Method == http.MethodPost && r.Form.Get("textInput") == "" {
		m.BadRequestHandler(w, r)
		return
	}

	m.app.GetInput().Flags["input"] = r.Form.Get("textInput")
	banner := m.app.GetInput().BannerFile[r.Form.Get("Font")]

	if banner == "" {
		m.BadRequestHandler(w, r)
		return
	}

	m.app.GetInput().Flags["font"] = banner
	err = helpers.FileContents(banner)
	if err != nil {
		m.ServerErrorHandler(w, r)
		return
	}

	output := ascii.Output(m.app.GetInput().Flags["input"])
	nonasciis := ascii.NonAsciiOutput()

	td := data.StringMap
	td["ascii"] = output
	td["nonasciis"] = nonasciis

	renders.RenderTemplate(w, "ascii.page.html", data)
}

// LoginHandler handles user login and session creation
func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var session *models.Session
		cookie, err := r.Cookie("session_id")
		if err == nil {
			session, _ = m.app.GetSessionManager().GetSession(cookie.Value)
		}

		if session == nil {
			renders.RenderTemplate(w, "login.page.html", m.app.GetTemplateData())
			return
		}

		if session != nil && session.CRSFToken != "" && r.Method == "GET" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			m.ServerErrorHandler(w, r)
			return
		}

		form := m.app.GetTemplateData().Form
		form.Errors.Clear()

		username := r.Form.Get("username")
		form.Required(r, "username")

		if !form.IsValidForm() {
			w.WriteHeader(http.StatusUnauthorized)
			renders.RenderTemplate(w, "login.page.html", m.app.GetTemplateData())
			return
		}

		if username != "" {
			session := m.app.GetSessionManager().CreateSession()
			var data = m.app.GetSessionManager().GetSessionData(session.CRSFToken)
			data.StringMap["username"] = data.CapitalizeFirst(username)

			http.SetCookie(w, &http.Cookie{
				Name:     sessionCookieName,
				Value:    session.CRSFToken,
				Path:     "/",
				Expires:  time.Now().Add(30 * time.Minute),
				HttpOnly: true,
				Secure:   false, // TODO replace with true for production
			})

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

	}
}

// LogoutHandler handles user logout by deleting the session
func (m *Repository) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	sessionID := cookie.Value
	var data = m.app.GetSessionManager().GetSessionData(sessionID)
	td := data.StringMap
	td["username"] = ""
	sm := m.app.GetSessionManager()
	sm.DeleteSession(sessionID)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Secure:   false, // TODO replace with true for production
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// NotFoundHandler handles unknown routes; 404 status
func (m *Repository) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		renders.RenderTemplate(w, "login.page.html", m.app.GetTemplateData())
		return
	}

	data := m.app.GetSessionManager().GetSessionData(cookie.Value)
	w.WriteHeader(http.StatusNotFound)
	renders.RenderTemplate(w, "notfound.page.html", data)
}

// BadRequestHandler handles bad requests routes
func (m *Repository) BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		renders.RenderTemplate(w, "login.page.html", m.app.GetTemplateData())
		return
	}

	data := m.app.GetSessionManager().GetSessionData(cookie.Value)
	w.WriteHeader(http.StatusBadRequest)
	renders.RenderTemplate(w, "badrequest.page.html", data)
}

// ServerErrorHandler handles server failures that result in status 500
func (m *Repository) ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		renders.RenderTemplate(w, "login.page.html", m.app.GetTemplateData())
		return
	}

	data := m.app.GetSessionManager().GetSessionData(cookie.Value)
	w.WriteHeader(http.StatusInternalServerError)
	renders.RenderTemplate(w, "serverError.page.html", data)
}

// AboutHandler handles the about page route '/about'
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		renders.RenderTemplate(w, "login.page.html", m.app.GetTemplateData())
		return
	}

	data := m.app.GetSessionManager().GetSessionData(cookie.Value)
	renders.RenderTemplate(w, "about.page.html", data)
}

// ContactHandler handles the contact page route '/contact'
func (m *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	data := m.app.GetSessionManager().GetSessionData(cookie.Value)

	if r.Method == http.MethodGet {
		form := m.app.GetTemplateData().Form
		form.Errors.Clear()
		m.app.GetTemplateData().StringMap["success"] = ""

		renders.RenderTemplate(w, "contact.page.html", data)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			m.BadRequestHandler(w, r)
			return
		}

		form := m.app.GetTemplateData().Form
		form.Errors.Clear()

		form.Required(r, "name", "email", "message")

		if !form.IsValidForm() {
			w.WriteHeader(http.StatusBadRequest)
			renders.RenderTemplate(w, "contact.page.html", m.app.GetTemplateData())
			return
		}

		m.app.GetTemplateData().StringMap["success"] = "email successfully sent"

		w.WriteHeader(http.StatusAccepted)
		renders.RenderTemplate(w, "contact.page.html", data)
	}
}

// DownloadHandler handles file download requests.
func (m *Repository) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("ascii-art.txt")
	m.app.GetInput().Flags["output"] = "ascii-art.txt"

	output := ascii.Output(m.app.GetInput().Flags["input"])
	ascii.LogOutput(strings.ReplaceAll(output, "$", " "))

	w.Header().Set("Content-Disposition", "attachment; filename=\""+filePath+"\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, filePath)
}
