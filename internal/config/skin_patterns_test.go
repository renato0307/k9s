package config_test

import (
	"testing"

	"github.com/derailed/k9s/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestSkinPatternsLoad(t *testing.T) {
	sp := config.NewSkinPatterns()
	assert.Nil(t, sp.LoadSkinPatterns("testdata/skin_patterns.yml"))

	assert.Equal(t, 3, len(sp.Patterns))

	k, ok := sp.Patterns["dev"]
	assert.True(t, ok)
	assert.Equal(t, "dev-.*", k.Pattern)
	assert.Equal(t, "gruvbox-dark", k.Skin)

	k, ok = sp.Patterns["lab"]
	assert.True(t, ok)
	assert.Equal(t, "lab-.*", k.Pattern)
	assert.Equal(t, "stock", k.Skin)
}

func TestSkinPatternsMatch(t *testing.T) {
	sp := config.NewSkinPatterns()
	sp.LoadSkinPatterns("testdata/skin_patterns.yml")

	assert.Contains(t, sp.Match("dev-01"), "gruvbox-dark")
	assert.Contains(t, sp.Match("lab-01"), "stock")
	assert.Empty(t, sp.Match("nomatch-01"))
}
