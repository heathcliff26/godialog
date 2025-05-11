package main

import (
	"fmt"

	"github.com/heathcliff26/godialog"
)

func main() {
	fd := godialog.NewFileDialog()

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
