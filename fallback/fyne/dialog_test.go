package fyne

import (
	"sync"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/test"
	filedialog "github.com/heathcliff26/godialog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFyneFallbackDialog(t *testing.T) {
	a := test.NewApp()
	d := NewFyneFallbackDialog(a)

	require.NotNil(t, d, "Dialog should not be nil")

	assert := assert.New(t)

	assert.Equal(a, d.App, "App should be set correctly")
	assert.Equal(DefaultDialogHeight, d.Height, "Height should be set to default")
	assert.Equal(DefaultDialogWidth, d.Width, "Width should be set to default")
}

func TestDialogErrorWhenAppNil(t *testing.T) {
	d := &FyneFallbackDialog{}

	assert := assert.New(t)

	var wg sync.WaitGroup
	wg.Add(2)

	cb := func(path string, err error) {
		defer wg.Done()
		assert.Empty(path, "Path should be empty")
		assert.EqualError(err, "cannot open file dialog: fyne.App is nil", "Error message should match")
	}

	d.Open("Test", ".", nil, cb)
	d.Save("Test", ".", nil, cb)

	wg.Wait()
}

func TestShowFileDialog(t *testing.T) {
	a := test.NewApp()
	f := NewFyneFallbackDialog(a)
	f.Height = 200
	f.Width = 300

	w := a.NewWindow("Test")
	d := dialog.NewFileOpen(func(uri fyne.URIReadCloser, err error) {}, w)

	filters := filedialog.FileFilters{
		{
			Description: "Test",
			Extensions:  []string{".txt"},
		},
	}

	assert := assert.New(t)

	err := f.showFileDialog(".", filters, d, w)
	assert.NoError(err, "Error should be nil")
	assert.True(w.FixedSize(), "Window should be fixed size")

	err = f.showFileDialog("/not/a/valid/path", filters, d, w)
	assert.Error(err, "Should return an error for invalid path")
}
