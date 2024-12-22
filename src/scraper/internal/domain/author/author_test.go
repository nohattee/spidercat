package author_test

import (
	"testing"

	"github.com/nohattee/spidercat/src/scraper/internal/domain/author"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthor(t *testing.T) {
	expected := "test"
	authorName := author.NewAuthor(expected)
	assert.NotEmpty(t, authorName.ID())
	assert.Equal(t, expected, authorName.Name())
}
