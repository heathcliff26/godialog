package main

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/heathcliff26/godialog"
	fyneFallback "github.com/heathcliff26/godialog/fallback/fyne"
)

var (
	output *widget.Label
)

func main() {
	a := app.New()
	w := a.NewWindow("FileDialog Test App")
	w.Resize(fyne.NewSize(800, 600))

	// Output log field
	output = widget.NewLabel("")
	output.Wrapping = fyne.TextWrapWord
	output.Resize(fyne.NewSize(800, 400))

	// Configuration options

	// Title
	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Enter title")
	titleEntry.Text = "Test Dialog"

	// Start location
	startLocationEntry := widget.NewEntry()
	startLocationEntry.SetPlaceHolder("Enter start location")

	emptyStartButton := widget.NewButton("Empty", func() {
		startLocationEntry.SetText("")
	})
	homedirButton := widget.NewButton("Home", func() {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			appendLog(fmt.Sprintf("Error getting home directory: %v", err))
			return
		}
		startLocationEntry.SetText(homeDir)
	})
	pwdButton := widget.NewButton("PWD", func() {
		pwd, err := os.Getwd()
		if err != nil {
			appendLog(fmt.Sprintf("Error getting current directory: %v", err))
			return
		}
		startLocationEntry.SetText(pwd)
	})

	startLocationContainer := container.NewVBox(
		startLocationEntry,
		container.NewHBox(emptyStartButton, homedirButton, pwdButton),
	)

	// Filter entry
	filterEntry := widget.NewMultiLineEntry()
	filterEntry.SetPlaceHolder("Enter filters in the form Description;.ext1 .ext2")

	allFilterButton := widget.NewButton("All", func() {
		filterText := filterEntry.Text
		if filterText != "" && filterText[len(filterText)-1] != '\n' {
			filterText += "\n"
		}
		filterEntry.SetText(filterText + "All files; ")
	})

	goFilterButton := widget.NewButton("Text", func() {
		filterText := filterEntry.Text
		if filterText != "" && filterText[len(filterText)-1] != '\n' {
			filterText += "\n"
		}
		filterEntry.SetText(filterText + "Text Files (*.txt);.txt")
	})
	markdownFilterButton := widget.NewButton("Markdown", func() {
		filterText := filterEntry.Text
		if filterText != "" && filterText[len(filterText)-1] != '\n' {
			filterText += "\n"
		}
		filterEntry.SetText(filterText + "Markdown (*.md);.md")
	})

	filterContainer := container.NewVBox(
		filterEntry,
		container.NewHBox(allFilterButton, goFilterButton, markdownFilterButton),
	)

	// Buttons
	openButton := widget.NewButton("Open File Dialog", func() {
		fd := prepFileDialog(startLocationEntry.Text, filterEntry.Text)

		appendLog(fmt.Sprintf("Open file dialog. folder: '%s', filters: '%v'", fd.InitialDirectory(), fd.Filters()))

		fd.Open(titleEntry.Text, func(path string, err error) {
			if err != nil {
				appendLog(fmt.Sprintf("Error: %v", err))
				return
			}
			appendLog(fmt.Sprintf("Selected file: %s", path))
		})
	})

	saveButton := widget.NewButton("Save File Dialog", func() {
		fd := prepFileDialog(startLocationEntry.Text, filterEntry.Text)

		appendLog(fmt.Sprintf("Save file dialog. folder: '%s', filters: '%v'", fd.InitialDirectory(), fd.Filters()))

		fd.Save(titleEntry.Text, func(path string, err error) {
			if err != nil {
				appendLog(fmt.Sprintf("Error: %v", err))
				return
			}
			appendLog(fmt.Sprintf("Saved file: %s", path))
		})
	})

	fyneOpenButton := widget.NewButton("Fyne Open", func() {
		fallback := fyneFallback.NewFyneFallbackDialog(a)
		fallback.Open(titleEntry.Text, startLocationEntry.Text, parseFilters(filterEntry.Text), func(path string, err error) {
			if err != nil {
				appendLog(fmt.Sprintf("Error: %v", err))
				return
			}
			appendLog(fmt.Sprintf("Selected file: %s", path))
		})
	})

	fyneSaveButton := widget.NewButton("Fyne Save", func() {
		fallback := fyneFallback.NewFyneFallbackDialog(a)
		fallback.Save(titleEntry.Text, startLocationEntry.Text, parseFilters(filterEntry.Text), func(path string, err error) {
			if err != nil {
				appendLog(fmt.Sprintf("Error: %v", err))
				return
			}
			appendLog(fmt.Sprintf("Saved file: %s", path))
		})
	})

	configForm := container.NewVBox(
		widget.NewLabel("Configuration"),
		widget.NewForm(
			widget.NewFormItem("Title", titleEntry),
			widget.NewFormItem("Start Location", startLocationContainer),
			widget.NewFormItem("Filters", filterContainer),
		),
	)

	buttons := container.NewHBox(openButton, saveButton, fyneOpenButton, fyneSaveButton)
	content := container.NewBorder(configForm, buttons, nil, nil, output)

	w.SetContent(content)
	w.ShowAndRun()
}

func prepFileDialog(startLocationText string, filterText string) godialog.FileDialog {
	filters := parseFilters(filterText)
	fd := godialog.NewFileDialog()
	fd.SetInitialDirectory(startLocationText)
	fd.SetFilters(filters)
	return fd
}

func parseFilters(filterText string) godialog.FileFilters {
	if filterText == "" {
		return nil
	}
	var filters godialog.FileFilters

	for _, line := range strings.Split(filterText, "\n") {
		s := strings.Split(line, ";")
		if len(s) != 2 {
			appendLog(fmt.Sprintf("Invalid filter format: %s", line))
			continue
		}
		filter := godialog.FileFilter{
			Description: s[0],
			Extensions:  strings.Split(s[1], " "),
		}
		filters = append(filters, filter)
	}

	return filters
}

func appendLog(message string) {
	fyne.Do(func() {
		output.SetText(output.Text + message + "\n")
	})
}
