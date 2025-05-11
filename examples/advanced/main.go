package main

import (
	"fmt"
	"os"

	"github.com/heathcliff26/godialog"
)

func main() {
	fd := godialog.NewFileDialog()

	pwd, _ := os.Getwd()
	fd.SetInitialDirectory(pwd)

	filters := godialog.FileFilters{
		{
			Description: "All Files",
			Extensions:  []string{""},
		},
		{
			Description: "Image Files (*.jpg, *.jpeg, *.png)",
			Extensions:  []string{".jpg", ".jpeg", ".png"},
		},
	}
	fd.SetFilters(filters)

	res := make(chan string)

	fd.Open("Test Dialog", func(s string, err error) {
		defer close(res)

		if err != nil {
			res <- fmt.Sprintf("Error: %v", err)
		} else {
			res <- fmt.Sprintf("Selected file: '%s'", s)
		}
	})

	println(<-res)
}
