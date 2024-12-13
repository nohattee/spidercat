package source_test

import (
	"testing"

	"scraper/internal/domain/source"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var s *source.Source
	var err error

	s, err = source.Parse("https://www.google.com/search?q=test")
	if assert.NoError(t, err) {
		assert.Equal(t, source.GoogleDomain, s.Domain())
	}

	s, err = source.Parse("https://google.com/search?q=test")
	if assert.NoError(t, err) {
		assert.Equal(t, source.GoogleDomain, s.Domain())
	}

	s, err = source.Parse("https://facebook.com/")
	assert.Error(t, err)
	assert.Nil(t, s)

	s, err = source.Parse("https://www.facebook.com/")
	assert.Error(t, err)
	assert.Nil(t, s)
}
