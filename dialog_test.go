package godialog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileDialog(t *testing.T) {
	assert.Equal(t, &FileDialog{}, NewFileDialog(), "NewFileDialog should return a new instance of FileDialog")
}
func TestFileDialogFilters(t *testing.T) {
	assert := assert.New(t)
	fd := &FileDialog{}

	filters := FileFilters{
		{"Image Files", []string{"*.jpg", "*.jpeg", "*.png"}},
	}
	fd.SetFilters(filters)
	assert.Equal(filters, fd.Filters(), "Filters should be set correctly")

	newFilter := FileFilter{
		"Text Files",
		[]string{"*.txt"},
	}
	fd.AddFilter(newFilter)
	assert.Equal(append(filters, newFilter), fd.Filters(), "Filter should be added correctly")
}

func TestFileDialogFallback(t *testing.T) {
	fd := &FileDialog{}

	// Test setting and getting fallback
	mockFallback := &mockFallbackDialog{}
	fd.SetFallback(mockFallback)
	assert.Same(t, mockFallback, fd.Fallback())
}

// Mock implementation of FallbackDialog for testing
type mockFallbackDialog struct{}

func (m *mockFallbackDialog) Open(title string, initialDirectory string, filters FileFilters, cb DialogCallback) {
	cb("", nil)
}

func (m *mockFallbackDialog) Save(title string, initialDirectory string, filters FileFilters, cb DialogCallback) {
	cb("", nil)
}
