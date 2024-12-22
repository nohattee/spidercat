package category_test

import (
	"testing"

	"github.com/nohattee/spidercat/src/scraper/internal/domain/category"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	expected := "test"
	testCategory := category.NewCategory(expected)
	assert.NotEmpty(t, testCategory.ID())
	assert.Equal(t, expected, testCategory.Name())
}
