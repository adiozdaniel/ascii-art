package handlers

import (
	"context"
	"fmt"
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

	// Log the session ID for debugging purposes
	fmt.Fprintf(w, "Session ID: %s\n", session.CRSFToken)
	fmt.Printf("Session ID: %s\n", session.CRSFToken)

	// Render the home page
	renders.RenderTemplate(w, "home.page.html", nil)
}

// SubmitHandler handles the output route '/ascii-art'
func (m *Repository) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	session, ok := r.Context().Value(m.AppData.GetSessionManager()).(*models.Session)
	if !ok {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Session ID: %s\n", session.CRSFToken)
	fmt.Printf("Session ID: %s\n", session.CRSFToken)
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
	err := helpers.FileContents(banner)
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
