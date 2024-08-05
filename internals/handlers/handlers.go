package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/models"
	"github.com/adiozdaniel/ascii-art/internals/renders"
	"github.com/adiozdaniel/ascii-art/pkg/helpers"
)

// get the app state manager
var (
	sm      = models.GetStateManager()
	appData = sm.GetInput()
	td      = sm.GetTemplateData()
	ck      = sm.GetSessionManager()
)

const sessionCookieName = "session_id"

// Repository handles HTTP requests and application state
type Repository struct {
	AppData *models.StateManager
}

// NewRepository creates a new Repository instance
func NewRepo(sm *models.StateManager) *Repository {
	return &Repository{AppData: sm}
}

// HomeHandler handles the homepage route '/'
func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	var session *models.Session

	if err == nil {
		session, _ = m.AppData.GetSessionManager().GetSession(cookie.Value)
	}

	if session == nil {
		session = m.AppData.GetSessionManager().CreateSession()

		http.SetCookie(w, &http.Cookie{
			Name:    "session_id",
			Value:   session.CRSFToken,
			Path:    "/",
			Expires: time.Now().Add(30 * time.Minute),
		})
	}

	ctx := context.WithValue(r.Context(), ck.SessionKey, session)
	r = r.WithContext(ctx)

	if session.CRSFToken != "" && r.Method == "GET" {
		renders.RenderTemplate(w, "home.page.html", nil)
	}
}

// SubmitHandler handles the output route '/ascii-art'
func (m *Repository) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	var session *models.Session

	if err == nil {
		session, _ = m.AppData.GetSessionManager().GetSession(cookie.Value)
	}

	// fmt.Printf("Session ID: %s\n", session.CRSFToken)  // Debugging
	if session == nil || session.CRSFToken == "" {
		m.NotFoundHandler(w, r)
		return
	}

	if r.FormValue("textInput") == "" && r.Method != "POST" {
		renders.RenderTemplate(w, "ascii.page.html", nil)
		return
	}

	if r.Method == "POST" && r.FormValue(("textInput")) == "" {
		m.BadRequestHandler(w, r)
		return
	}

	appData.Flags["input"] = r.FormValue("textInput")
	banner := appData.BannerFile[r.FormValue("Font")]

	if banner == "" {
		m.BadRequestHandler(w, r)
	}

	appData.Flags["font"] = banner
	err = helpers.FileContents(banner)
	if err != nil {
		m.NotFoundHandler(w, r)
		return
	}

	output := ascii.Output(appData.Flags["input"])
	nonasciis := ascii.NonAsciiOutput()

	td.StringMap["ascii"] = output
	td.StringMap["nonasciis"] = nonasciis

	renders.RenderTemplate(w, "ascii.page.html", td)
}

// NotFoundHandler handles unknown routes; 404 status
func (m *Repository) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renders.RenderTemplate(w, "notfound.page.html", nil)
}

// LoginHandler handles user login and session creation
func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renders.RenderTemplate(w, "login.page.html", nil)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")

		if username != "" {
			session := m.AppData.GetSessionManager().CreateSession()

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

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
}

// BadRequestHandler handles bad requests routes
func (m *Repository) BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	renders.RenderTemplate(w, "badrequest.page.html", nil)
}

// ServerErrorHandler handles server failures that result in status 500
func (m *Repository) ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	renders.RenderTemplate(w, "serverError.page.html", nil)
}

// AboutHandler handles the about page route '/about'
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.html", nil)
}

// ContactHandler handles the contact page route '/contact'
func (m *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "contact.page.html", nil)
}
