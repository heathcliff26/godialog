//go:build linux

package godialog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConvertFiltersToFreedesktopFilter(t *testing.T) {
	filter := FileFilters{
		{"Text Files", []string{".txt", ".md"}},
		{"Image Files", []string{".png", ".jpg"}},
		{"Nothing", []string{}},
		{"Nil", nil},
	}
	expected := []freedesktopFilter{
		{
			Name: "Text Files",
			Rules: []freedesktopFilterRule{
				{Pattern: "*.txt"},
				{Pattern: "*.md"},
			},
		},
		{
			Name: "Image Files",
			Rules: []freedesktopFilterRule{
				{Pattern: "*.png"},
				{Pattern: "*.jpg"},
			},
		},
		{
			Name: "Nothing",
		},
		{
			Name: "Nil",
		},
	}

	converted := convertFiltersToFreedesktopFilter(filter)

	assert.Equal(t, expected, converted, "Should convert filters correctly")
}

func TestDialogDoesNotBlock(t *testing.T) {
	fd := NewFileDialog()

	t.Run("Open", func(t *testing.T) {
		done := make(chan struct{})

		go func() {
			fd.Open("Test Open", func(path string, err error) {})
			close(done)
		}()

		select {
		case <-done:
			// Test passed
		case <-time.After(1 * time.Second):
			t.Error("Dialog did not return in time")
		}
	})
	t.Run("Save", func(t *testing.T) {
		done := make(chan struct{})

		go func() {
			fd.Save("Test Open", func(path string, err error) {})
			close(done)
		}()

		select {
		case <-done:
			// Test passed
		case <-time.After(1 * time.Second):
			t.Error("Dialog did not return in time")
		}
	})
}
