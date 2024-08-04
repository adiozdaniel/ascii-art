package config

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

func (a *AppConfig) CreateBannerCache() (map[string]string, error) {
	myCache := make(map[string]string)
	baseDir := instance.input.GetProjectRoot("views/static", "banners")

	bannerDir := filepath.Join(baseDir, "*.txt")
	banners, err := filepath.Glob(bannerDir)
	if err != nil {
		return myCache, fmt.Errorf("error globbing banners: %v", err)
	}

	for _, banner := range banners {
		contents, err := os.ReadFile(banner)
		if err != nil {
			return myCache, fmt.Errorf("error reading banner file %s: %v", banner, err)
		}
		a.BannerFileCache[filepath.Base(banner)] = string(contents)
	}

	return myCache, nil
}
