package handlers

import (
	"net/http"
	"time"

	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/internals/config"
	"github.com/adiozdaniel/ascii-art/internals/models"
	"github.com/adiozdaniel/ascii-art/internals/renders"
	"github.com/adiozdaniel/ascii-art/utils"
)

// Repository is a struct to hold the application configuration
type Repository struct {
	App *config.AppConfig
}

// Repo is a global variable to hold the Repository instance
var Repo *Repository

// NewRepo initializes a new Repository instance
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers initializes the repository instance and sets the handlers
func NewHandlers(m *Repository) {
	Repo = m
}

// HomeHandler handles the homepage route '/'
func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	m.SetCookieHandler(w, r)
	renders.RenderTemplate(w, r, "home.page.html", nil)
}

// SubmitHandler handles the output route '/ascii-art'
func (m *Repository) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renders.RenderTemplate(w, r, "ascii.page.html", &models.TemplateData{})
		return
	}

	if r.Method == "POST" && r.FormValue(("textInput")) == "" {
		Repo.BadRequestHandler(w, r)
		return
	}

	utils.Inputs.Input = r.FormValue("textInput")
	banner := utils.BannerFiles[r.FormValue("FileName")]

	if banner == "" {
		utils.Inputs.BannerPath = utils.BannerFiles["standard"]
	} else {
		utils.Inputs.BannerPath = banner
	}

	_, err := ascii.FileContents()
	if err != nil {
		Repo.ServerErrorHandler(w, r)
		return
	}

	output := ascii.Output(utils.Inputs.Input)
	noasciis := utils.NonAsciiOutput()
	asciiMap := map[string]string{
		"Body":    "\n" + output,
		"NoAscii": noasciis,
	}

	renders.RenderTemplate(w, r, "ascii.page.html", &models.TemplateData{Ascii: asciiMap})
}

// NotFoundHandler handles unknown routes; 404 status
func (m *Repository) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renders.RenderTemplate(w, r, "notfound.page.html", nil)
}

// BadRequestHandler handles bad requests routes
func (m *Repository) BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	renders.RenderTemplate(w, r, "badrequest.page.html", nil)
}

// ServerErrorHandler handles server failures that result in status 500
func (m *Repository) ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	renders.RenderTemplate(w, r, "serverError.page.html", nil)
}

// AboutHandler handles the about page route '/about'
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "about.page.html", nil)
}

// ContactHandler handles the contact page route '/contact'
func (m *Repository) ContactHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "contact.page.html", nil)
}

// SetCookieHandler sets a session cookie for the client
func (m *Repository) SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	m.App.Sessions = make(map[string]string)
	sessionID := generateSessionID()
	m.App.Sessions[sessionID] = r.RemoteAddr

	expiration := time.Now().Add(15 * time.Minute)
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  expiration,
		HttpOnly: true,
		Secure:   false, // TODO Set to true before deploying to production
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("Session cookie has been set"))
}

// GetSessionHandler retrieves session data based on the session cookie
func (m *Repository) GetSessionHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}

	sessionID := cookie.Value
	userID, ok := m.App.Sessions[sessionID]
	if !ok {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Session is valid for user: " + userID))
}

// Generate a session ID (in a real application, use a more secure method)
func generateSessionID() string {
	return time.Now().Format("20060102150405")
}
