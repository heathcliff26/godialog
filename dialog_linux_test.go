//go:build linux

package godialog

import (
	"testing"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestParseDBusResponse(t *testing.T) {
	tMatrix := []struct {
		Name          string
		Input         *dbus.Signal
		Path          string
		ErrorContains string
		Ignored       bool
	}{
		{
			Name: "ValidResponse",
			Input: &dbus.Signal{
				Name: DBusResponseName,
				Body: []interface{}{
					uint32(0),
					map[string]dbus.Variant{
						"uris": dbus.MakeVariant([]string{"file:///home/user/file.txt"}),
					},
				},
			},
			Path: "/home/user/file.txt",
		},
		{
			Name:          "Nil",
			ErrorContains: "received nil signal from dbus",
		},
		{
			Name:          "InvlidBodyLength",
			Input:         &dbus.Signal{Name: DBusResponseName, Body: []interface{}{uint32(0)}},
			ErrorContains: "invalid response from dbus, invalid response body",
		},
		{
			Name:          "NoResponseSignal",
			Input:         &dbus.Signal{Name: DBusResponseName, Body: []interface{}{"not-a-number", "some other value"}},
			ErrorContains: "invalid response from dbus, no response signal",
		},
		{
			Name:  "Cancelled",
			Input: &dbus.Signal{Name: DBusResponseName, Body: []interface{}{uint32(1), "some other value"}},
		},
		{
			Name:          "NoResults",
			Input:         &dbus.Signal{Name: DBusResponseName, Body: []interface{}{uint32(0), "some other value"}},
			ErrorContains: "invalid response from dbus, no results",
		},
		{
			Name: "NoUris",
			Input: &dbus.Signal{
				Name: DBusResponseName,
				Body: []interface{}{
					uint32(0),
					map[string]dbus.Variant{
						"not-uris": dbus.MakeVariant([]string{"file:///home/user/file.txt"}),
					},
				},
			},
			ErrorContains: "invalid response from dbus, no uris provided",
		},
		{
			Name: "UrisWrongType",
			Input: &dbus.Signal{
				Name: DBusResponseName,
				Body: []interface{}{
					uint32(0),
					map[string]dbus.Variant{
						"uris": dbus.MakeVariant(0),
					},
				},
			},
			ErrorContains: "invalid response from dbus, uris have the wrong type",
		},
		{
			Name: "NoUris",
			Input: &dbus.Signal{
				Name: DBusResponseName,
				Body: []interface{}{
					uint32(0),
					map[string]dbus.Variant{
						"uris": dbus.MakeVariant([]string{}),
					},
				},
			},
			ErrorContains: "response indicated success but no path was selected",
		},
		{
			Name: "MisformedPath",
			Input: &dbus.Signal{
				Name: DBusResponseName,
				Body: []interface{}{
					uint32(0),
					map[string]dbus.Variant{
						"uris": dbus.MakeVariant([]string{"file:///%ÖÖ"}),
					},
				},
			},
			ErrorContains: "failed to unescape path",
		},
		{
			Name: "WrongResponseName",
			Input: &dbus.Signal{
				Name: "NotTheResponse",
				Body: []interface{}{
					uint32(0),
					map[string]dbus.Variant{
						"uris": dbus.MakeVariant([]string{"file:///home/user/file.txt"}),
					},
				},
			},
			Ignored: true,
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			require := require.New(t)
			var path string
			var err error
			var ignored bool
			require.NotPanics(func() {
				path, ignored, err = parseDBusResponse(tCase.Input)
			}, "parseDBusResponse should not panic")
			if tCase.Path != "" {
				require.NoError(err, "Should not return an error")
				require.Equal(tCase.Path, path, "Should return correct path")
			} else if tCase.ErrorContains != "" {
				require.Error(err, "Should return an error")
				require.Empty(path, "should not return a path")
				require.Contains(err.Error(), tCase.ErrorContains, "Error message should contain expected text")
			} else {
				require.NoError(err, "Should not return an error")
				require.Empty(path, "should not return a path")
				require.Equal(tCase.Ignored, ignored, "Should match expected ignore value")
			}
		})
	}
}
