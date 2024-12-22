package tag_test

import (
	"testing"

	"github.com/nohattee/spidercat/src/scraper/internal/domain/tag"

	"github.com/stretchr/testify/assert"
)

func TestNewTag(t *testing.T) {
	expected := "test"
	testTag := tag.NewTag(expected)
	assert.NotEmpty(t, testTag.ID())
	assert.Equal(t, expected, testTag.Name())
}
