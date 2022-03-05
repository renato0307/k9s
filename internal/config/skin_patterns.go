package config

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v2"
)

// K9sSkinPatterns manages K9s SkinPatterns.
var K9sSkinPatterns = filepath.Join(K9sHome(), "skin_patterns.yml")

// SkinPatterns represents a collection of skin patterns.
type SkinPatterns struct {
	Patterns map[string]Pattern `yaml:"skinPatterns"`
}

// HotKey describes a K9s skin pattern.
type Pattern struct {
	Pattern string `yaml:"pattern"`
	Skin    string `yaml:"skin"`

	regexp *regexp.Regexp
}

// NewSkinPatterns returns a new plugin.
func NewSkinPatterns() SkinPatterns {
	return SkinPatterns{
		Patterns: make(map[string]Pattern),
	}
}

// Load K9s plugins.
func (sp SkinPatterns) Load() error {
	return sp.LoadSkinPatterns(K9sSkinPatterns)
}

// LoadSkinPatterns loads plugins from a given file.
func (sp SkinPatterns) LoadSkinPatterns(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var sps SkinPatterns
	if err := yaml.Unmarshal(f, &sps); err != nil {
		return err
	}
	for k, v := range sps.Patterns {
		v.regexp = regexp.MustCompile(v.Pattern)
		sp.Patterns[k] = v
	}

	return nil
}

// Match finds a skin for a context using a regular expression.
// It returns the first skin that matches the context name.
func (sp SkinPatterns) Match(context string) string {
	for _, v := range sp.Patterns {
		result := v.regexp.FindString(context)
		if result != "" {
			skinFile := fmt.Sprintf("%s.yml", v.Skin)
			return filepath.Join(K9sHome(), skinFile)
		}
	}

	return ""
}
