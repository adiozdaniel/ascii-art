package models

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

// AppConfig holds configuration for the application.
type AppConfig struct {
	UseCache        bool
	TemplateCache   map[string]*template.Template
	BannerFileCache map[string]string
}

// App is the instance of AppConfig.
func App() *AppConfig {
	return &AppConfig{
		TemplateCache:   make(map[string]*template.Template),
		BannerFileCache: make(map[string]string),
		UseCache:        false,
	}
}

// CreateBannerCache caches banner files
func (a *AppConfig) CreateBannerCache() (map[string]string, error) {
	baseDir := instance.input.GetProjectRoot("views/static", "banners")
	bannerDir := filepath.Join(baseDir, "*.txt")
	banners, err := filepath.Glob(bannerDir)
	if err != nil {
		return nil, fmt.Errorf("error globbing banners: %v", err)
	}

	for _, banner := range banners {
		contents, err := os.ReadFile(banner)
		if err != nil {
			return nil, fmt.Errorf("error reading banner file %s: %v", banner, err)
		}
		a.BannerFileCache[filepath.Base(banner)] = string(contents)
	}

	return a.BannerFileCache, nil
}

// CreateTemplateCache is a helper function to cache all HTML templates as a map
func (a *AppConfig) CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	baseDir := instance.input.GetProjectRoot("views", "templates")

	templatesDir := filepath.Join(baseDir, "*.page.html")
	pages, err := filepath.Glob(templatesDir)
	if err != nil {
		return myCache, fmt.Errorf("error globbing templates: %v", err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, fmt.Errorf("error parsing page %s: %v", name, err)
		}

		layoutsPath := filepath.Join(baseDir, "*.layout.html")
		matches, err := filepath.Glob(layoutsPath)
		if err != nil {
			return myCache, fmt.Errorf("error finding layout files: %v", err)
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(layoutsPath)
			if err != nil {
				return myCache, fmt.Errorf("error parsing layout files: %v", err)
			}
		}

		a.TemplateCache[name] = ts
	}
	return a.TemplateCache, nil
}
